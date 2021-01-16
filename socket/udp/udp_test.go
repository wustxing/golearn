package udp

import (
	"fmt"
	"net"
	"strconv"
	"testing"
	"time"
)

func TestUDP(t *testing.T){
	go startServer()
	startClient()
}

func TestTime(t *testing.T){
	fmt.Println(time.Now().UnixNano())
}

func startServer(){
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
		Zone: "",
	})

	if err != nil {
		panic(err)
	}
	var receiveCount int32
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println(err)
			break
		}

		receiveCount++
		//fmt.Printf("addr:%v data:%v count:%v\n", addr, string(data[:n]), n)
		fmt.Println("server receiveCount:", receiveCount)
		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Println(err)
			continue
		}

		//go func() {
		//	for i:=0;i<1000;i++{
		//		_, err = listen.WriteToUDP([]byte("received success!"), addr)
		//		if err != nil {
		//			fmt.Println(err)
		//			continue
		//		}
		//		time.Sleep(time.Millisecond*100)
		//	}
		//}()
	}
}

func startClient(){
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 9090,
		Zone: "",
	})

	if err != nil {
		panic(err)
	}

	go func(){
		var receiveCount int32
		var latencyTotal int64
		for{
			result := make([]byte, 4096)
			n, _, err := conn.ReadFromUDP(result)
			if err != nil {
				fmt.Println(err)
			}
			receiveCount++
			s:=string(result[:n])
			t,_:=strconv.ParseInt(s,10,64)
			latencyTotal+=(time.Now().UnixNano()-t)
			fmt.Printf("client receiveCount:%d latency:%dms \n", receiveCount,latencyTotal/(int64(receiveCount)*(1000*1000)))
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
	time.Sleep(time.Second*5)
}
