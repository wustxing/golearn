package tcp

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"testing"
	"time"
)

const TargetPort = 2000

func TestTCP(t *testing.T){
	go startServer()
	time.Sleep(time.Second)
	startClient()
}

func startServer(){
	listen, err := net.Listen("tcp", "0.0.0.0:9999")
		if err != nil {
			fmt.Println("listen error:", err)
			return
		}

		for {
			conn, err := listen.Accept()
			if err != nil {
				fmt.Println("accept error:", err)
				break
			}
			go HandleConn(conn)
		}
}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 100)
	var receiveCount int32
	var writeElapse int64
	for {
		//read
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		receiveCount++
		before:=time.Now().UnixNano()
		_,err=conn.Write(buffer[0:n])
		if err!=nil{
			fmt.Println(err)
		}
		writeElapse+=(time.Now().UnixNano()-before)
		fmt.Printf("receive count:%d,writeelapse:%d \n", receiveCount,writeElapse/1000000)
	}
}


func startClient(){
	conn, err := net.Dial("tcp", ":1000")
	if err != nil {
		fmt.Println("dial error", err)
	}

	go func(){
		var receiveCount int32
		var latencyTotal int64
		for{
			result := make([]byte, 19)
		//	n, err := conn.Read(result)
			_,err = io.ReadFull(conn,result)
			if err != nil {
				fmt.Println(err)
			}
			receiveCount++
			s:=string(result)
			t,err:=strconv.ParseInt(s,10,64)
			if err!=nil{
				fmt.Println(err)
			}
			latency:=(time.Now().UnixNano()-t)
			latencyTotal+=latency
			fmt.Printf("client receiveCount:%d currLatency:%d,averLatency:%dms \n", receiveCount,latency/1000000,latencyTotal/(int64(receiveCount)*(1000*1000)))
		}
	}()

	for i:=0;i<=1000;i++{
		t:=time.Now().UnixNano()
		s:=fmt.Sprintf("%v",t)
		_, err = conn.Write([]byte(s))
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond*1)
	}
	time.Sleep(time.Second*10)

	////	defer conn.Close()
	//buf := make([]byte, 100)
	//conn.SetReadDeadline(time.Now().Add(time.Second))
	//_, err = conn.Read(buf)
	//if err != nil {
	//	e, ok := err.(net.Error)
	//	fmt.Println(err, e, ok, e.Temporary(), e.Timeout())
	//}
	//conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	//_, err = conn.Read(buf)
	//if err != nil {
	//	e, ok := err.(net.Error)
	//	fmt.Println(err, e, ok, e.Temporary(), e.Timeout())
	//}
}