package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "b", "c"}

	s = s[:2]

	fmt.Println(s)
}
