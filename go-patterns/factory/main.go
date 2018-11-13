package main

import (
	"github.com/0990/golearn/go-patterns/factory/data"
	"log"
)

func main() {
	s := data.NewStore(data.DiskStorage)
	f, err := s.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	f.Write([]byte("data"))
	defer f.Close()
}
