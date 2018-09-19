package xsync

import (
	"fmt"
	"testing"
	"time"
)

func TestSync(t *testing.T) {

	for i := 0; i < 10; i++ {
		tk := Sequence.Get()
		tk.Done()
		b := tk.Wait()
		fmt.Println("wait:", i, b)
		time.Sleep(time.Second)
		tk.Done()
	}
}
