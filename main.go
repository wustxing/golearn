package main

import (
	"fmt"
	"github.com/0990/golearn/util"
)

func main() {
	path := "hello.txt"
	fmt.Println(util.Md5File(path))
	fmt.Println(util.Md5Files(path))
}
