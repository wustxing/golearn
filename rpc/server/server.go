package server

import (
	"fmt"
	"github.com/0990/golearn/rpc"
	"github.com/golang/protobuf/proto"
)

type ServerType int

const (
	_ ServerType = iota
	Gate
	Center
	Game
	Users
)

type Server interface {
	Send(proto.Message)
	Request(proto.Message, func(proto.Message, error)) error
	Call(proto.Message) (proto.Message, error)
	Route(sesid int32, msg proto.Message)
}

type server struct {
	rpcClient   *rpc.RPCClient
	serverid    int32
	serverTopic string //目标服务器nats的topic,暂为服务器id
}

func NewServer(client *rpc.RPCClient, serverid int32) Server {
	return &server{
		rpcClient:   client,
		serverid:    serverid,
		serverTopic: fmt.Sprintf("%v", serverid),
	}
}

func (p *server) Send(msg proto.Message) {
	p.rpcClient.SendMsg(p.serverTopic, msg)
}

func (p *server) Request(msg proto.Message, f func(proto.Message, error)) error {
	p.rpcClient.Request(p.serverTopic, msg, f)
	return nil
}

func (p *server) Call(msg proto.Message) (proto.Message, error) {
	return p.rpcClient.Call(p.serverTopic, msg)
}

func (p *server) Route(sesid int32, msg proto.Message) {
	p.rpcClient.Route2Server(p.serverTopic, sesid, msg)
}
