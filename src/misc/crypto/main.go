package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {

	err := GeneratePrivateKey(4096)
	if err != nil {
		fmt.Printf("error message:%v", err)
	}
}

func GeneratePrivateKey(bits int) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	stream := x509.MarshalPKCS1PrivateKey(privateKey)

	block := &pem.Block{
		Type:  "private key",
		Bytes: stream,
	}

	/*
		file, err := os.Create("private.pem")
		if err != nil {
			return err
		}*/

	err = pem.Encode(os.Stdout, block)
	if err != nil {
		return err
	}
	return nil

}
