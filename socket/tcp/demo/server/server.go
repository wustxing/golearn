package server

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	conn   net.Conn
	Server *server
}

type server struct {
	address                  string
	onNewClientCallBack      func(c *Client)
	onClientConnectionClosed func(c *Client, err error)
	onNewMessage             func(c *Client, message string)
}

func (c *Client) listen() {
	reader := bufio.NewReader(c.conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			c.conn.Close()
			c.Server.onClientConnectionClosed(c, err)
			if opErr, ok := err.(*net.OpError); ok {
				fmt.Println(opErr)
			}
			return
		}
		c.Server.onNewMessage(c, message)
	}
}

func (c *Client) Send(message string) error {
	_, err := c.conn.Write([]byte(message))
	return err
}

func (c *Client) SendBytes(b []byte) error {
	_, err := c.conn.Write(b)
	return err
}

func (c *Client) Conn() net.Conn {
	return c.conn
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (s *server) OnNewClient(callback func(c *Client)) {
	s.onNewClientCallBack = callback
}

func (s *server) OnClientConnectionClosed(callback func(c *Client, err error)) {
	s.onClientConnectionClosed = callback
}

func (s *server) OnNewMessage(callback func(c *Client, message string)) {
	s.onNewMessage = callback
}

func (s *server) Listen() {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		fmt.Errorf("address:%v,%v", s.address, err)
	}
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		fmt.Println("accept new connect",conn.LocalAddr(),conn.RemoteAddr())
		client := &Client{
			conn:   conn,
			Server: s,
		}
		go client.listen()
		s.onNewClientCallBack(client)
	}
}

func New(address string) *server {
	fmt.Println("creating server with address", address)
	server := &server{
		address: address,
	}

	server.OnNewMessage(func(c *Client, message string) {})
	server.OnNewClient(func(c *Client) {})
	server.OnClientConnectionClosed(func(c *Client, err error) {})
	return server
}
