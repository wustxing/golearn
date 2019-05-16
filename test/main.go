package main

import "fmt"

type I interface {
	Get() int
	Set(int)
}

//2
type S struct {
	Age int
}

func (s S) Get() int {
	return s.Age
}

func (s S) Set(age int) {
	s.Age = age
}

//3
func f(i I) {
	i.Set(10)
	fmt.Println(i.Get())
}

func main() {
	a := make([]int32, 0, 10)
	a = append(a, 1)
	a = append(a, 2)
	fmt.Println(a[1:])
}
