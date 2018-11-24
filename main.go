package main

import (
	"fmt"
)

type Config struct {
	name string
}

var c *Config

func main() {
	a := 1
	m := 2

	b := &a
	c := b
	fmt.Println(b, c)
	b = &m
	fmt.Println(a, b, c)
}
