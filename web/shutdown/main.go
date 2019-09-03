package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var listenAddr = flag.String("listen-addr", ":5000", "server listen address")

func main() {
	flag.Parse()

	logger := log.New(os.Stdout, "http:", log.LstdFlags)

	//TODO 为什么长度为1
	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	server := newWebserver(logger)
	go gracefullShutdown(server, logger, quit, done)

	logger.Println("Server is ready to handle requests at", *listenAddr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s:%v\n", listenAddr, err)
	}

	<-done
	logger.Println("Server stoped")
}

func gracefullShutdown(server *http.Server, logger *log.Logger, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	logger.Println("Server is shutting down...")
	//给服务端最多30秒的关闭时间，如果服务端还没关好，强制结束整个服务
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Could not gracefully shutdown the serfver:%v\n", err)
	}
	logger.Println("Server shut down")
	close(done)
}

func newWebserver(logger *log.Logger) *http.Server {
	router := http.NewServeMux()

	router.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	return &http.Server{
		Addr:         *listenAddr,
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}
