package main

import (
	"github.com/looplab/fsm"
	"fmt"
)

type Door struct{
	To string
	FSM *fsm.FSM
}

func(d *Door)enterState(e *fsm.Event){
	fmt.Printf("the door to %s is %s\n",d.To,e.Dst)
}
func(d *Door)leaveState(e *fsm.Event){
	fmt.Printf("leave state,the door to %s is %s\n",d.To,e.Dst)
}

func NewDoor(to string)*Door{
	d:=&Door{
		To:to,
	}
	d.FSM = fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name:"open",Src:[]string{"closed"},Dst:"open"},
			{Name:"close",Src:[]string{"open"},Dst:"closed"},
		},
		fsm.Callbacks{
			"enter_state":func(e *fsm.Event){d.enterState(e)},
			"leave_state":func(e *fsm.Event){d.leaveState(e)},
		},
	)
	return d
}

func main(){
	door:=NewDoor("heaven")
	err:=door.FSM.Event("open")

	if err!=nil{
		fmt.Println(err)
	}

	err=door.FSM.Event("close")
	if err!=nil{
		fmt.Println(err)
	}
}
