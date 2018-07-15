package main

import "fmt"

type GetID interface {
	GetID() int32
}

type Card struct{
	id int32
}

func(p *Card)GetID()int32{
	return p.id
}

func(p *Card)SetID(value int32){
	p.id = value
}
var getID GetID
func main(){
	getID:=Card{id:1}

	//getID.SetID(2)

	newCard:=getID

	newCard.SetID(2)
	fmt.Println(getID,newCard)
}
