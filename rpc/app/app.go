package app

import (
	"github.com/0990/golearn/rpc"
	"github.com/0990/golearn/rpc/server"
	"github.com/0990/golearn/rpc/service"
	"github.com/golang/protobuf/proto"
	"sync"
)

type app struct {
	sync.Mutex
	sid2server map[int32]server.Server
	rpcClient  *rpc.RPCClient
	serverid   int32
	worker     service.Worker
}

func NewApp(serverid int32) (*app, error) {
	rpcClient, err := rpc.NewRPCClient(serverid)
	if err != nil {
		return nil, err
	}
	p := new(app)
	p.worker = service.NewWorker()
	p.serverid = serverid
	p.rpcClient = rpcClient
	return p, nil
}

func (p *app) Run() {
	p.rpcClient.Run()
	p.worker.Run()
}

func (p *app) RegisterMsg(msg proto.Message, f func(sesid int32, msg proto.Message)) {

}

func (p *app) RegisterRequest(msg proto.Message, f func(serverid int32, msg proto.Message)) {

}

func (p *app) GetServerById(serverid int32) server.Server {
	p.Lock()
	defer p.Unlock()
	if v, ok := p.sid2server[serverid]; ok {
		return v
	}
	s := server.NewServer(p.rpcClient, serverid)
	p.sid2server[serverid] = s
	return s
}

//TODO add
func (p *app) GetServerByType(serverType server.ServerType) server.Server {
	return nil
}
