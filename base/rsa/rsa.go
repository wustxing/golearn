package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func RSAEncrypt(publicKey []byte,data []byte)([]byte,error){
	block,_:=pem.Decode(publicKey)
	if block==nil{
		return nil,errors.New("public key error")
	}

	pub,err:=x509.ParsePKIXPublicKey(block.Bytes)
	if err!=nil{
		return nil,err
	}

	p:=pub.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader,p,data)
}

func RSADecrypt(privateKey []byte,cipher []byte,)([]byte,error){
	block,_:=pem.Decode(privateKey)
	if block==nil{
		return nil,errors.New("private key error!")
	}

	priv,err:=x509.ParsePKCS1PrivateKey(block.Bytes)
	if err!=nil{
		return nil,err
	}

	return rsa.DecryptPKCS1v15(rand.Reader,priv,cipher)
}

