package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	var data map[string]interface{}
	mapD :=map[string]int{"apple":5,"lettuce":7}
	str,_:=json.Marshal(mapD)
	fmt.Println(str)
	json.Unmarshal(str,&data)
	fmt.Print(data)

}
