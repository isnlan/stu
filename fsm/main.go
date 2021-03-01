package main

import (
	"fmt"

	"github.com/looplab/fsm"
)

type Door struct {
	To  string
	FSM *fsm.FSM
}

func NewDoor(to string) *Door {
	d := &Door{
		To: to,
	}

	d.FSM = fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: ""},
		},
		fsm.Callbacks{
			//"enter_state": func(e *fsm.Event) { d.enterState(e) },
			//"enter_open": func(e *fsm.Event) {
			//	fmt.Println("enter open", e)
			//},
			//"leave_open": func(e *fsm.Event) {
			//	fmt.Println("leave open", e)
			//},
			"open": func(event *fsm.Event) {
				fmt.Println("openning")
			},
			"close": func(event *fsm.Event) {
				fmt.Println("close111")
			},
		},
	)

	return d
}

func (d *Door) enterState(e *fsm.Event) {
	fmt.Printf("The door to %s is %s\n", d.To, e.Dst)
}

func main() {
	door := NewDoor("heaven")
	fmt.Println("---")
	err := door.FSM.Event("open")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("---")

	err = door.FSM.Event("close")
	if err != nil {
		fmt.Println(err)
	}
}
