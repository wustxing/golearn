package main

import (
	"flag"
	"fmt"
	"github.com/0990/golearn/util"
	"strings"
	"time"
)

var (
	files = flag.String("f", "hello1.txt,hello.txt,hello2.txt,hello3.txt,emailsvr,emailsvr1", "md5 file")
)

func main() {
	flag.Parse()
	fs := strings.Split(*files, ",")
	now := time.Now()
	defer func() {
		duration := time.Since(now)
		fmt.Println(duration)
	}()
	md5, err := util.MD5FileSync1(fs...)
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	fmt.Println(md5)
}
