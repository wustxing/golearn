package client

import (
	"net"
	"time"
)

type Client struct {
	conn net.Conn
}

func New(addr string) (*Client, error) {
	conn, err := net.DialTimeout("tcp", addr, 3*time.Second)
	if err != nil {
		return nil, err
	}

	client := &Client{
		conn: conn,
	}

	return client, nil
}

func (c *Client) Send(message string) ( error) {
	_, err := c.conn.Write([]byte(message))
	return err

	//reader := bufio.NewReader(c.conn)
	//resp, err := reader.ReadString('\n')
	//if err != nil {
	//	return "", err
	//}
	//return resp, err
}
