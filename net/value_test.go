package net

import (
	"testing"

	"github.com/micro-plat/lib4go/assert"
)

func TestV1x(t *testing.T) {
	values, err := NewValuesByQuery("b1=v3&a1=v1&a2=v2")
	assert.Equal(t, err, nil)
	assert.Equal(t, values.Get("b1"), "v3")
	assert.Equal(t, values.Get("a1"), "v1")
}

func TestSign(t *testing.T) {
	values := NewValues()
	values.Set("b1", "v3")
	values.Set("a1", "v1")
	values.Set("a2", "v2")
	a := values.Join("=", "&")
	assert.Equal(t, a, "b1=v3&a1=v1&a2=v2")
	values.Sort()
	a = values.Join("=", "&")
	assert.Equal(t, a, "a1=v1&a2=v2&b1=v3")
	e := values.Encode()
	assert.Equal(t, len(e), 17)

	f := values.Join("=", "&", "key", "123123")
	assert.Equal(t, f, "a1=v1&a2=v2&b1=v3&key=123123")

	g := values.Join("=", "&", "123123")
	assert.Equal(t, g, "a1=v1&a2=v2&b1=v3&123123")

	h := values.Join("", "", "123123")
	assert.Equal(t, h, "a1v1a2v2b1v3123123")

}
