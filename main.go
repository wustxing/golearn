package main

import "fmt"

type parent struct {
	parenter parenter
}

func (p *parent) turnOffLight() {
	fmt.Println("turn off light1")
}
func (p *parent) goToBed() {
	fmt.Println("go to bed")
}

func (p *parent) toSleep() {
	p.parenter.turnOffLight()
	p.parenter.goToBed()
}

type parenter interface {
	turnOffLight()
	goToBed()

	toSleep()
}

type child struct {
	parent
}

func (p *child) turnOffLight() {
	fmt.Println("turn off light child")
}
func (p *child) toSleep() {
	p.turnOffLight()
	p.goToBed()
}

func main() {
	child := &child{}
	child.parenter = child
	var people parenter = child
	people.toSleep()
}
