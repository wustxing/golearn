package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, err := regexp.MatchString("^100", "10011111")
	fmt.Println(match, err)
}
