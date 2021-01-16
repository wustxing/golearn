package main

import (
	"fmt"
	"golang.org/x/net/icmp"
	"html/template"
	"net"
)

func main(){
	netaddr,err:=net.ResolveIPAddr("ip4","10.225.136.159")
	if err!=nil{
		panic(err)
	}
	conn,err:=net.ListenIP("ip4:icmp",netaddr)
	if err!=nil{
		panic(err)
	}
	for{
		buf:=make([]byte,1024)
		n,addr,err:=conn.ReadFrom(buf)
		fmt.Println(err)
		msg,_:=icmp.ParseMessage(1,buf[0:n])
		fmt.Println(n,addr,msg.Type,msg.Code,msg.Checksum)
	}
	template.HTMLEscapeString()
}
