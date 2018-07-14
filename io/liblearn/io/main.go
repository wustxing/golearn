package main

import (
	//"bufio"
	"fmt"
	"strings"
	//"log"
	"os"
	//	"strings"
)

func main() {

	reader := strings.NewReader("helloworld")
	reader.Seek(-2, os.SEEK_END)

	r, _, _ := reader.ReadRune()
	fmt.Printf("%c\n", r)

}
