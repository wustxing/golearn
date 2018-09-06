package main

import (
	"flag"
	"fmt"
)

var username = flag.String("name","xujialong1","Input your name")
var age = flag.String("age","6","Input your age")
func main(){
	flag.Parse()
	fmt.Println("Hello",*username,*age)
	flag.Usage()
}
