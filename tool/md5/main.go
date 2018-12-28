package main

import (
	"flag"
	"fmt"
	"github.com/0990/golearn/util"
	"strings"
)

var (
	files = flag.String("f", "hello.1txt", "md5 file")
)

func main() {
	flag.Parse()
	fs := strings.Split(*files, ",")

	md5, err := util.MD5FileSync(fs...)
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	fmt.Println(md5)
}
