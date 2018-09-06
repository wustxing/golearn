package main

import (
	"fmt"
	"time"
)

func main() {
	//t1 := time.Now()
	//t2 := t1.Add(-1 * time.Minute)
	//
	//str1 := t1.Format("2006-01-02 15:04:05")
	//str2 := t2.Format("2006-01-02 15:04:05")
	//
	//t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-01-02 15:04:05", time.Local)
	//fmt.Println(t1, t1.Before(t2), str1, str2, t)

	timer := time.AfterFunc(time.Second, func() {

	})
	fmt.Println(timer)
	timer.Stop()
	if timer == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("non nil")
	}
	fmt.Println(timer)
}
