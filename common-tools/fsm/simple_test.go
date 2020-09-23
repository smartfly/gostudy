package fsm

import (
	"fmt"
	"testing"

	"github.com/looplab/fsm"
)

/*
 * @author taohu
 * @date 2020/8/28
 * @time 上午10:10
 * @desc please describe function
**/

func TestBasicFunc(t *testing.T) {
	f := fsm.NewFSM("start", fsm.Events{
		{Name: "submit", Src: []string{"start"}, Dst: "checking"},
		{Name: "check-fail", Src: []string{"checking"}, Dst: "check-fail"},
		{Name: "check-suc", Src: []string{"checking"}, Dst: "check-suc"},
	}, fsm.Callbacks{})

	fmt.Println(f.Current())
	err := f.Event("submit")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f.Current())
	_ = f.Event("check-fail")
	fmt.Println(f.Current())
	err = f.Event("fail")
	if err != nil {
		fmt.Println(err)
	}
}

type Door struct {
	To  string
	FSM *fsm.FSM
}

func (d *Door) enterState(e *fsm.Event) {
	fmt.Printf("The door to %s is %s\n", d.To, e.Dst)
}

func NewDoor(to string) *Door {
	d := &Door{To: to}
	d.FSM = fsm.NewFSM("closed", fsm.Events{
		{Name: "open", Src: []string{"closed"}, Dst: "open"},
		{Name: "close", Src: []string{"open"}, Dst: "closed"},
	}, fsm.Callbacks{
		"enter_state": func(event *fsm.Event) {
			d.enterState(event)
		},
	})
	return d
}

func TestStructFunc(t *testing.T) {
	door := NewDoor("heaven")
	err := door.FSM.Event("open")
	if err != nil {
		fmt.Println(err)
	}
	err = door.FSM.Event("close")
	if err != nil {
		fmt.Println(err)
	}
}
