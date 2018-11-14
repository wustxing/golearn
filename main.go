package main

import "fmt"

type PropPack struct {
	maps map[int32]int
}

func main() {
	p := PropPack{}
	p.maps = make(map[int32]int)
	a := 1
	p.maps[1] = a

	b := 2
	s := p
	s.maps[1] = b
	fmt.Println(p, s, p.maps[1])
}
