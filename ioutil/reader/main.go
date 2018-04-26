package main

import (
	"bufio"
	"strings"
	"fmt"
	"time"
)

func main(){
	reader:=bufio.NewReader(strings.NewReader("http;//hellowomrld\tIt is the home"))

	go Peek(reader)
	time.Sleep(time.Microsecond*200)
	go reader.ReadBytes('\t')
	time.Sleep(time.Second*100)
}

func Peek(reader *bufio.Reader){
	line,_:=reader.Peek(2)
	fmt.Printf("%s\n",line)
	time.Sleep(time.Second)
	fmt.Printf("%s\n",line)
}
