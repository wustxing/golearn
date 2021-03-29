package main

import (
	"fmt"
)

const data = `{"id":1,"obj":{"name":"xu","age":18}}`

type User interface {
	Id() int32
}

type user struct{
	id int32
}

func(p *user)Id()int32{
	return p.id
}

func delUser(id int32, users []*user) {
	for i := 0; i < len(users); {
		if users[i].Id() == id {
			users = append(users[:i], users[i+1:]...)
		} else {
			i++
		}
	}
	fmt.Println(users)
}

func main() {
	var a interface{}
	a = 1
	if a==0{
		fmt.Println("0")
	}
	if a==1{
		fmt.Println("1")
	}
}
