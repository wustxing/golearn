package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func main() {
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	privateKeyOut, _ := os.Create("privateKey")
	pem.Encode(privateKeyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	privateKeyOut.Close()

	publicKeyOut, _ := os.Create("publicKey")
	pem.Encode(publicKeyOut, &pem.Block{Type: "RSA Public KEY", Bytes: x509.MarshalPKCS1PublicKey(&pk.PublicKey)})
	publicKeyOut.Close()
}
