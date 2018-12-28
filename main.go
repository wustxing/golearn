package main

import (
	"fmt"
	"github.com/0990/golearn/util"
)

func main() {
	path := "hello.txt"
	path1 := "hello1.txt"
	fmt.Println(util.Md5File(path))
	fmt.Println(util.Md5FileAsync(path, path1))
	fmt.Println(util.MD5FileSync(path, path1))
}
