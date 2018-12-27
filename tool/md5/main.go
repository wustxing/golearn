package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/0990/golearn/util"
	"os"
	"strings"
)

var (
	files = flag.String("f", "hello.1txt", "md5 file")
)

func main() {
	flag.Parse()
	fs := strings.Split(*files, ",")

	md5, err := util.Md5Files(fs...)
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	fmt.Println(md5)
	bufio.NewScanner(os.Stdin).Scan()
}
