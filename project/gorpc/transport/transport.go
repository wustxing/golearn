package transport

import (
	"encoding/binary"
	"io"
	"net"
)

type Transport struct {
	conn net.Conn
}

func NewTransport(conn net.Conn) *Transport {
	return &Transport{conn}
}

func (t *Transport) Send(data []byte) error {
	buf := make([]byte, 4+len(data))
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))
	copy(buf[4:], data)
	_, err := t.conn.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

func (t *Transport) Read() ([]byte, error) {
	header := make([]byte, 4)
	_, err := io.ReadFull(t.conn, header)
	if err != nil {
		return nil, err
	}

	dataLen := binary.BigEndian.Uint32(header)
	data := make([]byte, dataLen)
	_, err = io.ReadFull(t.conn, data)
	if err != nil {
		return nil, err
	}

	return data, nil

}
