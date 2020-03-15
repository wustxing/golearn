package main

import (
	"fmt"
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
	v := r.FormValue("order")
	fmt.Println(v)
	fmt.Println("request url:", r.Host)
	fmt.Fprintln(w, "hello world")
}
