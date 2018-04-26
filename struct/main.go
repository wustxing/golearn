package main

import (
	"fmt"
)

type hello struct {
	world
	ll
}

type world struct {
	i string
}

type ll struct {
	j string
}

func (w *world) say() {
	fmt.Println("hello")
}
func (l *ll) hi() {
	fmt.Println("hi")
}

func main() {
	hello := hello{}
	hello.say()
	hello.hi()
}
