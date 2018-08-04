package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now(), time.Now().Unix())
	fmt.Println(time.Now().UTC(), time.Now().UTC().Unix())
}
