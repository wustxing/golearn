package main

import (
	"fmt"
	"testing"
)

func Test_Sort(t *testing.T) {
	persons := PersonSlice{
		Person{
			Age: 1,
		},
		Person{
			Age: 5,
		},
		Person{
			Age: 2,
		},
	}
	persons[1], persons[2] = persons[2], persons[1]
	fmt.Printf("after sort:%v", persons)
}
