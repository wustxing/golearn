package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex


func main(){
	m = new(sync.RWMutex)

	go read(1)
	go read(2)
	time.Sleep(5*time.Second)
}

func read(i int){
	println(i,"read start")
	m.Lock()

	println(i,"reading")
	time.Sleep(1*time.Second)
	m.Unlock()

	println(i,"read over")
}
