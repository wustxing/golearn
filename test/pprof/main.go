package main

import (
	"github.com/0990/golearn/test/pprof/data"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("https://github.com/EDDYCJY"))
		}
	}()

	//see in http://127.0.0.1:6060/debug/pprof/
	//go tool pprof http://localhost:6060/debug/pprof/profile?seconds=60
	//go tool pprof http://localhost:6060/debug/pprof/heap
	//go tool pprof http://localhost:6060/debug/pprof/block
	//go tool pprof http://localhost:6060/debug/pprof/mutex
	http.ListenAndServe("0.0.0.0:6060", nil)
}
