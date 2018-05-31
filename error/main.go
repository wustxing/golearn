package main

import (
	"errors"
)

func main(){
	//var err error
	check := func(f func()error) {
		err:=f()
		if err != nil {
			panic(err)
		}
	}
	check(func() error{
		return errors.New("hello")
		//fmt.Println(err)
	})
}
