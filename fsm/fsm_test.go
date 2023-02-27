package fsm

import (
	"context"
	"fmt"
	"testing"
)

func TestFsm(t *testing.T) {
	fsm := NewFSM(
		"init",
		Events{
			{Name: "open", Src: []string{"init"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		Callbacks{
			"enter_state": func(ctx context.Context, e *Event) {
				fmt.Println("enter_state:", e.Dst)
			},
			"open": func(ctx context.Context, e *Event) {
				fmt.Println("open:", e.Args)
				e.FSM.SetResult(map[string]interface{}{
					"a": "1",
				})
			},
			"close": func(ctx context.Context, e *Event) {
				fmt.Println("closed:", e.Args)
			},
		},
	)
	fmt.Println("start:", fsm.Current())
	if err := fsm.Event(context.Background(), "open"); err != nil {
		fmt.Println(err)
	}
	fmt.Println("end:", fsm.Current(), fsm.GetResult())

	t.Fail()

}
