package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
		handleConn(c)
	}
}

func handleConn(c net.Conn) {
	go ConnRead(c)
}

type HistoryReader struct {
	rd   io.Reader
	buff []byte
}

func (h *HistoryReader) Buff() []byte {
	return h.buff
}

func (h *HistoryReader) Read(p []byte) (int, error) {
	n, err := h.rd.Read(p)
	h.buff = append(h.buff, p[0:n]...)
	return n, err
}

func ConnRead(c net.Conn) {
	hReader := HistoryReader{
		rd: c,
	}
	reader := bufio.NewReader(&hReader)
	req, err := http.ReadRequest(reader)
	fmt.Println(string(hReader.Buff()))
	if err != nil {
		log.Println(err)
		return
	}

	host := req.Host
	if strings.Index(host, ":") == -1 {
		host += ":80"
	}
	s, err := net.DialTimeout("tcp", host, time.Second)
	if err != nil {
		c.Write([]byte("HTTP/1.1 404 Not found\r\n\r\n"))
		return
	}
	defer s.Close()
	if req.Method == http.MethodConnect {
		c.Write([]byte("HTTP/1.1 200 Connection established\r\n\r\n"))
	} else {
		s.Write(hReader.Buff())
	}
	go func() {
		io.Copy(c, s)
	}()
	io.Copy(s, c)
}
