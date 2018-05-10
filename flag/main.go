package main

import (
	"flag"
	"fmt"
)

func main(){
	username := flag.String("name","xujialong","Input your name")

	flag.Parse()

	fmt.Println("Hello",*username)
	flag.Usage()
}
