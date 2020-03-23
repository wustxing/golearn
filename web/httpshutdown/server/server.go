package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	server    *http.Server
	serverMux *http.ServeMux
}

func (p *Server) Init() {
	p.serverMux = http.NewServeMux()
	p.serverMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//writer.Write([]byte("wait 20"))
		log.Println("handle start")
		time.Sleep(time.Second * 20)
		//writer.Write([]byte("bye bye"))
		fmt.Fprint(writer, "bye")
		log.Println("handle end")
	})

	p.server = &http.Server{Addr: ":9090", Handler: p.serverMux}
}

func (p *Server) Run() {
	err := p.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func (p *Server) ShutDown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	p.server.SetKeepAlivesEnabled(false)

	err := p.server.Shutdown(ctx)
	return err
}
