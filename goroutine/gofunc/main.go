package main

import (
	"sync"
	"fmt"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup

	wg.Add(10)

	//for i:=0;i<10;i++{
	//	go func(){
	//		fmt.Println("i=",i)
	//		wg.Done()
	//	}()
	//}
	for i:=0;i<10;i++{
		fmt.Println(i)
		go func(i int){
			fmt.Println("i=",i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
