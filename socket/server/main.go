package main

import (
	"net"
	"fmt"
)

func main(){
	listen,err:=net.Listen("tcp",":8888")
	if err!=nil{
		fmt.Println("listen error:",err)
		return
	}

	for{
		conn,err:=listen.Accept()
		if err!=nil{
			fmt.Println("accept error:",err)
			break;
		}
		go HandleConn(conn)
	}

}

func HandleConn(conn net.Conn){
	defer conn.Close()
	for{
		//read

		//write
	}
}
