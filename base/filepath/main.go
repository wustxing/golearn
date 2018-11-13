package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	abs, err := filepath.Abs("hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(abs)

	dir := filepath.Dir(abs)
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)

	fmt.Println(filepath.Join(dir, "world"))
}
