package mqtt

import (
	"encoding/json"
	"fmt"
	"net"

	proto "github.com/huin/mqtt"
	"github.com/jeffallen/mqtt"
	"github.com/micro-plat/lib4go/queue"
)

// mqttClient memcache配置文件
type mqttClient struct {
	servers []string
	client  *mqtt.ClientConn
}

// New 根据配置文件创建一个redis连接
func New(addrs []string, raw string) (m *mqttClient, err error) {
	m = &mqttClient{servers: addrs}

	conf := &queue.Config{}
	if err := json.Unmarshal([]byte(raw), &conf); err != nil {
		return nil, err
	}

	conn, err := net.Dial("tcp", conf.Addr)
	if err != nil {
		return nil, err
	}
	cc := mqtt.NewClientConn(conn)
	if err = cc.Connect(conf.UserName, conf.Password); err != nil {
		return nil, fmt.Errorf("连接失败:%v(%s-%s/%s)", err, conf.Addr, conf.UserName, conf.Password)
	}
	m.client = cc
	return m, nil
}

// Push 向存于 key 的列表的尾部插入所有指定的值
func (c *mqttClient) Push(key string, value string) error {
	c.client.Publish(&proto.Publish{
		Header:    proto.Header{},
		TopicName: key,
		Payload:   proto.BytesPayload([]byte(value)),
	})
	return nil
}

// Pop 移除并且返回 key 对应的 list 的第一个元素。
func (c *mqttClient) Pop(key string) (string, error) {
	return "", fmt.Errorf("mqtt不支持pop方法")
}

// Pop 移除并且返回 key 对应的 list 的第一个元素。
func (c *mqttClient) Count(key string) (int64, error) {
	return 0, fmt.Errorf("mqtt不支持pop方法")
}

// Close 释放资源
func (c *mqttClient) Close() error {
	c.client.Disconnect()
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
