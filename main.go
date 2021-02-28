package main

import (
	"encoding/json"
	"fmt"
)

const data = `{"id":1,"obj":{"name":"xu","age":18}}`

func main() {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(data), &m)
	fmt.Println(m, err)

	var d []byte
	for _, v := range m {
		switch v.(type) {
		case map[string]interface{}:
			d, _ = json.Marshal(v)
		case []interface{}:
			d, _ = json.Marshal(v)
		default:

		}
	}

	fmt.Println(string(d))
}
