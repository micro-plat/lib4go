package tengoscript

import "github.com/d5/tengo/v2"

type Object = tengo.Object

type UserFunction = tengo.UserFunction

//Modules 供脚本使用的模块信息
type Modules map[string]map[string]Object

//Option 配置选项
type Option func(*VM)

//WithCodeBlockMode 使用代码块模式，预编译脚本调用call时执行脚本
func WithCodeBlockMode() Option {
	return func(lm *VM) {
		lm.mode = CodeBlockMode
	}
}

//WithMainFuncMode 使用main函数调用模式，启动时执行脚本，调用call时执行main函数
func WithMainFuncMode() Option {
	return func(lm *VM) {
		lm.mode = MainFuncMode
	}
}

//WithModule 添加模块
func WithModule(name string, exports map[string]Object) Option {
	return func(lm *VM) {
		lm.modules[name] = exports
	}
}
