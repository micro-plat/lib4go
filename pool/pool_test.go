package pool

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"
)

func Factory() (interface{}, error) {
	return 1, nil
}

func FactoryErr() (interface{}, error) {
	return nil, errors.New("错误的工厂方法")
}

func Close(interface{}) error {
	return nil
}

// TestNew 测试初始化一个pool
func TestNew(t *testing.T) {
	// 正常调用
	timeout := time.Second * 2
	config := &PoolConfigOptions{InitialCap: 1, MaxCap: 10, Factory: Factory, Close: Close, IdleTimeout: timeout}
	_, err := New(config)
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	// 传入的函数为nil
	config = &PoolConfigOptions{InitialCap: 1, MaxCap: 10, Factory: nil, Close: nil}
	_, err = New(config)
	if !strings.EqualFold(err.Error(), "invalid function settings") {
		t.Errorf("test fail %v", err)
	}

	// 传入的数字不对
	config = &PoolConfigOptions{InitialCap: -1, MaxCap: 10, Factory: Factory, Close: Close}
	_, err = New(config)
	if !strings.EqualFold(err.Error(), "invalid capacity settings") {
		t.Errorf("test fail %v", err)
	}

	// 传入的数字不对
	config = &PoolConfigOptions{InitialCap: 2, MaxCap: 1, Factory: Factory, Close: Close}
	_, err = New(config)
	if !strings.EqualFold(err.Error(), "invalid capacity settings") {
		t.Errorf("test fail %v", err)
	}

	// 传入的数字不对
	config = &PoolConfigOptions{InitialCap: 2, MaxCap: -1, Factory: Factory, Close: Close}
	_, err = New(config)
	if !strings.EqualFold(err.Error(), "invalid capacity settings") {
		t.Errorf("test fail %v", err)
	}

	// 工厂方法返回错误的信息
	config = &PoolConfigOptions{InitialCap: 2, MaxCap: 10, Factory: FactoryErr, Close: Close}
	_, err = New(config)
	if !strings.EqualFold(err.Error(), "factory is not able to fill the pool: 错误的工厂方法") {
		t.Errorf("test fail %v", err)
	}
}

// TestGet 测试从pool中获取一个连接
func TestGet(t *testing.T) {
	// pool中有连接
	timeout := time.Second * 2
	config := &PoolConfigOptions{InitialCap: 1, MaxCap: 10, Factory: Factory, Close: Close, IdleTimeout: timeout}
	c, err := New(config)
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	data, err := c.Get()
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	if data.(int) != 1 {
		t.Error("test fail")
	}

	// pool 中没有连接
	config = &PoolConfigOptions{InitialCap: 0, MaxCap: 10, Factory: Factory, Close: Close, IdleTimeout: timeout}
	c, err = New(config)
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	data, err = c.Get()
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	if data.(int) != 1 {
		t.Error("test fail")
	}

	// 超时获取
	config = &PoolConfigOptions{InitialCap: 2, MaxCap: 10, Factory: Factory, Close: Close, IdleTimeout: timeout}
	c, err = New(config)
	if err != nil {
		t.Errorf("test fail %v", err)
	}
	time.Sleep(time.Second * 3)

	data, err = c.Get()
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	if data.(int) != 1 {
		t.Error("test fail")
	}
}

// TestPut 测试将连接放回pool
func TestPut(t *testing.T) {
	maxCap := 3
	timeout := time.Second * 2
	config := &PoolConfigOptions{InitialCap: 2, MaxCap: maxCap, Factory: Factory, Close: Close, IdleTimeout: timeout}
	c, err := New(config)
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	// put一个nil
	err = c.Put(nil)
	if !strings.EqualFold(err.Error(), "connection is nil. rejecting") {
		t.Errorf("test fail %v", err)
	}

	// 获取一个，然后放回去
	conn, err := c.Get()
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	err = c.Put(conn)
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	// 连接池满了，放回去直接关闭
	config = &PoolConfigOptions{InitialCap: 2, MaxCap: maxCap, Factory: Factory, Close: Close, IdleTimeout: timeout}
	c, err = New(config)
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	conns := make([]interface{}, 0, maxCap+1)
	for i := 0; i < maxCap; i++ {
		conn, err := c.Get()
		if err != nil {
			t.Errorf("test fail %v", err)
		}
		conns = append(conns, conn)
	}

	conn, err = c.Get()
	if err != nil {
		t.Errorf("test fail %v", err)
	}
	conns = append(conns, conn)

	if len(conns) != maxCap+1 {
		t.Errorf("test fail actual : %d, except : %d", len(conns), maxCap+1)
	}

	// 开始放回去
	for _, conn := range conns {
		c.Put(conn)
	}
}

// TestClose 测试关闭一个链接
func TestClose(t *testing.T) {
	timeout := time.Second * 2
	config := &PoolConfigOptions{InitialCap: 2, MaxCap: 10, Factory: Factory, Close: Close, IdleTimeout: timeout}
	c, err := New(config)
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	conn, err := c.Get()
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	// 关闭这个连接，然后放回去
	err = c.Close(conn)
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	err = c.Put(conn)
	if err != nil {
		t.Errorf("test fail %v", err)
	}
}

// TestLen 测试获取pool中的连接数
func TestLen(t *testing.T) {
	// 正常调用
	initialCap := 2
	timeout := time.Second * 2
	config := &PoolConfigOptions{InitialCap: initialCap, MaxCap: 10, Factory: Factory, Close: Close, IdleTimeout: timeout}
	c, err := New(config)
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	// 取出一个，判断个数
	conn, _ := c.Get()
	if c.Len() != initialCap-1 {
		t.Errorf("test fail")
	}

	// 然后放回去，判断个数
	c.Put(conn)
	if c.Len() != initialCap {
		t.Errorf("test fail")
	}

	// 多方一个回去
	c.Put(conn)
	if c.Len() != initialCap+1 {
		t.Errorf("test fail")
	}

}

// TestRelease 测试释放pool中所有连接
func TestRelease(t *testing.T) {
	initialCap := 2
	timeout := time.Second * 2
	config := &PoolConfigOptions{InitialCap: initialCap, MaxCap: 10, Factory: Factory, Close: Close, IdleTimeout: timeout}
	c, err := New(config)
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	if c.Len() != initialCap {
		t.Errorf("test fail")
	}

	// 释放所有连接
	c.Release()

	if c.Len() != 0 {
		t.Errorf("test fail")
	}
}

// TestAutoRelease 测试自动释放超时的连接
func TestAutoRelease(t *testing.T) {
	initialCap := 2
	timeout := time.Second * 2
	config := &PoolConfigOptions{InitialCap: initialCap, MaxCap: initialCap * 2, Factory: Factory, Close: Close, IdleTimeout: timeout}
	c, err := New(config)
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	if c.Len() != initialCap {
		t.Errorf("test fail")
	}

	c.AutoReleaseStart()

	time.Sleep(time.Second * 3)
	if c.Len() != 0 {
		fmt.Println("test", c.Len())
		t.Errorf("test fail")
	}
}
