package main

import (
	"bytes"
	"encoding/binary"
	"github.com/imroc/biu"
	"strconv"
)

type dnsHeader struct {
	id      []byte //标识字段，客户端会解析服务器返回的DNS应答报文，获取ID值与请求报文设置的ID值做比较，如果相同，则认为是同一个DNS会话。
	qr      int64  //0表示查询报文，1表示响应报文;
	opcode  int64  //0（标准查询），其他值为1（反向查询）和2（服务器状态请求）,[3,15]保留值;
	aa      int64  //表示授权回答（authoritative answer）-- 这个比特位在应答的时候才有意义，指出给出应答的服务器是查询域名的授权解析服务器;
	tc      int64  //表示可截断的（truncated）--用来指出报文比允许的长度还要长，导致被截断;
	rd      int64  //表示期望递归(Recursion Desired) -- 这个比特位被请求设置，应答的时候使用的相同的值返回。如果设置了RD，就建议域名服务器进行递归解析，递归查询的支持是可选的;
	ra      int64  //表示支持递归(Recursion Available) -- 这个比特位在应答中设置或取消，用来代表服务器是否支持递归查询;
	z       int64  //保留值，暂未使用;
	rcode   int64  //0 : 没有错误;1 : 报文格式错误(Format error);2 : 服务器失败(Server failure);3 : 域名不存在(Name Error);4 : 域名服务器不支持查询类型(Not Implemented);5 : 拒绝(Refused)
	qdcount int64  //报文请求段中的问题记录数
	ancount int64  //报文回答段中的回答记录数
	nscount int64  //报文授权段中的授权记录数
	arcount int64  //报文附加段中的附加记录数
}

type dnsQuestion struct {
	qname  []byte //域名
	qtype  []byte //查询的协议类型
	qclass []byte //查询的类,比如，IN代表Internet
}

type dnsAnswer struct {
	aname  []byte //域名
	atype  []byte //查询的协议类型
	aclass []byte //查询的类,比如，IN代表Internet
	ttl    int64  //time to live存活时间
	rdlen  int64  //数据长度
	rdata  []byte //数据记录
}

type dnsRespone struct {
	header   dnsHeader
	question dnsQuestion
	answer   dnsAnswer
}

var qtypeDic = map[int64]string{
	1:   "A",     //IPv4地址
	2:   "NS",    //名字服务器
	5:   "CNAME", //规范名称定义主机的正式名字的别名
	6:   "SOA",   //开始授权标记一个区的开始
	11:  "WKS",   //熟知服务定义主机提供的网络服务
	12:  "PTR",   //指针把IP地址转化为域名
	13:  "HINFO", //主机信息给出主机使用的硬件和操作系统的表述
	15:  "MX",    //邮件交换把邮件改变路由送到邮件服务器
	28:  "AAAA",  //IPv6地址
	252: "AXFR",  //传送整个区的请求
	255: "ANY",   //对所有记录的请求
}

func getHeader(data []byte) dnsHeader {
	flags := data[2:4]

	flagsStr := biu.BytesToBinaryString(flags)
	//log.Printf("flags:%x  flagsBit: %q ",flags,flagsStr)

	header := &dnsHeader{
		id:      data[0:2],
		qr:      StringToInt64(flagsStr[1:2]),
		opcode:  StringToInt64(flagsStr[2:6]),
		aa:      StringToInt64(flagsStr[6:7]),
		tc:      StringToInt64(flagsStr[7:8]),
		rd:      StringToInt64(flagsStr[8:9]),
		ra:      StringToInt64(flagsStr[10:11]),
		z:       StringToInt64(flagsStr[11:14]),
		rcode:   StringToInt64(flagsStr[14:18]),
		qdcount: BytesToInt64(data[4:6]),
		ancount: BytesToInt64(data[6:8]),
		nscount: BytesToInt64(data[6:8]),
		arcount: BytesToInt64(data[10:12]),
	}

	return *header
}

func getQuestion(data []byte) dnsQuestion {

	qnameBytes := data[12 : len(data)-4]

	question := &dnsQuestion{
		qname:  qnameBytes,
		qtype:  data[len(data)-4 : len(data)-2], //BytesToInt64(data[len(data)-4:len(data)-2]),
		qclass: data[len(data)-2 : len(data)],   //BytesToInt64(data[len(data)-2:len(data)]),
	}

	return *question
}

//int64转 byte 数组
func Int64ToBytes(i int64) []byte {
	b_buf := bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.BigEndian, i)
	return b_buf.Bytes()[len(b_buf.Bytes())-2:]
}

//byte 数组转 int64
func BytesToInt64(buf []byte) int64 {
	bufStr := biu.BytesToBinaryString(buf)
	bufint64, _ := strconv.ParseInt(bufStr[len(bufStr)-5:len(bufStr)-1], 10, 64)
	return bufint64
}

//字符串转 int64
func StringToInt64(str string) int64 {
	intStr, _ := strconv.ParseInt(str, 10, 64)
	return intStr
}

func main() {

}
