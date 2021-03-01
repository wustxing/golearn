package main

import (
	"fmt"
	"math"
)

func main(){
	var a int32 =math.MaxInt32-100
	new := int64(a) + int64(101)
	if new > math.MaxInt32 {
		new = math.MaxInt32
	}
	fmt.Println(new)
	a = int32(new)
	fmt.Println(a)
	fmt.Println(math.MaxInt32 )
}


