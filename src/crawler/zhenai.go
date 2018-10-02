package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func main() {
	
	response, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		log.Printf("error %s", err)
	}
	
	defer response.Body.Close()
	
	e := determineEncoding(response.Body)
	
	utf8reader := transform.NewReader(response.Body, e.NewDecoder())
	
	if response.StatusCode != http.StatusOK {
		log.Printf("error http code: %d", response.StatusCode)
		return
	}
	
	body, _ := ioutil.ReadAll(utf8reader)
	
	getAllCity(body)
	
}

func getAllCity(contents []byte) {
	
	compile := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	all := compile.FindAllSubmatch(contents, -1)
	for _, m := range all {
		fmt.Printf("city:%s, url:%s\n", m[2], m[1])
	}
	
}

func determineEncoding(r io.Reader) encoding.Encoding {
	
	buffer, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	
	e, _, _ := charset.DetermineEncoding(buffer, "")
	
	return e
}
