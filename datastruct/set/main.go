package main

import (
	"gopkg.in/fatih/set.v0"
	"fmt"
)

func main(){
	s1:=set.New()
	s1.Add(1)
	s1.Add(2)

	s2:=set.New()
	s2.Add(2)
	s2.Add(3)

	diff:=set.Difference(s1,s2)
	fmt.Println(diff)
}
