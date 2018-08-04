package main

import (
	"fmt"
)

func main() {
	a := 1
	intArr := []*int{&a, nil}
	for _, v := range intArr {

		fmt.Println(v)
	}
	for i := 0; i < 2; i++ {
		switch i {
		case 0:
			fmt.Println("0")
			break
		case 1:
			fmt.Println("1")
		}
	}

}
