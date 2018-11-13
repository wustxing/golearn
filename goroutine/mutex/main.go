package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex

var global_int int

func main() {
	m = new(sync.RWMutex)

	go read(1)
	go read(2)
	time.Sleep(5 * time.Second)

	lock := sync.Mutex{}
	lock.Lock()
}

func read(i int) {
	println(i, "read start")
	m.RLock()
	global_int = i
	println(i, global_int, "reading")
	time.Sleep(2 * time.Second)
	m.RUnlock()

	println(i, "read over")
}
