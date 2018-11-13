package main

type LogProcessor struct {
	Reader    CustomReader
	Processor CustomProcessor
	Writer    CustomWriter
}

func main() {

	var reader CustomReader

	ch := make(chan []byte, 64)
	reader = &CustomFileReader{Path: "", ReadChannel: ch}

	var pro CustomProcessor
	out := make(chan []byte, 4096)
	pro = &CustomFileProcessor{InputChannel: ch, OutputChannel: out}

	var w CustomWriter

	w = &CustomFileWriter{WriteChannel: out, Path: "result.txt"}

	processor := &LogProcessor{Reader: reader, Writer: w, Processor: pro}

	for j := 0; j < 10; j++ {
		go processor.Reader.Read()
	}

	for i := 0; i < 10; i++ {
		go processor.Processor.Process()
	}

	go processor.Writer.Write()

	select {}

}
