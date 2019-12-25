package gocache

import (
	"fmt"
	"time"

	"github.com/micro-plat/lib4go/cache"
	gocache "github.com/zkfy/go-cache"
)

// cacheClient redis配置文件
type cacheClient struct {
	client *gocache.Cache
}

// New 根据配置文件创建一个redis连接
func New(addrs []string, conf string) (m *cacheClient, err error) {
	m = &cacheClient{}
	m.client = gocache.New(5*time.Minute, 10*time.Minute)
	return
}

// Get 根据key获取redis中的数据
func (c *cacheClient) Get(key string) (string, error) {
	v, ok := c.client.Get(key)
	if !ok {
		return "", nil
	}
	return v.(string), nil
}

//Decrement 增加变量的值
func (c *cacheClient) Decrement(key string, delta int64) (n int64, err error) {
	return c.client.DecrementInt64(key, delta)
}

//Increment 减少变量的值
func (c *cacheClient) Increment(key string, delta int64) (n int64, err error) {
	return c.client.IncrementInt64(key, delta)
}

//Gets 获取多条数据
func (c *cacheClient) Gets(key ...string) (r []string, err error) {
	r = make([]string, 0, len(key))
	for _, k := range key {
		v, ok := c.client.Get(k)
		if !ok {
			v = ""
		}
		r = append(r, v.(string))
	}
	return r, nil

}

// Add 添加数据到redis中,如果redis存在，则报错
func (c *cacheClient) Add(key string, value string, expiresAt int) error {
	return c.client.Add(key, value, time.Second*time.Duration(expiresAt))
}

// Set 更新数据到redis中，没有则添加
func (c *cacheClient) Set(key string, value string, expiresAt int) error {
	c.client.Set(key, value, time.Second*time.Duration(expiresAt))
	return nil
}

func (c *cacheClient) Delete(key string) error {
	c.client.Delete(key)
	return nil
}

// Delete 删除redis中的数据
func (c *cacheClient) Exists(key string) bool {
	_, ok := c.client.Get(key)
	return ok
}

// Delay 延长数据在redis中的时间
func (c *cacheClient) Delay(key string, expiresAt int) error {
	expires := time.Duration(expiresAt) * time.Second
	if expiresAt == 0 {
		expires = 0
	}
	v, ok := c.client.Get(key)
	if !ok {
		return fmt.Errorf("%s值不存在", key)
	}
	c.client.Set(key, v, expires)

}
func (c *cacheClient) Close() error {
	return nil
}

type cacheResolver struct {
}

func (s *cacheResolver) Resolve(address []string, conf string) (cache.ICache, error) {
	return New(address, conf)
}
func init() {
	cache.Register("gocache", &cacheResolver{})
}
