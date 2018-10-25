package main

import (
	"fmt"
)

type LogProcessor struct {
	Reader    CustomReader
	Processor CustomProcessor
}

func main() {

	var reader CustomReader

	ch := make(chan []byte)
	reader = &CustomFileReader{Path: "", ReadChannel: ch}

	var pro CustomProcessor
	out := make(chan []byte)
	pro = &CustomFileProcessor{InputChannel: ch, OutputChannel: out}

	for j := 0; j < 5; j++ {
		go reader.Read()
	}

	for i := 0; i < 10; i++ {
		go pro.Process()
	}

	for {

		fmt.Printf("received:%s\n", <-out)
	}

}
