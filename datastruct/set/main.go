package main

import (
	"fmt"
	"gopkg.in/fatih/set.v0"
	"log"
)

func main() {
	s1 := set.New()
	s1.Add(1)
	s1.Add(2)
	s1.Add(2)

	s2 := set.New()
	s2.Add(2)
	s2.Add(3)

	diff := set.Difference(s1, s2)
	fmt.Println(s1, s2, diff)

	log.SetFlags()
	for _, v := range s1.List() {
		intValue := v.(int)
		fmt.Println(intValue)
	}
}
