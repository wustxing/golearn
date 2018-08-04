package main

import "fmt"

type IP []int32

func (ip IP) Change() {
	ip[0] = 2
	fmt.Println(ip)
}

func main() {
	ip := IP([]int32{1, 2, 3, 4})
	ip.Change()
	fmt.Println(ip)
}
