package memcache

import (
	"strings"
	"testing"
	"time"
)

func TestMemcache(t *testing.T) {
	mem, err := NewJSON(`{"servers": ["192.168.0.166:11212"]}`)
	if err != nil {
		t.Error(err)
		return
	}

	// 正常流程
	key := "test"
	value := "value"
	if e := mem.Add(key, value, 5); e != nil {
		t.Errorf("存数据失败：%v", e)
	}

	v := mem.Get(key)
	if !strings.EqualFold(v, value) {
		t.Errorf("取数据失败：%s", v)
	}

	err = mem.Delay(key, 10)
	if err != nil {
		t.Errorf("延长存储时间失败:%v", err)
	}

	time.Sleep(8 * time.Second)

	v = mem.Get(key)
	if !strings.EqualFold(v, value) {
		t.Error("延长存储时间失败")
	}

	mem.Delete(key)
	v = mem.Get(key)
	if !strings.EqualFold(v, "") {
		t.Errorf("获取到了脏数:%s", v)
	}

	// 取一个不存在的数据
	key = "test1"
	value = "value1"
	if mem.Get(key) != "" {
		t.Error("get unkonw err:")
	}

	// Add一个存在的数据
	key = "test_add"
	value = "value_add"
	value2 := "value_add_2"
	if err := mem.Add(key, value, 10); err != nil {
		t.Errorf("存数据失败:%v", err)
	}
	if err := mem.Add(key, value2, 10); err == nil {
		t.Error("test fail")
	}

	// Set一个存在的数据
	key = "test_set"
	value = "value_set"
	value2 = "value_set_2"
	if err := mem.Set(key, value, 10); err != nil {
		t.Errorf("Set data err : %v", err)
	}
	if err := mem.Set(key, value2, 10); err != nil {
		t.Errorf("Set data err : %v", err)
	}

	// 配置文件错误
	mem, err = NewJSON(`192.168.0.166:11212`)
	if err == nil {
		t.Error(err)
		return
	}

	// 配置文件错误
	mem, err = NewJSON(`{"servers": ["192.168.0.165:11212"]}`)
	if err != nil {
		t.Error(err)
		return
	}
	if e := mem.Add(key, value, 5); e == nil {
		t.Error("test fail")
	}

	v = mem.Get(key)
	if !strings.EqualFold(v, "") {
		t.Error("test fail")
	}

	err = mem.Delay(key, 10)
	if err == nil {
		t.Error("test fail")
	}

	mem.Delete(key)
}
