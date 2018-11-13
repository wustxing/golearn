package main

import (
	"fmt"
	"github.com/0990/golearn/util"
)

type PropPack struct {
	PropID int32
	Num    int32
}

func main() {

	fmt.Println(util.GetRandomString(6))
}

func PropToString(propList []PropPack) (s string) {
	for i, v := range propList {
		s += fmt.Sprintf("%d,%d", v.PropID, v.Num)
		if i < len(propList) {
			s += ";"
		}
	}
	return
}
