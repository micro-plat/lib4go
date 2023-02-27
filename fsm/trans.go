package fsm

const (
	callbackNone int = iota
	callbackBeforeEvent
	callbackLeaveState
	callbackEnterState
	callbackAfterEvent
)

// transitioner is an interface for the FSM's transition function.
type transitioner interface {
	transition(*FSM) error
}

// transitionerStruct is the default implementation of the transitioner
// interface. Other implementations can be swapped in for testing.
type transitionerStruct struct{}

// Transition completes an asynchronous state change.
//
// The callback for leave_<STATE> must previously have called Async on its
// event to have initiated an asynchronous state transition.
func (t transitionerStruct) transition(f *FSM) error {
	if f.transition == nil {
		return NotInTransitionError{}
	}
	f.transition()
	return nil
}
