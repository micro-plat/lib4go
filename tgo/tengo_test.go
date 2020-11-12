package tgo

import (
	"testing"

	"github.com/d5/tengo/v2/stdlib"
	"github.com/micro-plat/lib4go/assert"
)

func TestModule(t *testing.T) {

	cases := []struct {
		script      string
		moduleName  string
		moduleFunc  CallableFunc
		resultName  string
		resultValue interface{}
	}{
		{script: `http:=import("http");client:=http.get("abc")`, moduleName: "get", moduleFunc: FuncASRS(func(v string) string { return "12345" }), resultName: "client", resultValue: "12345"},
		{script: `http:=import("http");client:=http.get("abc")`, moduleName: "get", moduleFunc: FuncASRS(func(v string) string { return "true" }), resultName: "client", resultValue: "true"},
	}

	for _, cs := range cases {
		vm, err := New(cs.script, NewModule("http").Add(cs.moduleName, cs.moduleFunc))
		assert.Equal(t, nil, err)
		result, err := vm.Run()
		assert.Equal(t, nil, err)
		assert.Equal(t, cs.resultValue, result.GetString(cs.resultName))
	}
}
func TestStrings(t *testing.T) {

	cases := []struct {
		script      string
		moduleName  string
		moduleFunc  CallableFunc
		resultName  string
		resultValue []interface{}
	}{
		{script: `http:=import("http");client:=[http.get("abc"),"123"]`, moduleName: "get", moduleFunc: FuncASRS(func(v string) string { return "12345" }), resultName: "client", resultValue: []interface{}{"12345", "123"}},
		{script: `http:=import("http");client:=[http.get("abc"),"123"]`, moduleName: "get", moduleFunc: FuncASRS(func(v string) string { return "true" }), resultName: "client", resultValue: []interface{}{"true", "123"}},
	}

	for _, cs := range cases {
		vm, err := New(cs.script, NewModule("http").Add(cs.moduleName, cs.moduleFunc))
		assert.Equal(t, nil, err)
		result, err := vm.Run()
		assert.Equal(t, nil, err)
		assert.Equal(t, cs.resultValue, result.GetArray(cs.resultName))
	}
}
func TestBool(t *testing.T) {
	cases := []struct {
		script      string
		moduleName  string
		moduleFunc  CallableFunc
		resultName  string
		resultValue bool
	}{
		{script: `http:=import("http");client:=http.get("abc")`, moduleName: "get", moduleFunc: FuncASRS(func(v string) string { return "true" }), resultName: "client", resultValue: true},
		{script: `http:=import("http");client:=http.get("abc")`, moduleName: "get", moduleFunc: FuncASRS(func(v string) string { return "false" }), resultName: "client", resultValue: false},
	}

	for _, cs := range cases {
		vm, err := New(cs.script, NewModule("http").Add(cs.moduleName, cs.moduleFunc))
		assert.Equal(t, nil, err)
		result, err := vm.Run()
		assert.Equal(t, nil, err)
		assert.Equal(t, cs.resultValue, result.GetBool(cs.resultName))
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
