package fsm

import (
	"context"
	"fmt"
)

// Event is the info that get passed as a reference in the callbacks.
type Event struct {
	// FSM is an reference to the current FSM.
	FSM *FSM

	// Event is the event name.
	Event string

	// Src is the state before the transition.
	Src string

	// Dst is the state after the transition.
	Dst string

	// Err is an optional error that can be returned from a callback.
	Err error

	// Args is an optional list of arguments passed to the callback.
	Args []interface{}

	// canceled is an internal flag set if the transition is canceled.
	canceled bool

	// async is an internal flag set if the transition should be asynchronous
	async bool

	// cancelFunc is called in case the event is canceled.
	cancelFunc func()
}

// Cancel can be called in before_<EVENT> or leave_<STATE> to cancel the
// current transition before it happens. It takes an optional error, which will
// overwrite e.Err if set before.
func (e *Event) Cancel(err ...error) {
	e.canceled = true
	e.cancelFunc()

	if len(err) > 0 {
		e.Err = fmt.Errorf("已取消:%w", err[0])
	}
}

// Async can be called in leave_<STATE> to do an asynchronous state transition.
//
// The current state transition will be on hold in the old state until a final
// call to Transition is made. This will complete the transition and possibly
// call the other callbacks.
func (e *Event) Async() {
	e.async = true
}

type Callback func(context.Context, *Event)

type eKey struct {
	// event is the name of the event that the keys refers to.
	event string

	// src is the source from where the event can transition.
	src string
}

// cKey is a struct key used for keeping the callbacks mapped to a target.
type cKey struct {
	// target is either the name of a state or an event depending on which
	// callback type the key refers to. It can also be "" for a non-targeted
	// callback like before_event.
	target string

	// callbackType is the situation when the callback will be run.
	callbackType int
}
