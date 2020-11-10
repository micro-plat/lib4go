package tgo

import (
	"testing"

	"github.com/d5/tengo/v2/stdlib"
	"github.com/micro-plat/lib4go/assert"
)

func TestLib(t *testing.T) {

	cases := []struct {
		script string
		fname  string
		fun    CallableFunc
		name   string
		expect interface{}
	}{
		{script: `http:=import("http");client:=http.get("abc")`, fname: "get", name: "client", expect: "12345", fun: FuncASRS(func(v string) string { return "12345" })},
		{script: `http:=import("http");client:=http.get("abc")`, fname: "get", name: "client", expect: "true", fun: FuncASRS(func(v string) string { return "true" })},
	}

	for _, cs := range cases {
		vm, err := New(cs.script, NewModule("http").Add(cs.fname, cs.fun))
		assert.Equal(t, nil, err)
		result, err := vm.Run()
		assert.Equal(t, nil, err)
		assert.Equal(t, cs.expect, result.GetString(cs.name))
	}

}
func BenchmarkTest(t *testing.B) {

	cases := []struct {
		script string
		fname  string
		fun    CallableFunc
		name   string
		expect interface{}
	}{
		{script: `http:=import("http");client:=http.get("abc")`, fname: "get", name: "client", expect: "12345", fun: stdlib.FuncASRS(func(v string) string { return "12345" })},
	}

	for _, cs := range cases {
		vm, err := New(cs.script, NewModule("http").Add(cs.fname, cs.fun))
		assert.Equal(t, nil, err)
		result, err := vm.Run()
		assert.Equal(t, nil, err)
		assert.Equal(t, cs.expect, result.GetString(cs.name))
	}

}
