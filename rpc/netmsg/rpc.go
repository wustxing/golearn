package netmsg

import (
	"github.com/0990/golearn/rpc/msg"
	"github.com/gogo/protobuf/proto"
)

func Marshal(message proto.Message, seqid int32) []byte {
	rpc := &msg.RPC{}
	rpc.Data, _ = proto.Marshal(message)
	rpc.Seqid = seqid
	//TODO msgid
	rpc.Msgid = 0
	data, _ := proto.Marshal(rpc)
	return data
}

func Unmarshal(data []byte, message proto.Message) (seqid, msgid int32, err error) {
	respRPC := &msg.RespRPC{}
	err = proto.Unmarshal(data, respRPC)
	if err != nil {
		return 0, 0, err
	}
	err = proto.Unmarshal(respRPC.Data, message)
	if err != nil {
		return 0, 0, err
	}
	return respRPC.Seqid, respRPC.Msgid, nil
}

func Unmarshal(data []byte) (seqid, msgid int32, data []byte, err error) {
	respRPC := &msg.RespRPC{}
	err = proto.Unmarshal(data, respRPC)
	if err != nil {
		return 0, 0, err
	}
	err = proto.Unmarshal(respRPC.Data, message)
	if err != nil {
		return 0, 0, err
	}
	return respRPC.Seqid, respRPC.Msgid, nil
}
