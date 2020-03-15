package main

import (
	"fmt"
	"net/url"
)

func main() {
	ret, err := url.Parse("dm.baidu.com:5000")

	fmt.Println(ret, err)
}
