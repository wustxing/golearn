package main

import "fmt"

type teacher struct{
	people
}
type people struct{
}

func(p *people)ShowA(){
	fmt.Println("show a")
	p.ShowB()
}
func (p *people)ShowB(){
	fmt.Println("people show b")
}
func(t *teacher)ShowB(){
	fmt.Println("teacher show b")
}

func main(){
	teacher:=teacher{}
	teacher.ShowB()
	teacher.ShowA()
}
