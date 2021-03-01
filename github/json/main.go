package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const data = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		panic(err)
	}
	//age := m["age"]
	fmt.Println(reflect.TypeOf(m["age"]))
}
