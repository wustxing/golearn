package main

import (
	"fmt"
	"strconv"
)

const (
	LCardCommon = 1 << iota
	LCardSpecial
	LCardEquip
)

func main() {
	str := fmt.Sprintf("%d", 101)
	i, err := strconv.ParseInt(str, 2, 32)
	if err != nil {
		panic(err)
	}
	println(LCardCommon, LCardSpecial, i&LCardCommon, i&LCardSpecial, i&LCardEquip)
}
