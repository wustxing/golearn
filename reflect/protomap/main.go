package main

import (
	"reflect"
	"github.com/0990/golearn/reflect/protomap/msg"
	"fmt"
	"github.com/golang/protobuf/proto"
	"errors"
	"unsafe"
	"strconv"
)

type MessageInfo struct{
	msgType reflect.Type
	msgHandler MessageHandler
}

type MessageHandler func(msgid uint32,msg interface{})
var(
	msg_map = make(map[uint32]MessageInfo)
)

func MessageHandler_test(msgid uint32,msg interface{}){
	p:=msg.(*hello.Hello)
	fmt.Println(p.Name,p.Id)
}

func main(){
	RegisterMessage(&hello.Hello{},MessageHandler_test)

	hello:=&hello.Hello{
		Name:"xujialong",
	}
	data,_:=proto.Marshal(hello)

	msgID:=GetHash(reflect.TypeOf(hello).String())
	int36Str :=strconv.FormatUint(uint64(msgID),36)
	dataHead:=[]byte(int36Str)
	for i:=len(int36Str);i<7;i++{
		dataHead = append([]byte{'0'},dataHead...)
	}
	fmt.Println(string(dataHead),"msgID:",msgID)

	sendData :=make([]byte,0)
	sendData = append([]byte(dataHead),data...)
	sendDataContent :=sendData[7:]
	sendDataHead:=sendData[:7]
	sendmsgid,_:=strconv.ParseUint(string(sendDataHead),36,32)
	fmt.Println(len(sendData),":",len(data),",getDatalen:",len(sendDataContent),"sendDataHead",len(sendDataHead),"msgid",sendmsgid)
	HandleRawData(uint32(sendmsgid),data)
}


func RegisterMessage(msg interface{},handler MessageHandler){
	var info MessageInfo
	info.msgType = reflect.TypeOf(msg.(proto.Message))
	info.msgHandler = handler
	msgid:=GetHash(info.msgType.String())
	msg_map[msgid] = info
}

func HandleRawData(msgid uint32,data []byte)error{
	if info,ok:=msg_map[msgid];ok{
		msg:=reflect.New(info.msgType.Elem()).Interface()
		err:=proto.Unmarshal(data,msg.(proto.Message))
		if err!=nil{
			return err
		}
		info.msgHandler(msgid,msg)
		return err
	}
	fmt.Println("not found msgid")
	return errors.New("not found msgid")
}

const (
	c1_32 uint32 = 0xcc9e2d51
	c2_32 uint32 = 0x1b873593
)

// GetHash returns a murmur32 hash for the data slice.
func GetHash(str string) uint32 {
	data :=[]byte(str)
	// Seed is set to 37, same as C# version of emitter
	var h1 uint32 = 37

	nblocks := len(data) / 4
	var p uintptr
	if len(data) > 0 {
		p = uintptr(unsafe.Pointer(&data[0]))
	}

	p1 := p + uintptr(4*nblocks)
	for ; p < p1; p += 4 {
		k1 := *(*uint32)(unsafe.Pointer(p))

		k1 *= c1_32
		k1 = (k1 << 15) | (k1 >> 17) // rotl32(k1, 15)
		k1 *= c2_32

		h1 ^= k1
		h1 = (h1 << 13) | (h1 >> 19) // rotl32(h1, 13)
		h1 = h1*5 + 0xe6546b64
	}

	tail := data[nblocks*4:]

	var k1 uint32
	switch len(tail) & 3 {
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0])
		k1 *= c1_32
		k1 = (k1 << 15) | (k1 >> 17) // rotl32(k1, 15)
		k1 *= c2_32
		h1 ^= k1
	}

	h1 ^= uint32(len(data))

	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16

	return (h1 << 24) | (((h1 >> 8) << 16) & 0xFF0000) | (((h1 >> 16) << 8) & 0xFF00) | (h1 >> 24)
}




