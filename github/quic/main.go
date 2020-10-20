package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/lucas-clemente/quic-go"
	"io"
	"log"
	"math/big"
	"time"
)

func main(){
	addr:="127.0.0.1:9999"
	go func(){
		log.Fatal(startServer(addr))
	}()
	err:=startClient(addr)
	if err!=nil{
		panic(err)
	}
	fmt.Println(err)
}

type loggingWriter struct{ io.Writer }

func (w loggingWriter) Write(b []byte) (int, error) {
	fmt.Printf("Server: Got '%s'\n", string(b))
	return w.Writer.Write(b)
}

func startServer(addr string)error{
	listener,err:=quic.ListenAddr(addr,generateTLSConfig(),nil)
	if err!=nil{
		return err
	}
	for{
		sess,err:=listener.Accept(context.Background())
		if err!=nil{
			return err
		}
		go func(){
			for{
				stream,err:=sess.AcceptStream(context.Background())
				if err!=nil{
					fmt.Println(err)
					return
				}
				fmt.Println("stream open:",stream.StreamID())
				go func(){
					_,err=io.Copy(loggingWriter{stream},stream)
					fmt.Println(err)
				}()
			}
		}()
	}
}

func startClient(addr string)error{
	tlsConf:=&tls.Config{
		InsecureSkipVerify: true,
		NextProtos: []string{"quic-echo-example"},
	}
	session,err:=quic.DialAddr(addr,tlsConf,nil)
	if err!=nil{
		return err
	}
	go func(){
		stream,err:=session.OpenStreamSync(context.Background())
		if err!=nil{
			panic(err)
		}

		_,err = stream.Write([]byte("hi"))
		if err!=nil{
		    fmt.Println(err)
			return
		}
	}()
	stream,err:=session.OpenStreamSync(context.Background())
	if err!=nil{
		return err
	}
	message:="hello world!!!"
	fmt.Printf("Client:Sending:`%s`\n",message)
	_,err = stream.Write([]byte(message))
	if err!=nil{
		return err
	}

	fmt.Println(session.Context().Err())
	session.CloseWithError(1000,"close test")
	time.Sleep(time.Second)
	fmt.Println(session.Context().Err())


	buf:=make([]byte,len(message))
	_,err = io.ReadFull(stream,buf)
	if err!=nil{
		return err
	}
	fmt.Printf("Client:Got `%s`\n",buf)
	return nil
}



func generateTLSConfig()*tls.Config{
	key,err:=rsa.GenerateKey(rand.Reader,1024)
	if err!=nil{
		panic(err)
	}
	template:=x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER,err:=x509.CreateCertificate(rand.Reader,&template,&template,&key.PublicKey,key)
	if err!=nil{
		panic(err)
	}
	keyPEM:=pem.EncodeToMemory(&pem.Block{
		Type:    "RSA PRIVATE KEY",
		Bytes:   x509.MarshalPKCS1PrivateKey(key),
	})
	certPEM:=pem.EncodeToMemory(&pem.Block{
		Type:    "CERTIFICATE",
		Headers: nil,
		Bytes:   certDER,
	})
	tlsCert,err:=tls.X509KeyPair(certPEM,keyPEM)
	if err!=nil{
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos: []string{"quic-echo-example"},
	}
}


