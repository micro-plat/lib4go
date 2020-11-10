package tengoscript

import (
	"testing"

	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"
	"github.com/micro-plat/lib4go/assert"
)

func TestLib(t *testing.T) {

	cases := []struct {
		script  string
		modules map[string]Object
		name    string
		expect  interface{}
	}{
		{script: `http:=import("http");client:=http.get("abc")`, name: "client", expect: "12345", modules: map[string]Object{
			"get": &tengo.UserFunction{Name: "get", Value: stdlib.FuncASRS(func(v string) string { return "12345" })}}},
		{script: `http:=import("http");client:=http.get("abc")`, name: "client", expect: "true", modules: map[string]Object{
			"get": &tengo.UserFunction{Name: "get", Value: stdlib.FuncASRS(func(v string) string { return "true" })}}},
	}

	for _, cs := range cases {
		vm, err := New(cs.script, WithModule("http", cs.modules))
		assert.Equal(t, nil, err)
		result, err := vm.Run()
		assert.Equal(t, nil, err)
		assert.Equal(t, cs.expect, result.GetString(cs.name))
	}

}
func BenchmarkTest(t *testing.B) {
	cases := []struct {
		script  string
		modules map[string]Object
		name    string
		expect  interface{}
	}{
		{script: `http:=import("http");client:=http.get("abc")`, name: "client", expect: "12345", modules: map[string]Object{
			"get": &tengo.UserFunction{Name: "get", Value: stdlib.FuncASRS(func(v string) string { return "12345" })}}},
	}

	for _, cs := range cases {
		vm, err := New(cs.script, WithModule("http", cs.modules))
		assert.Equal(t, nil, err)
		t.ResetTimer()
		for i := 0; i < t.N; i++ {
			result, err := vm.Run()
			assert.Equal(t, nil, err)
			assert.Equal(t, cs.expect, result.GetString(cs.name))
		}
	}

}
