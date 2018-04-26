package event

import (
	"testing"
	"time"
)

// TestNewDelayCallback 测试创建一个延迟回调对象
func TestNewDelayCallback(t *testing.T) {
	delayTime := time.Second * 2
	firstDelay := time.Second * 2
	_, err := NewDelayCallback(delayTime, firstDelay, nil)
	if err == nil {
		t.Error("test fail")
	}

	_, err = NewDelayCallback(delayTime, firstDelay, func(...interface{}) {

	})
	if err != nil {
		t.Errorf("test fail %v", err)
	}
}

// TestDelayCallback 测试延迟回调
func TestDelayCallback(t *testing.T) {
	delayTime := time.Second * 2
	firstDelay := time.Second * 4
	number := 1
	d, err := NewDelayCallback(delayTime, firstDelay, func(data ...interface{}) {
		num := data[0].(int)
		number += num
	})
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	d.Push(2)
	// 没有执行，判断值
	if number != 1 {
		t.Errorf("test fail %d", number)
	}

	time.Sleep(firstDelay)
	time.Sleep(time.Second)

	// 第一次执行之后，判断值
	if number != 3 {
		t.Errorf("test fail %d", number)
	}

	time.Sleep(delayTime)
	time.Sleep(time.Second)

	// 延迟回调执行了，没有添加触发时间，判断值
	if number != 3 {
		t.Errorf("test fail %d", number)
	}

	d.Push(3)

	time.Sleep(delayTime)
	time.Sleep(time.Second)
	// 添加触发事件之后，延迟执行回调函数，判断值
	if number != 6 {
		t.Errorf("test fail %d", number)
	}
	d.Close()

	number = 1
	// 测试进入回调函数才开始push值
	d, err = NewDelayCallback(delayTime, firstDelay, func(data ...interface{}) {
		num := data[0].(int)
		number += num
	})
	if err != nil {
		t.Errorf("test fail %v", err)
	}
	time.Sleep(delayTime + firstDelay)
	d.Push(2)

	time.Sleep(delayTime)
	if number != 3 {
		t.Errorf("test fail %d", number)
	}
}

// TestDelayClose 测试关闭回调对象
func TestDelayClose(t *testing.T) {
	delayTime := time.Second * 2
	firstDelay := time.Second * 1
	number := 1
	d, err := NewDelayCallback(delayTime, firstDelay, func(data ...interface{}) {
		num := data[0].(int)
		number += num
	})
	if err != nil {
		t.Errorf("test fail %v", err)
	}

	d.Push(2)
	// 没有执行，判断值
	if number != 1 {
		t.Errorf("test fail %d", number)
	}

	time.Sleep(firstDelay)
	time.Sleep(time.Second)

	// 第一次执行之后，判断值
	if number != 3 {
		t.Errorf("test fail %d", number)
	}

	d.Push(3)
	d.Close()

	time.Sleep(delayTime)

	// 关闭回调对象之后判断值
	if number != 3 {
		t.Errorf("test fail %d", number)
	}

	// 直接关闭
	b, err := NewDelayCallback(delayTime, firstDelay, func(data ...interface{}) {
		num := data[0].(int)
		number += num
	})
	if err != nil {
		t.Errorf("test fail %v", err)
	}
	b.Close()
}
