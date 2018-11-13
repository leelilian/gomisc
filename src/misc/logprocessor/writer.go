package main

import (
	"bufio"
	"fmt"
	"os"
)

type CustomWriter interface {
	Write()
}

type CustomFileWriter struct {
	WriteChannel chan []byte
	Path         string
}

func (cw *CustomFileWriter) CreateFile() error {
	file, err := os.Create(cw.Path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func (cw *CustomFileWriter) Write() {

	file, err := os.OpenFile(cw.Path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	newWriter := bufio.NewWriter(file)

	for {
		content := <-cw.WriteChannel
		// 使用NewWriter方法返回的io.Writer缓冲默认大小为4096，也可以使用NewWriterSize方法设置缓存的大小
		fmt.Printf("write:%s\n", string(content))
		// 将文件写入缓存
		if _, err = newWriter.Write(content); err != nil {
			fmt.Println(err)
		}
		// 从缓存写入到文件中
		if err = newWriter.Flush(); err != nil {
			fmt.Println(err)
		}

	}

}
