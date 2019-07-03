package rpc

import (
	"fmt"
	"github.com/0990/golearn/rpc/msg"
	"github.com/0990/golearn/rpc/netmsg"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
)

var ErrTimeOut = errors.New("rpc: timeout")
var ErrNoKnow = errors.New("rpc: unknow")

type RPCClient struct {
	conn        *nats.Conn
	clientTopic string
	serverid    int32
}

func NewRPCClient(serverid int32) (*RPCClient, error) {
	p := &RPCClient{}
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}
	p.conn = conn
	p.serverid = serverid
	p.clientTopic = fmt.Sprintf("%v", serverid)
	return p, nil
}

//阻塞式
func (p *RPCClient) Call(serverTopic string, message proto.Message) (proto.Message, error) {
	ret := make(chan proto.Message)
	call := CreateCall(func(msg proto.Message, err error) {
		ret <- msg
	})

	data := netmsg.Marshal(message, call.seqid)

	err := p.send(serverTopic, data)
	if err != nil {
		return nil, err
	}

	select {
	case result, ok := <-ret:
		if !ok {
			return nil, errors.New("client closed")
		}
		return result, nil
	case <-time.After(time.Second * 10):
		GetCallWithDel(call.seqid)
		return nil, ErrTimeOut
	}

	return nil, ErrNoKnow
}

//非阻塞式
func (p *RPCClient) Request(serverTopic string, message proto.Message, onRecv func(proto.Message, error)) error {
	call := CreateCall(onRecv)
	data := netmsg.Marshal(message, call.seqid)

	err := p.send(serverTopic, data)
	if err != nil {
		return err
	}

	time.AfterFunc(time.Second*10, func() {
		//TODO 放在主线程中工作
		if call, ok := GetCallWithDel(call.seqid); ok {
			call.onRecv(nil, ErrTimeOut)
		}
	})

	return nil
}

//仅发送
func (p *RPCClient) SendMsg(serverTopic string, msg proto.Message) {
	data := netmsg.Marshal(msg, 0)
	p.conn.Publish(serverTopic, data)
}

func (p *RPCClient) Run() {
	go p.ReadLoop()
}

func (p *RPCClient) ReadLoop() error {
	sub, err := p.conn.SubscribeSync(p.clientTopic)
	if err != nil {
		return err
	}

	for {
		m, err := sub.NextMsg(time.Minute)
		if err != nil && err == nats.ErrTimeout {
			continue
		} else if err != nil {
			return err
		}
		rpc := &msg.RPC{}
		err = proto.Unmarshal(m.Data, rpc)
		if err != nil {
			logrus.WithError(err)
		}
		switch rpc.Type {
		case msg.RPC_Request:
			//rpc request
			//TODO unmarshal 构造func(sender,fun)回调
		case msg.RPC_Response:
			//rpc response
			if v, ok := GetCallWithDel(seqid); ok {
				v.onRecv(rpc, nil)
			}
		case msg.RPC_RouteServer:
			//route server
			//TODO unmarshal 构造func(session,fun)回调
		case msg.RPC_RouteGate:
			//route gate
			//TODO 根据sesid发送到对应客户端
		case msg.RPC_Normal:
			//normal
			//TODO unmarshal 构造func(sender,fun)回调

		default:
			logrus.WithField("rpcType", rpc.Type).Error("not support rpc type")
		}
	}
}

func (p *RPCClient) send(topic string, data []byte) error {
	return p.conn.Publish(topic, data)
}

func (p *RPCClient) Route2Server(topic string, sesid int32, message proto.Message) {
	data, _ := proto.Marshal(message)
	rpc := &msg.RPC{
		Type:     msg.RPC_RouteServer,
		Sesid:    sesid,
		Senderid: p.serverid,
		Msgid:    1,
		Data:     data,
	}

	rpcData, _ := proto.Marshal(rpc)
	p.send(topic, rpcData)
}
