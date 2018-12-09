package main

import (
	"fmt"
	"sort"
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
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Age > persons[j].Age
	})
	fmt.Printf("after sort:%v", persons)
}
