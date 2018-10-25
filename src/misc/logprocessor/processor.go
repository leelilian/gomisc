package main

import (
	"strings"
)

type CustomProcessor interface {
	Process()
}

type CustomFileProcessor struct {
	InputChannel  chan []byte
	OutputChannel chan []byte
}

func (cp *CustomFileProcessor) Process() {

	for {
		line := string(<-cp.InputChannel)
		line = strings.ToUpper(line)
		cp.OutputChannel <- []byte(line)
	}
}
