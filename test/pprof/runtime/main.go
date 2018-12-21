package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sync"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file.")
)

func main() {
	log.Println("begin")
	flag.Parse()
	f, err := os.Create("cpuprofile.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 30; i++ {
		nums := fibonacci(i)
		fmt.Println(nums)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go print(&wg)
	wg.Wait()
}

func print(group *sync.WaitGroup) {
	func() {
		for i := 0; i < 10000; i++ {
			fmt.Println(i)
		}
		group.Done()
	}()
}

func fibonacci(num int) int {
	if num < 2 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}
