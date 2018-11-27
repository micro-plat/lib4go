package net

import "testing"
import "github.com/micro-plat/lib4go/ut"

func TestSign(t *testing.T) {
	values := NewValues()
	values.Set("b1", "v3")
	values.Set("a1", "v1")
	values.Set("a2", "v2")
	a := values.Join("=", "&")
	ut.Expect(t, a, "b1=v3&a1=v1&a2=v2")
	values.Sort()
	a = values.Join("=", "&")
	ut.Expect(t, a, "a1=v1&a2=v2&b1=v3")
	e := values.Encode()
	ut.Expect(t, len(e), 17)

	f := values.Join("=", "&", "key", "123123")
	ut.Expect(t, f, "a1=v1&a2=v2&b1=v3&key=123123")
	g := values.Join("=", "&", "123123")
	ut.Expect(t, g, "a1=v1&a2=v2&b1=v3&123123")
}
