package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.Int("p", 8888, "listen port")

func main() {
	flag.Parse()

	listenAddr := fmt.Sprintf(":%d", *port)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("")))

	s := http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
