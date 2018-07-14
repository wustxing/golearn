package main

import (
	"fmt"
	"time"
)

type Hi struct{
	id int32
}

var HiChan chan Hi

func SayHi(hi *Hi){
	fmt.Println(hi)
	hi.id = 3
	fmt.Println(hi)
}

func main(){
	HiChan = make(chan Hi,10)
	hi1:=&Hi{
		2,
	}
	HiChan<-*hi1
	HiChan<-*hi1
	hi1.id = 3

	go func(){
		for{
			select{
			case hi:=<-HiChan:
				fmt.Println(hi)
			}
		}
	}()

	time.Sleep(5*time.Second)

}
