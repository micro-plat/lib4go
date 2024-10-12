package invokes

import (
	"testing"
	"time"

	"github.com/micro-plat/lib4go/assert"
)

func TestInvoke(t *testing.T) {
	executes := 0
	debouncedFunc := Debounce(func() {
		executes++
	}, 500*time.Millisecond)

	// 模拟高频事件
	for i := 0; i < 5; i++ {
		debouncedFunc()
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(time.Second)
	assert.Equal(t, 1, executes)
}
func TestThrottle(t *testing.T) {
	executes := 0
	throttledFunc := Throttle(func() {
		executes++
	}, 500*time.Millisecond)

	// 模拟高频事件
	for i := 0; i < 10; i++ {
		throttledFunc()
	}
	time.Sleep(time.Second)
	assert.Equal(t, 1, executes)
}
