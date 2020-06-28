package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abcde_145_190"
	index := strings.LastIndex(str, "_")
	if index > 0 {
		fmt.Println(str[index+1:])
	}

	fmt.Println(strings.ReplaceAll("abc_%d_cde_%d", "%d", "*"))
}
