package main

import (
	"fmt"
	"github.com/davyxu/goobjfmt"
)

type MyData struct {
	Name string
	Age  int32
}

func main() {
	binaryOpt()
}

func binaryOpt() {
	input := &MyData{
		Name: "xujialong",
		Age:  12,
	}
	body, err := goobjfmt.BinaryWrite(input)
	if err != nil {
		panic("writeError")
	}
	x := &MyData{}
	err = goobjfmt.BinaryRead(body, x)
	if err != nil {
		panic("read error")
	}
	fmt.Printf("%v", x)
}
