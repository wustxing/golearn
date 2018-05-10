package main

import (
	"strings"
	"fmt"
)

func main(){
	fields:=make(map[string]string)

	fields["hello"] = "m"
	fields["world"] = "j"

	fieldstring := make([]string,0,len(fields))

	for k,_ :=range fields{
		fieldstring = append(fieldstring, k);
	}

	fmt.Println(strings.Join(fieldstring,","))
}
