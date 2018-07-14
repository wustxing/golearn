package main

import (
	"github.com/0990/golearn/util/snowflake"
	"fmt"
)

func main(){
	snowflake,_:=snowflake.NewSnowflake(0)

	for i:=0;i<100;i++{
		key :=snowflake.Generate()
		fmt.Printf("%b,%d\n",key,key)
	}

}
