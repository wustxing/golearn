package main

import (
	"fmt"
)

func main() {

	defer_call()
	fmt.Println("helloworld")
}

func defer_call(){
	defer func(){
		fmt.Println("A")
	}()
	defer func(){
		fmt.Println("B")
	}()
	defer func(){
		fmt.Println("C")
	}()
	panic("xujialong")
	fmt.Println("defercall")
}
