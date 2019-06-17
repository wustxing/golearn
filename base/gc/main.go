package main

import (
	"fmt"
	"time"
)

const (
	windowSize = 200000
	msgCount   = 1000000
)

type (
	message []byte
	buffer  [windowSize]message
)

var worst time.Duration
var total time.Duration

//func mkMessage(n int) message {
//	m := make(message, 1024)
//	//for i := range m {
//	//	//	m[i] = byte(n)
//	//	//}
//	return m
//}

func pushMsg(b *buffer, highID int) {
	start := time.Now()
	m := make(message, 1024)
	(*b)[highID%windowSize] = m
	elapsed := time.Since(start)
	if elapsed > worst {
		worst = elapsed
	}
	total += elapsed
}

func main() {
	var b buffer
	for i := 0; i < msgCount; i++ {
		pushMsg(&b, i)
	}
	fmt.Println("Worst push time: ", worst)
	fmt.Println("Average push time: ", total/msgCount)
}
