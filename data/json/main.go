package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int32
	Name string
}

func main() {
	var data map[string]interface{}
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	str, _ := json.Marshal(mapD)
	fmt.Println(str)
	json.Unmarshal(str, &data)
	fmt.Print(data)

	u := &User{
		ID:   1,
		Name: "hello",
	}

	str, _ = json.Marshal(u)
	fmt.Println(string(str))
}
