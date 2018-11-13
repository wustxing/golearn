package main

import (
	"fmt"
	"sanguosha.com/sgs_nuyan/gameutil"
)

type User struct {
	id int
}

func main() {
	resultMap := make(map[int]int32)
	for i := 0; i < 1000; i++ {
		resultMap[Random()]++
	}
	fmt.Println(resultMap)
}

func Random() int {
	randNum := gameutil.RandNum(2)
	intArr := []int32{1, 2}
	for i, v := range intArr {
		if v > randNum {
			return i
		}
	}
	return 100
}
