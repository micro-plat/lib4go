package tgo

import "github.com/d5/tengo/v2"

//Module 供脚本使用的模块信息
type Module struct {
	name   string
	object map[string]tengo.Object
}

//NewModule 构建模块
func NewModule(name string) *Module {
	return &Module{
		name:   name,
		object: make(map[string]tengo.Object),
	}
}

//Add 添加函数
func (m *Module) Add(method string, f CallableFunc) *Module {
	m.object[method] = &UserFunction{Name: method, Value: f}
	return m
}
