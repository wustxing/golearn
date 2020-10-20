package main

import (
	"flag"
	"fmt"
	"github.com/0990/socks5"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

var c = flag.String("c","","remoteIP:port(client mode)")
var ss5Addr = flag.String("ss5","","ss5 server addr(client mode)")
var port = flag.Int("port",3010,"listen port(server mode)")

const SEND_COUNT = 10

func main(){
	flag.Parse()

	if *c==""{
		err:=runServer(*port)
		if err!=nil{
			log.Fatal(err)
		}
	}else{
		runClient(*c,*ss5Addr)
	}
}

func runServer(listenPort int)error{
	lisAddr:=fmt.Sprintf(":%d",listenPort)

	log.Println("listenAddr:",lisAddr)

	lis,err:=net.Listen("tcp",lisAddr)
	if err!=nil{
		return err
	}

	for{
		conn,err:=lis.Accept()
		if err!=nil{
			return err
		}
		go func(conn2 net.Conn){
			defer conn.Close()

			buf:=make([]byte,19)
			for{
				_,err:=io.ReadFull(conn,buf)
				if err!=nil{
					return
				}
				s:=fmt.Sprintf("%s:%d",string(buf),time.Now().UnixNano())
				fmt.Println(s)
				go func(){
					conn.Write([]byte(s))
				}()
			}
		}(conn)
	}
}

type Dialer interface {
	Dial(network, addr string) (net.Conn, error)
}

func runClient(remoteAddr string,ss5Addr string){
	var dialer Dialer
	if ss5Addr!=""{
		dialer = socks5.NewClient(socks5.ClientCfg{
			ServerAddr: ss5Addr,
			UserName:   "",
			Password:   "",
		})
	}else{
		dialer = &net.Dialer{}
	}

	conn,err:=dialer.Dial("tcp",remoteAddr)
	if err!=nil{
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go read(conn,&wg)

	for i:=0;i< SEND_COUNT;i++{
		s := fmt.Sprintf("%d", time.Now().UnixNano())
		_, err := conn.Write([]byte(s))
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}

	wg.Wait()
}

func read(conn net.Conn,group *sync.WaitGroup){
	defer group.Done()

	buf:=make([]byte,39)
	var latencyTotal,sendLatencyTotal,recvLatencyTotal int64
	var count int64
	for{
		conn.SetReadDeadline(time.Now().Add(time.Second*2))
		_,err:=io.ReadFull(conn,buf)
		if err!=nil{
			break
		}
		s:=string(buf)
		t:=strings.Split(s,":")
		if len(t)!=2{
			break
		}
		count++
		t0,_:= strconv.Atoi(t[0])
		t1,_:= strconv.Atoi(t[1])
		sendT := int64(t0)
		recvT := int64(t1)
		now:=time.Now().UnixNano()
		latency:=(now-sendT)/1000000
		sendLatency :=(recvT-sendT)/1000000
		recvLatency :=(now-recvT)/1000000
		fmt.Printf("send:%d recv:%d latency:%d \n",sendLatency,recvLatency,latency)
		latencyTotal+=latency
		sendLatencyTotal+=sendLatency
		recvLatencyTotal+=recvLatency
		if count== SEND_COUNT {
			break
		}
	}

	averLatency:=latencyTotal/count
	averSendLatency:=sendLatencyTotal/count
	averRecvLatency:=recvLatencyTotal/count

	fmt.Println("average:")
	fmt.Printf("send:%d recv:%d latency:%d \n",averSendLatency,averRecvLatency,averLatency)
}
