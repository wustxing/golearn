package main

type User struct {
	age int
}

type Users []*User

func (users Users) SayHello() []*User {
	return users
}

func main() {
	users := []*User{}
	Users(users).SayHello()
}
