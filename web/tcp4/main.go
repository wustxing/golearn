package main

import (
	"net"
	"net/http"
)

func main(){
	http.HandleFunc("/hello", ServerHandler)
	server := &http.Server{}
	l, err := net.Listen("tcp4", ":80")
	if err != nil {
		panic(err)
	}
	err = server.Serve(l)
}

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
