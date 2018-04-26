package metrics

import (
	"fmt"
	"testing"

	"github.com/qxnw/lib4go/ut"
)

func TestRingCounter1(t *testing.T) {
	counter := NewRPSC(6, 10)

	counter.mark(1, int64(0))
	ut.ExpectSkip(t, counter.counter, int64(1))
	counter.mark(1, int64(1))
	ut.ExpectSkip(t, counter.counter, int64(2))
	counter.mark(1, int64(2))
	ut.ExpectSkip(t, counter.counter, int64(3))
	counter.mark(1, int64(3))
	ut.ExpectSkip(t, counter.counter, int64(4))
	counter.mark(1, int64(4))
	ut.ExpectSkip(t, counter.counter, int64(5))
	counter.mark(1, int64(5))
	ut.ExpectSkip(t, counter.counter, int64(6))
	counter.mark(1, int64(6))
	ut.ExpectSkip(t, counter.counter, int64(6))

	counter.mark(12, int64(12))
	ut.ExpectSkip(t, counter.counter, int64(12))

	counter.mark(13, int64(100))
	ut.ExpectSkip(t, counter.counter, int64(13))

	counter.mark(0, int64(101))
	ut.ExpectSkip(t, counter.counter, int64(13))

}
func TestRingCounter2(t *testing.T) {
	counter := NewRPSC(6, 10)
	for i := 0; i < 6; i++ {
		counter.mark(1, int64(i))
	}

	counter.mark(1, int64(8))
	ut.ExpectSkip(t, counter.counter, int64(4))
}
func TestRingCounter3(t *testing.T) {
	counter := NewRPSC(6, 10)
	for i := 0; i < 6; i++ {
		counter.mark(1, int64(i))
	}
	counter.mark(1, int64(10))
	ut.ExpectSkip(t, counter.counter, int64(2))
}
func TestRingCounter4(t *testing.T) {
	counter := NewRPSC(6, 10)
	counter.mark(1, int64(10))
	ut.ExpectSkip(t, counter.counter, int64(1))
}
func TestRingCounter5(t *testing.T) {
	counter := NewRPSC(6, 10)
	counter.mark(1, int64(6))
	ut.ExpectSkip(t, counter.counter, int64(1))
	counter.mark(1, int64(11))
	ut.ExpectSkip(t, counter.counter, int64(2))
	//counter.mark(1, int64(17))
	//ut.ExpectSkip(t, counter.counter, int32(1))
}
func TestRingCounter6(t *testing.T) {
	counter := NewRPSC(6, 10)
	counter.mark(1, int64(2))
	ut.ExpectSkip(t, counter.counter, int64(1))
	counter.mark(1, int64(7))
	ut.ExpectSkip(t, counter.counter, int64(2))
	fmt.Println("---------------")
	counter.mark(0, int64(9))
	ut.ExpectSkip(t, counter.counter, int64(1))
}
func TestRingCounter61(t *testing.T) {
	counter := NewRPSC(60, 3600)
	counter.mark(1, int64(1))
	ut.ExpectSkip(t, counter.counter, int64(1))
	counter.mark(1, int64(7))
	ut.ExpectSkip(t, counter.counter, int64(2))
	fmt.Println("---------------")
	counter.mark(0, int64(9))
	ut.ExpectSkip(t, counter.counter, int64(2))
	counter.mark(1, int64(59))
	ut.ExpectSkip(t, counter.counter, int64(3))

	counter.mark(1, int64(61))
	ut.ExpectSkip(t, counter.counter, int64(3))

	counter.mark(1, int64(130))
	ut.ExpectSkip(t, counter.counter, int64(1))

}

func BenchmarkRing1(b *testing.B) {
	counter := NewRPSC(6, 10)
	for i := 0; i < b.N; i++ {
		counter.mark(int64(i), int64(i))
	}
}
func TestRingCounter62(t *testing.T) {
	counter := NewRPSC(60, 3600)
	for i := 0; i < 1000; i++ {
		counter.mark(1, int64(1))
	}
	ut.ExpectSkip(t, counter.counter, int64(1000))
}
