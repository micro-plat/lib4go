package tengoscript

import (
	"github.com/d5/tengo/v2"
	"github.com/micro-plat/lib4go/types"
)

const (

	//CodeBlockMode 代码块模式
	CodeBlockMode int = 0

	//MainFuncMode main函数调用模式
	MainFuncMode int = 1
)

//VM lua虚拟机
type VM struct {
	script   *tengo.Script
	compiled *tengo.Compiled
	modules  map[string]map[string]Object
	mode     int
}

//New 构建虚拟机
func New(scope string, opts ...Option) (*VM, error) {
	vm := &VM{
		modules: make(map[string]map[string]Object),
	}
	for _, opt := range opts {
		opt(vm)
	}

	//加载脚本
	script := tengo.NewScript([]byte(scope))

	//加载模块
	modules := tengo.NewModuleMap()
	for k, v := range vm.modules {
		modules.AddBuiltinModule(k, v)
	}
	script.SetImports(modules)

	//编译脚本
	compiled, err := script.Compile()
	if err != nil {
		return nil, err
	}
	vm.compiled = compiled
	return vm, nil
}

//Run 执行脚本
func (v *VM) Run(input ...interface{}) (types.XMap, error) {
	mp := types.NewXMap()
	mp.Append(input...)
	script := v.compiled.Clone()
	for k, value := range mp {
		if err := script.Set(k, value); err != nil {
			return nil, err
		}
	}
	if err := script.Run(); err != nil {
		return nil, err
	}
	return var2Mpa(script), nil
}
