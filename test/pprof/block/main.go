package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

var mutex sync.Mutex

func main() {

	runtime.SetBlockProfileRate(1 * 1000 * 1000)
	var wg sync.WaitGroup
	wg.Add(1)
	mutex.Lock()

	go worker(&wg)
	fmt.Println("sleep 2 start")
	time.Sleep(2 * time.Millisecond)
	fmt.Println("sleep 2 end")
	mutex.Unlock()
	wg.Wait()
	writeProfTo("block", "block.bprof")

}

func worker(wg *sync.WaitGroup) {
	fmt.Println("worker start")
	defer wg.Done()
	mutex.Lock()
	time.Sleep(10 * time.Millisecond)
	mutex.Unlock()
	fmt.Println("worker end")
}

func writeProfTo(name, fn string) {
	p := pprof.Lookup(name)
	if p == nil {
		fmt.Errorf("%s prof not found", name)
		return
	}
	f, err := os.Create(fn)
	if err != nil {
		fmt.Errorf("%v", err.Error())
		return
	}
	defer f.Close()
	err = p.WriteTo(f, 0)
	if err != nil {
		fmt.Errorf("%v", err.Error())
		return
	}
}
