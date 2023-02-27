package fsm

//EventDesc 事件说明
type EventDesc struct {
	// Name 事件名称
	Name string

	//Src 原状态
	Src []string

	//Dst 目标状态
	Dst string
}

// Events 事件组
type Events []EventDesc

// Callbacks 事件调用函数
type Callbacks map[string]Callback
