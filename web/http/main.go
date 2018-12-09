package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("listen 8888 start...")
	http.HandleFunc("/", ServerHandler)
	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request url:", r.URL.String())
	fmt.Fprintln(w, "hello world")
}
