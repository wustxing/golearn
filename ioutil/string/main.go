package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(strings.Count("fivevev","ve"))
	fmt.Printf("Fields are:%q",strings.FieldsFunc((" foo bar baz"),func(r rune)bool{
		if r=='b'{
			return true
		}
		return false
	}))

	fmt.Println(strings.Split(" foo bar baz","b"))
}
