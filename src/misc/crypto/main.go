package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	
	err := GeneratePrivateKey(4096)
	if err != nil {
		fmt.Printf("error message:%v", err)
	}
	
	pub, err := readKey("public.pem")
	if err != nil {
		fmt.Printf("loading public key error:%v", err)
	}
	
	content := []byte("this is a test message,老子写着玩")
	result, err := encrypt(content, pub)
	if err != nil {
		fmt.Printf("encrypt  error:%v", err)
	}
	
	file, err := os.Create("encrypt")
	if err != nil {
		fmt.Printf("create file error:%v", err)
	}
	defer file.Close()
	
	encrypted := base64.StdEncoding.EncodeToString(result)
	
	length, err := file.WriteString(encrypted)
	if err != nil {
		fmt.Printf("write file error:%v", err)
	}
	if len(result) != length {
		fmt.Printf("write file error:%v", err)
	}
	
	fmt.Printf("%s\n", encrypted)
	
	private, err := readKey("private.pem")
	if err != nil {
		fmt.Printf("loading private key error:%v", err)
	}
	
	raw, err := decrypt(result, private)
	
	if err != nil {
		fmt.Printf("decrypt error:%v\n", err)
		
	}
	
	fmt.Printf("%s\n", raw)
	
	salt := []byte("fuck")
	
	signature, err := sign(content, private, salt)
	if err != nil {
		fmt.Printf("sign error:%v\n", err)
	}
	succ, err := verify(content, signature, pub, salt)
	fmt.Printf("verify result:%v\n", succ)
	fmt.Printf("verify error:%v\n", err)
	
}

func GeneratePrivateKey(bits int) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	
	stream := x509.MarshalPKCS1PrivateKey(privateKey)
	
	block := &pem.Block{
		Type:  "RSA Private Key",
		Bytes: stream,
	}
	
	privateFile, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	defer privateFile.Close()
	
	err = pem.Encode(privateFile, block)
	if err != nil {
		return err
	}
	
	publicKey := privateKey.PublicKey
	
	// publicKey := public.(*rsa.PublicKey)
	pubstream := x509.MarshalPKCS1PublicKey(&publicKey)
	
	block = &pem.Block{
		Type:  "RSA Public Key",
		Bytes: pubstream,
	}
	
	publicFile, err := os.Create("public.pem")
	if err != nil {
		return err
	}
	defer publicFile.Close()
	
	err = pem.Encode(publicFile, block)
	if err != nil {
		return err
	}
	
	return nil
	
}

func readKey(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
	
}

func encrypt(src []byte, publickey []byte) ([]byte, error) {
	
	block, _ := pem.Decode(publickey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	
	return rsa.EncryptPKCS1v15(rand.Reader, pub, src)
	
}

func decrypt(src []byte, privatekey []byte) ([]byte, error) {
	
	block, _ := pem.Decode(privatekey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	
	return rsa.DecryptPKCS1v15(rand.Reader, private, src)
	
}

func sign(src []byte, privatekey []byte, salt []byte) ([]byte, error) {
	
	block, _ := pem.Decode(privatekey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	
	/*
	hash := sha256.New()
	hash.Write(src)
	data := hash.Sum(salt)*/
	data := sha256.Sum256(append(src, salt...))
	
	return rsa.SignPKCS1v15(rand.Reader, private, crypto.SHA256, data[:])
	
}
func verify(src []byte, signature []byte, publickey []byte, salt []byte) (bool, error) {
	
	block, _ := pem.Decode(publickey)
	if block == nil {
		return false, errors.New("public key error")
	}
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return false, err
	}
	/*
	hash := sha256.New()
	hash.Write(src)
	data := hash.Sum(salt)*/
	
	data := sha256.Sum256(append(src, salt...))
	
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, data[:], signature)
	if err != nil {
		return false, err
	}
	return true, nil
}
