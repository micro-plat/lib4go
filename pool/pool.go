package pool

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

//pool 存放链接信息
type pool struct {
	mu          sync.Mutex
	conns       chan *idleConn
	factory     func() (interface{}, error)
	close       func(interface{}) error
	done        bool
	idleTimeout time.Duration
}

type idleConn struct {
	conn interface{}
	t    time.Time
}

//Len 连接池中已有的连接
func (c *pool) Len() int {
	return len(c.getConns())
}

var (
	//ErrClosed 连接池已经关闭Error
	ErrClosed = errors.New("pool is closed")

	// AutoReleaseTime 定时自动清除
	TimeOut time.Duration
)

//IPool 基本方法
type IPool interface {
	Get() (interface{}, error)

	Put(interface{}) error

	Close(interface{}) error

	Release()

	Len() int

	AutoReleaseStart()
}

//PoolConfigOptions 连接池相关配置
type PoolConfigOptions struct {
	//连接池中拥有的最小连接数
	InitialCap int
	//连接池中拥有的最大的连接数
	MaxCap int
	//生成连接的方法
	Factory func() (interface{}, error)
	//关闭链接的方法
	Close func(interface{}) error
	//链接最大空闲时间，超过该时间则将失效
	IdleTimeout time.Duration
}

//New 初始化链接
func New(config *PoolConfigOptions) (IPool, error) {
	if config.InitialCap < 0 || config.MaxCap <= 0 || config.InitialCap > config.MaxCap {
		return nil, errors.New("invalid capacity settings")
	}

	/*add by champly 2016年12月12日14:06:14*/
	if config.Factory == nil || config.Close == nil {
		return nil, errors.New("invalid function settings")
	}

	// 添加设置自动清理超时连接
	TimeOut = config.IdleTimeout
	if TimeOut == 0 {
		TimeOut = time.Hour * 30
	}
	/*end*/

	c := &pool{
		conns:       make(chan *idleConn, config.MaxCap),
		factory:     config.Factory,
		close:       config.Close,
		idleTimeout: config.IdleTimeout,
	}
	for i := 0; i < config.InitialCap; i++ {
		conn, err := c.factory()
		if err != nil {
			c.Release()
			return nil, fmt.Errorf("factory is not able to fill the pool: %s", err)
		}
		c.conns <- &idleConn{conn: conn, t: time.Now()}
	}

	return c, nil
}

//getConns 获取所有连接
func (c *pool) getConns() chan *idleConn {
	c.mu.Lock()
	conns := c.conns
	c.mu.Unlock()
	return conns
}

//Get 从pool中取一个连接
func (c *pool) Get() (interface{}, error) {
	conns := c.getConns()
	if conns == nil || c.done {
		return nil, ErrClosed
	}
	for {
		select {
		case wrapConn := <-conns:
			if wrapConn == nil {
				return nil, ErrClosed
			}
			// 判断是否超时，超时则丢弃
			if timeout := c.idleTimeout; timeout > 0 {
				if wrapConn.t.Add(timeout).Before(time.Now()) {
					//丢弃并关闭该链接
					c.Close(wrapConn.conn)
					continue
				}
			}

			return wrapConn.conn, nil
		default:
			conn, err := c.factory()
			if err != nil {
				return nil, err
			}

			return conn, nil
		}
	}
}

//Put 将连接放回pool中
func (c *pool) Put(conn interface{}) error {
	if conn == nil {
		return errors.New("connection is nil. rejecting")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conns == nil || c.done {
		return c.Close(conn)
	}

	select {
	case c.conns <- &idleConn{conn: conn, t: time.Now()}:
		return nil
	default:
		return c.Close(conn)
	}
}

//Close 关闭单条连接
func (c *pool) Close(conn interface{}) error {
	if conn == nil {
		return errors.New("connection is nil. rejecting")
	}
	return c.close(conn)
}

//Release 释放连接池中所有链接
func (c *pool) Release() {
	c.mu.Lock()
	conns := c.conns
	c.conns = nil
	c.factory = nil
	closeFun := c.close
	c.close = nil
	c.mu.Unlock()

	if conns == nil {
		return
	}

	close(conns)
	for wrapConn := range conns {
		closeFun(wrapConn.conn)
	}
}

/*add by champly 2016年12月12日16:24:35*/
// AutoReleaseTimeout 自动清除超时的连接
func (c *pool) AutoReleaseStart() {
	go func() {
		for {
			select {
			case <-time.After(TimeOut):
				c.clear()
				if c.done {
					return
				}
			}
		}
	}()
}

func (c *pool) clear() {
	start := time.Now()
	if c.Len() > 0 {
		c.mu.Lock()
		conns := c.conns
		defer c.mu.Unlock()
		if conns == nil {
			return
		}

		fmt.Println("pool中总连接数：", len(c.conns))
		length := len(c.conns)
		for i := 0; i < length; i++ {
			wrapConn, ok := <-conns
			if ok {
				if wrapConn == nil {
					return
				}
				// 如果超时
				if timeout := c.idleTimeout; timeout > 0 {
					if wrapConn.t.Add(timeout).Before(time.Now()) {
						c.Close(wrapConn.conn)
						continue
					}
				}
				// 如果没有超时, 写回idleConn
				c.conns <- &idleConn{conn: wrapConn.conn, t: time.Now()}
			}
		}
	}

	fmt.Println("总共耗时：", time.Now().Sub(start), " 现在连接数：", len(c.conns))
}

/*end*/
