package main

import "fmt"

func main() {
	hi := &hello{}
	hi.card = &Card1{id: 1}
	hi.hi = append(hi.hi, &Card1{id: 1})
	hi.hi = append(hi.hi, &Card1{id: 2})
	b := hi.GetSlice()
	//fmt.Println(hi, b)
	b[0].SetID(100)
	b = append(b[:1], b[2:]...)
	//judgeSelectList = append(judgeSelectList[:i], judgeSelectList[i+1:]...)
	fmt.Println(hi.hi[0].GetID(), hi.hi, b[0].GetID(), b)
}

type hello struct {
	hi   []Card
	card Card
}
type Card interface {
	GetID() int32
	SetID(int32)
}

type Card1 struct {
	id int32
}

func (c *Card1) GetID() int32 {
	return c.id
}

func (c *Card1) SetID(id int32) {
	c.id = id
}

func (h *hello) GetSlice() []Card {

	return h.hi
}

func SayHi(hi ...int) {
	fmt.Println(hi)
}
