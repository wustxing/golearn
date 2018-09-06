package main

import "fmt"

type stateMachineFunc func() stateMachineFunc

var machine stateMachineFunc

func initStateMachine() stateMachineFunc{
	var state1, state2  stateMachineFunc

	state1 = func() stateMachineFunc {
		fmt.Println("enter state1 function")
		return state2
	}

	state2 = func() stateMachineFunc {
		fmt.Println("enter state2 function")
		return nil
	}

	return state1
}

func event() {
	m := machine()
	if m != nil {
		machine = m
	}
}

func main(){
	machine = initStateMachine();
	event();
	event();
	event();
}