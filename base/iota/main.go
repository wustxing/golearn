package main

import "fmt"

const (
	a = -2
	b = iota << 1 //b=0
	c             //c=1
	d
	e
	f
	g = iota
)

func main() {
	fmt.Println(b, c, d, e, f, g)
}
