package invokes

import (
	"sync"
	"time"
)

// 防抖函数类型
type DebounceFunc func()

// Debounce 创建一个防抖函数，该函数只会在指定的持续时间内没有被再次调用后才执行。
func Debounce(f DebounceFunc, wait time.Duration) DebounceFunc {
	var mu sync.Mutex
	var timer *time.Timer

	return func() {
		mu.Lock()
		defer mu.Unlock()
		if timer != nil {
			timer.Stop()
		}
		timer = time.AfterFunc(wait, f)
	}
}

// 节流函数类型
type ThrottleFunc func()

// Throttle 创建一个节流函数，该函数在指定的持续时间内最多只执行一次。
func Throttle(f ThrottleFunc, wait time.Duration) ThrottleFunc {
	var mu sync.Mutex
	var lastRan time.Time

	return func() {
		mu.Lock()
		defer mu.Unlock()
		now := time.Now()
		// 如果当前时间与上次执行时间的差值大于等于指定的持续时间，则执行函数。
		if now.Sub(lastRan) >= wait {
			lastRan = now
			// 可选：在新的goroutine中执行f，以避免阻塞调用者。
			// 如果不需要异步执行，可以去掉go关键字。
			go f()
		}
	}
}
