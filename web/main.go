package main

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

var mux map[string]func(w http.ResponseWriter, r *http.Request)

func main() {
	server := http.Server{
		Addr:        ":8081",
		Handler:     &myhandle{},
		ReadTimeout: 5 * time.Second,
	}
	mux = make(map[string]func(w http.ResponseWriter, r *http.Request))

	mux["/"] = bytes.Index
}

type myHandle struct{}

func (*myHandle) ServerHttp(w http.ResponseWriter, r *http.Request) {
	log.Println("请求url:", r.URL.String())
	log.Println("请求方法:", r.Method)

	r.ParseForm()
	log.Println("请求报文:", r)
	log.Println("请求的参数:", r.Form)

	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
	} else {
		fileTer(w, r)
	}
}

type BaseJsonBean struct {
	Code    int         `json:"Code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
