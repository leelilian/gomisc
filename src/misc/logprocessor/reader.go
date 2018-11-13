package main

import (
	"fmt"
	"math/rand"
	"time"
)

type CustomReader interface {
	Read()
}

type CustomFileReader struct {
	Path        string
	ReadChannel chan []byte
}

func (lf *CustomFileReader) Read() {

	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	for {

		var result []byte
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < 18; i++ {
			result = append(result, bytes[r.Intn(len(bytes))])
		}

		fmt.Printf("generated:%s\n", result)
		lf.ReadChannel <- result
	}

}
