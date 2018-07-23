package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/micro-plat/lib4go/logger"
	"github.com/micro-plat/lib4go/net"
	"github.com/micro-plat/lib4go/utility"

	"github.com/micro-plat/lib4go/queue"
	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
)

// MQTTClient memcache配置文件
type MQTTClient struct {
	servers []string
	client  *client.Client
	once    sync.Once
	Logger  *logger.Logger
	done    bool
}

//New 根据配置文件创建一个redis连接
func New(addrs []string, raw string) (m *MQTTClient, err error) {
	m = &MQTTClient{servers: addrs, Logger: logger.GetSession("mqtt.publisher", logger.CreateSession())}

	conf := &queue.Config{}
	if err := json.Unmarshal([]byte(raw), &conf); err != nil {
		return nil, err
	}
	cc := client.New(&client.Options{
		ErrorHandler: func(err error) {
			m.Logger.Error("mqtt.publisher出错:", err)
		},
	})
	cert, err := m.getCert(conf)
	if err != nil {
		return nil, err
	}
	if err := cc.Connect(&client.ConnectOptions{
		Network:   "tcp",
		Address:   conf.Addr,
		UserName:  []byte(conf.UserName),
		Password:  []byte(conf.Password),
		ClientID:  []byte(fmt.Sprintf("%s-%s", net.GetLocalIPAddress(), utility.GetGUID()[0:6])),
		TLSConfig: cert,
		KeepAlive: 3,
	}); err != nil {
		return nil, fmt.Errorf("连接失败:%v(%s-%s/%s)", err, conf.Addr, conf.UserName, conf.Password)
	}
	m.client = cc
	return m, nil
}

func (c *MQTTClient) getCert(conf *queue.Config) (*tls.Config, error) {
	if conf.CertPath == "" {
		return nil, nil
	}
	b, err := ioutil.ReadFile(conf.CertPath)
	if err != nil {
		return nil, fmt.Errorf("读取证书失败:%s(%v)", conf.CertPath, err)
	}
	roots := x509.NewCertPool()
	if ok := roots.AppendCertsFromPEM(b); !ok {
		return nil, fmt.Errorf("failed to parse root certificate")
	}
	return &tls.Config{
		RootCAs: roots,
	}, nil
}

// Push 向存于 key 的列表的尾部插入所有指定的值
func (c *MQTTClient) Push(key string, value string) error {
	if c.done {
		return fmt.Errorf("队列已关闭")
	}
	return c.client.Publish(&client.PublishOptions{
		QoS:       mqtt.QoS0,
		TopicName: []byte(key),
		Message:   []byte(value),
	})
}

// Pop 移除并且返回 key 对应的 list 的第一个元素。
func (c *MQTTClient) Pop(key string) (string, error) {
	return "", fmt.Errorf("mqtt不支持pop方法")
}

// Pop 移除并且返回 key 对应的 list 的第一个元素。
func (c *MQTTClient) Count(key string) (int64, error) {
	return 0, fmt.Errorf("mqtt不支持pop方法")
}

// Close 释放资源
func (c *MQTTClient) Close() error {
	c.done = true
	c.once.Do(func() {
		c.client.Disconnect()
		c.client.Terminate()
	})
	return nil
}

type redisResolver struct {
}

func (s *redisResolver) Resolve(address []string, conf string) (queue.IQueue, error) {
	return New(address, conf)
}
func init() {
	queue.Register("mqtt", &redisResolver{})
}
