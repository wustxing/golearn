package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"reflect"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	value := gjson.Parse(json)
	value.ForEach(func(key, value gjson.Result) bool {
		v := value.Value()
		fmt.Println(reflect.TypeOf(v))
		fmt.Println(key, value.Value())
		return true
	})
	//value
	//value := gjson.Get(json, "name.last")
	//value.Value()
	println(value.Get("name").Type)
}
