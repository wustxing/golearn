package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var port = flag.Int("port",9090,"listen port")

func main() {
	flag.Parse()

	listenAddr:=fmt.Sprintf(":%d",*port)
	fmt.Println(listenAddr)

	http.HandleFunc("/hello", ServerHandler)
	err := http.ListenAndServe(listenAddr, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for k,v:=range r.Form{
		fmt.Printf("k:%s,v:%s\n",k,v[0])
	}
	fmt.Println("method:", r.Method)

	data, err := ioutil.ReadAll(r.Body)
	if err == nil {
		fmt.Println("body:", string(data))
	}
	w.Write([]byte("hello"))
}
