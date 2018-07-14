package _select

import (
	"fmt"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int,1)
	string_chan :=make(chan string,1)

	int_chan <-1
	string_chan <-"hellow"
	//这里会随机选一个，还不是全部执行
	select {
		case value:=<-int_chan:
			fmt.Println(value)
		case value:=<-string_chan:
			fmt.Println(value)
	}

}
