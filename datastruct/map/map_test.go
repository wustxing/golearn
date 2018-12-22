package main

import (
	"fmt"
	"testing"
)

type Student struct {
	age int
}

func TestMap(t *testing.T) {
	testMap := make(map[int32]Student)
	testMap[1] = Student{1}

	s := testMap[1]
	s.age = 2
	s1 := testMap[1]
	s1.age = 3
	for _, v := range testMap {
		v.age = 4
	}
	//testMap[1].age = 4
	fmt.Printf("%p,%p,%v", &s, &s1, testMap[1])
}
