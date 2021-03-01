package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://oapi.dingtalk.com/robot/send?access_token=7065c78a340a83a69ec4d2803fffcd810cc9f2ba9eda31cb0596828ba2b0a5ee"

func main() {
	data := `
{
    "msgtype": "text", 
    "text": {
        "content": "我就是我, 是不一样的烟火@156xxxx8827"
    }, 
    "at": {
        "atMobiles": [
            "156xxxx8827", 
            "189xxxx8325"
        ], 
        "isAtAll": false
    }
}
`

	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(data)))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	data1,err:=ioutil.ReadAll(resp.Body)
	fmt.Println(string(data1))
}
