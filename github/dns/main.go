
package main

import (
	"encoding/hex"
	"fmt"
	"github.com/miekg/dns"
	"net"
)

func main() {
	//data:=[]byte{}

	data, err := hex.DecodeString("0001010000010000000000000a74787468696e6b696e6703636f6d0000010001")
	if err != nil {
		panic(err)
	}

	req:=new(dns.Msg)
	err=req.Unpack(data)
	if err!=nil{
		panic(err)
	}

	fmt.Println(req.Question)

	conn,err:=net.Dial("udp","8.8.8.8:53")
	if err!=nil{
		panic(err)
	}

	_,err = conn.Write(data)
	if err!=nil{
		panic(err)
	}

	b := make([]byte, 2048)
	n, err := conn.Read(b)
	if err != nil {
		panic(err)
	}

	resp:=new(dns.Msg)
	err=resp.Unpack(b[0:n])
	if err!=nil{
		panic(err)
	}

	fmt.Println(resp.Answer)


	//var dig dnsutil.Dig
	//dig.SetDNS("8.8.8.8") //or ns.xxx.com
	//a, err := dig.A("www.baidu.com")  // dig google.com @8.8.8.8
	//
	//for _,v:=range a{
	//	fmt.Println(v.String())
	//}
	////fmt.Println()
	//fmt.Println(a, err)
}