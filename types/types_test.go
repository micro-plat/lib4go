package types

import (
	"testing"

	"github.com/micro-plat/lib4go/assert"
)

func TestIsNotEmpty(t *testing.T) {
	if IsEmpty(900) {
		t.Error("nil:false")
	}
	if IsEmpty("222") {
		t.Error("'':false")
	}
	if IsEmpty(10) {
		t.Error("'':false")
	}
	if IsEmpty([]string{"abc"}) {
		t.Error("[]string{}:false")
	}
	if IsEmpty(map[string]string{"a": "b"}) {
		t.Error("[]string{}:false")
	}
	if IsEmpty(map[string]interface{}{"e": 100}) {
		t.Error("[]string{}:false")
	}
	if IsEmpty(struct{ a int }{a: 100}) {
		t.Error("struct{}{}:false")
	}
	s := struct{ a string }{a: "abc"}
	if IsEmpty(&s) {
		t.Error("struct{}{}:false")
	}
	v := make(chan int, 1)
	v <- 1
	if IsEmpty(v) {
		t.Error("chan int:false")
	}
}
func TestIsEmpty(t *testing.T) {
	if !IsEmpty(nil) {
		t.Error("nil:false")
	}
	if !IsEmpty("") {
		t.Error("'':false")
	}
	if !IsEmpty(0) {
		t.Error("'':false")
	}
	if !IsEmpty([]string{}) {
		t.Error("[]string{}:false")
	}
	if !IsEmpty(map[string]string{}) {
		t.Error("map[string]string{}:false")
	}
	if !IsEmpty(map[string]interface{}{}) {
		t.Error("map[string]interface{}{}:false")
	}
	if !IsEmpty(struct{}{}) {
		t.Error("struct{}{}:false")
	}
	s := struct{}{}
	if !IsEmpty(&s) {
		t.Error("struct{}{}:false")
	}
	if !IsEmpty(make(chan int, 0)) {
		t.Error("chan int:false")
	}
}
func BenchmarkTest(b *testing.B) {
	assert.Equal(b, DecodeString("3", "2", "3", "3", "2", "4"), "2")

}

func TestDecode1(t *testing.T) {
	assert.Equal(t, DecodeString(1, 2, 3), "")
	assert.Equal(t, DecodeString("1", 2, 3), "1")
	assert.Equal(t, DecodeString(2, 2, 3), "3")
	assert.Equal(t, DecodeString(1, 2, 3, 4), "4")
	assert.Equal(t, DecodeString(3, 2, 3, 4), "4")
	assert.Equal(t, DecodeString(3, 2, 3, 3, 2, 4), "2")

}
func TestDecode2(t *testing.T) {
	assert.Equal(t, DecodeInt(1, 2, 3, 1), 1)
	assert.Equal(t, DecodeInt(2, 2, 3, 2), 3)
	assert.Equal(t, DecodeInt(1, 2, 3, 4), 4)
	assert.Equal(t, DecodeInt(3, 2, 3, 4), 4)
	assert.Equal(t, DecodeInt(3, 2, 3, 3, 2, 4), 2)

}
func TestDecode3(t *testing.T) {
	assert.Equal(t, DecodeInt(1, 2, "3", 1), 1)
	assert.Equal(t, DecodeInt(2, 2, "3", 3), 3)
	assert.Equal(t, DecodeInt(1, 2, "3", "4"), 4)
	assert.Equal(t, DecodeInt(3, 2, "3", "4"), 4)
	assert.Equal(t, DecodeInt(3, 2, "3", 3, "2", "4"), 2)
	assert.Equal(t, DecodeInt(3, 2, "3", "3", "2", "4"), 2)
}
func TestDecode4(t *testing.T) {
	assert.Equal(t, DecodeInt(1, 1, 0, 1), 0)
	assert.Equal(t, DecodeInt(0, 1, 0, 1), 1)
	assert.Equal(t, DecodeInt("0", "1", 0, 1), 1)
	assert.Equal(t, DecodeInt("1", "1", 0, 1), 0)

	assert.Equal(t, DecodeInt("0", 1, 0, 1), 1)
	assert.Equal(t, DecodeInt("1", 1, 0, 1), 0)
	assert.Equal(t, DecodeInt(float64(1), 1, 0, 1), 0)
}
