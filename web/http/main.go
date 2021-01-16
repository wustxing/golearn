package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("listen 9999 start...")
	http.HandleFunc("/", ServerHandler)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	data, err := ioutil.ReadAll(r.Body)
	if err == nil {
		fmt.Println("body:", string(data))
	}
	w.Write([]byte("hello"))
}
