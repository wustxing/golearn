package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"
)

func TestRSA(t *testing.T){
	pri,pub,err:=GenRSAKey(1024)
	if err!=nil{
		t.Fatal(err)
	}
	fmt.Printf("public key:\n%s\n",string(pub))
	fmt.Printf("private key:\n%s\n",string(pri))

	data:=[]byte("hello 0990")
	chipher,err:=RSAEncrypt(pub,data)
	if err!=nil{
		t.Fatal(err)
	}

	out,err:=RSADecrypt(pri,chipher)
	if err!=nil{
		t.Fatal(err)
	}

	if !bytes.Equal(data,out){
		t.FailNow()
	}
}

func GenRSAKey(bits int)(private []byte,public []byte,err error){
	privateKey,err:=rsa.GenerateKey(rand.Reader,bits)
	if err!=nil{
		return nil,nil,err
	}


	derSteam:=x509.MarshalPKCS1PrivateKey(privateKey)
	priBlock :=&pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   derSteam,
	}

	publicKey:=&privateKey.PublicKey
	derPkix,err :=x509.MarshalPKIXPublicKey(publicKey)
	if err!=nil{
		return nil,nil,err
	}

	publicBlock:=&pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   derPkix,
	}

	return pem.EncodeToMemory(priBlock),pem.EncodeToMemory(publicBlock),nil
}
