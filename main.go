package main

import "fmt"

type Helloer interface {
	SayHello()
}

type Hier interface {
	SayHi()
}

type Student struct {
}

func (s *Student) SayHello() {
	fmt.Println("hello")
}

func (s *Student) SayHi() {
	fmt.Println("hi")
}

func (s *Student) Name() string {
	return "xujialong"
}

func main() {
	s := new(Student)
	FromHelloToHi(s)

}

func FromHelloToHi(h Helloer) {
	n := h.(interface {
		Name() string
	}).Name()
	fmt.Println(n)
}
