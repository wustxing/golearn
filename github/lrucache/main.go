package main

import (
	"fmt"
	"github.com/hashicorp/golang-lru"
)

func main() {
	cache, _ := lru.New(128)
	for i := 0; i < 256; i++ {
		cache.Add(i, i)
	}

	if cache.Len() != 128 {
		panic("bad len")
	}

	fmt.Println(cache.Get(120))
	fmt.Println(cache.Get(255))
}
