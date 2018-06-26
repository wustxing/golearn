package main

import "fmt"
import (
	"time"
)

func main(){
	//runtime.GOMAXPROCS(4)
	const n = 100000
	starTime:=time.Now()
	leftMost := make(chan int)
	var right chan int
	left:=leftMost

	for i:=0;i<n;i++{
		right = make(chan int)
		go f(left,right)
		left = right
	}
	go func(c chan int){
		c<-1
	}(left)
	fmt.Println(<-leftMost)
	elapsed:=time.Since(starTime)
	fmt.Println(elapsed)
}

func f(left,right chan int){
	left<-1+<-right
}
