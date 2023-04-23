package fsm

import (
	"context"
	"fmt"
	"strings"
	"sync"
)

type FSM struct {
	// current 当前状态
	current string

	//currentResult 当前结果
	currentResult interface{}

	transitions map[eKey]string

	// callbacks 回调函数
	callbacks map[cKey]Callback

	transition      func()
	transitionerObj transitioner

	stateMu sync.RWMutex
	eventMu sync.Mutex

	metadataMu sync.RWMutex
	metadata   map[string]interface{}
}

func NewFSM(initial string, events []EventDesc, callbacks map[string]Callback) *FSM {
	f := &FSM{
		current:         initial,
		transitions:     make(map[eKey]string),
		callbacks:       make(map[cKey]Callback),
		metadata:        make(map[string]interface{}),
		transitionerObj: &transitionerStruct{},
	}

	allEvents := make(map[string]bool) //所有事件
	allStates := make(map[string]bool) //所有状态
	for _, e := range events {
		for _, src := range e.Src {
			f.transitions[eKey{e.Name, src}] = e.Dst
			allStates[src] = true
			allStates[e.Dst] = true
		}
		allEvents[e.Name] = true
	}

	//处理回调事件
	for name, fn := range callbacks {
		var target string
		var callbackType int

		switch {
		case strings.HasPrefix(name, "before_"):
			target = strings.TrimPrefix(name, "before_")
			if target == "event" {
				target = ""
				callbackType = callbackBeforeEvent
			} else if _, ok := allEvents[target]; ok {
				callbackType = callbackBeforeEvent
			}
		case strings.HasPrefix(name, "leave_"):
			target = strings.TrimPrefix(name, "leave_")
			if target == "state" {
				target = ""
				callbackType = callbackLeaveState
			} else if _, ok := allStates[target]; ok {
				callbackType = callbackLeaveState
			}
		case strings.HasPrefix(name, "enter_"):
			target = strings.TrimPrefix(name, "enter_")
			if target == "state" {
				target = ""
				callbackType = callbackEnterState
			} else if _, ok := allStates[target]; ok {
				callbackType = callbackEnterState
			}
		case strings.HasPrefix(name, "after_"):
			target = strings.TrimPrefix(name, "after_")
			if target == "event" {
				target = ""
				callbackType = callbackAfterEvent
			} else if _, ok := allEvents[target]; ok {
				callbackType = callbackAfterEvent
			}
		default:
			target = name
			if _, ok := allStates[target]; ok {
				callbackType = callbackEnterState
			} else if _, ok := allEvents[target]; ok {
				callbackType = callbackAfterEvent
			}
		}

		if callbackType != callbackNone {
			f.callbacks[cKey{target, callbackType}] = fn
		}
	}

	return f
}

// Current 返回当前状态
func (f *FSM) Current() string {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()
	return f.current
}

// Is 与当前状态相同则返回true
func (f *FSM) Is(state string) bool {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()
	return state == f.current
}

// SetState 修改当前状态，不会触发任何事件回调
func (f *FSM) SetState(state string) {
	f.stateMu.Lock()
	defer f.stateMu.Unlock()
	f.current = state
}

// Can 判断当前事件是否可以触发
func (f *FSM) Can(event string) bool {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()
	_, ok := f.transitions[eKey{event, f.current}]
	return ok && (f.transition == nil)
}

// SetResult 设置当前事件执行结果
func (f *FSM) SetResult(result interface{}) {
	f.currentResult = result
}

// GetResult 获取当前事件执行结果
func (f *FSM) GetResult() interface{} {
	return f.currentResult
}

// AvailableTransitions 返回当前状态可执行的事件列表
func (f *FSM) AvailableTransitions() []string {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()
	var transitions []string
	for key := range f.transitions {
		if key.src == f.current {
			transitions = append(transitions, key.event)
		}
	}
	return transitions
}

// Cannot 判断当前事件是否为不能触发
func (f *FSM) Cannot(event string) bool {
	return !f.Can(event)
}

// Metadata 获取元数据
func (f *FSM) Metadata(key string) (interface{}, bool) {
	f.metadataMu.RLock()
	defer f.metadataMu.RUnlock()
	dataElement, ok := f.metadata[key]
	return dataElement, ok
}

// SetMetadata 设置元数据
func (f *FSM) SetMetadata(key string, dataValue interface{}) {
	f.metadataMu.Lock()
	defer f.metadataMu.Unlock()
	f.metadata[key] = dataValue
}

// DeleteMetadata 删除元数据
func (f *FSM) DeleteMetadata(key string) {
	f.metadataMu.Lock()
	delete(f.metadata, key)
	f.metadataMu.Unlock()
}

// Event 事件转化调用
func (f *FSM) Event(ctx context.Context, event string, args ...interface{}) error {
	f.eventMu.Lock()
	var unlocked bool
	defer func() {
		if !unlocked {
			f.eventMu.Unlock()
		}
	}()

	f.stateMu.RLock()
	defer f.stateMu.RUnlock()

	//当事件正在执行
	if f.transition != nil {
		return InTransitionError{event}
	}

	//获取事件
	dst, ok := f.transitions[eKey{event, f.current}]
	if !ok {
		for ekey := range f.transitions {
			if ekey.event == event {
				return InvalidEventError{event, f.current}
			}
		}
		return UnknownEventError{event}
	}

	//处理事件取消
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	e := &Event{f, event, f.current, dst, nil, args, false, false, cancel}

	//执行预处理函数
	err := f.beforeEventCallbacks(ctx, e)
	if err != nil {
		return err
	}

	//没有处理函数
	if f.current == dst {
		f.afterEventCallbacks(ctx, e)
		return NoTransitionError{fmt.Errorf("%s %w", f.current, e.Err)}
	}

	//创建处理函数
	transitionFunc := func(ctx context.Context, async bool) func() {
		return func() {
			if ctx.Err() != nil {
				if e.Err == nil {
					e.Err = ctx.Err()
				}
				return
			}

			f.stateMu.Lock()
			f.current = dst
			f.transition = nil // treat the state transition as done
			f.stateMu.Unlock()

			if !async {
				f.eventMu.Unlock()
				unlocked = true
			}
			f.enterStateCallbacks(ctx, e)
			f.afterEventCallbacks(ctx, e)

		}
	}

	f.transition = transitionFunc(ctx, false)

	if err = f.leaveStateCallbacks(ctx, e); err != nil {
		if _, ok := err.(CanceledError); ok {
			f.transition = nil
		} else if asyncError, ok := err.(AsyncError); ok {
			ctx, cancel := uncancelContext(ctx)
			e.cancelFunc = cancel
			asyncError.Ctx = ctx
			asyncError.CancelTransition = cancel
			f.transition = transitionFunc(ctx, true)
			return asyncError
		}
		return err
	}

	// Perform the rest of the transition, if not asynchronous.
	f.stateMu.RUnlock()
	defer f.stateMu.RLock()
	err = f.doTransition()
	if err != nil {
		return InternalError{}
	}

	return e.Err
}

// Transition wraps transitioner.transition.
func (f *FSM) Transition() error {
	f.eventMu.Lock()
	defer f.eventMu.Unlock()
	return f.doTransition()
}

// doTransition wraps transitioner.transition.
func (f *FSM) doTransition() error {
	return f.transitionerObj.transition(f)
}

// beforeEventCallbacks 执行 before_ 函数
func (f *FSM) beforeEventCallbacks(ctx context.Context, e *Event) error {

	//调用事件预处理
	if fn, ok := f.callbacks[cKey{e.Event, callbackBeforeEvent}]; ok {
		fn(ctx, e)
		if e.canceled {
			return CanceledError{e.Err}
		}
	}

	//调用全局预处理
	if fn, ok := f.callbacks[cKey{"", callbackBeforeEvent}]; ok {
		fn(ctx, e)
		if e.canceled {
			return CanceledError{e.Err}
		}
	}
	return nil
}

// leaveStateCallbacks 执行 leave_ 函数
func (f *FSM) leaveStateCallbacks(ctx context.Context, e *Event) error {
	//完成当前事件的leave_函数
	if fn, ok := f.callbacks[cKey{f.current, callbackLeaveState}]; ok {
		fn(ctx, e)
		if e.canceled {
			return CanceledError{e.Err}
		} else if e.async {
			return AsyncError{Err: e.Err}
		}
	}

	//处理系统leave_处理函数
	if fn, ok := f.callbacks[cKey{"", callbackLeaveState}]; ok {
		fn(ctx, e)
		if e.canceled {
			return CanceledError{e.Err}
		} else if e.async {
			return AsyncError{Err: e.Err}
		}
	}
	return nil
}

// enterStateCallbacks 处理 enter_ 函数
func (f *FSM) enterStateCallbacks(ctx context.Context, e *Event) {
	if fn, ok := f.callbacks[cKey{f.current, callbackEnterState}]; ok {
		fn(ctx, e)
	}
	if fn, ok := f.callbacks[cKey{"", callbackEnterState}]; ok {
		fn(ctx, e)
	}
}

// afterEventCallbacks 处理 after_ 函数
func (f *FSM) afterEventCallbacks(ctx context.Context, e *Event) {
	if fn, ok := f.callbacks[cKey{e.Event, callbackAfterEvent}]; ok {
		fn(ctx, e)
	}
	if fn, ok := f.callbacks[cKey{"", callbackAfterEvent}]; ok {
		fn(ctx, e)
	}
}
