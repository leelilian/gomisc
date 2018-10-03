package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	if len(url) == 0 {
		return nil, fmt.Errorf("empty url")
	}
	response, err := http.Get(url)
	if err != nil {
		log.Printf("error %v", err)
		return nil, err
	}

	defer response.Body.Close()

	reader := bufio.NewReader(response.Body)

	e := determineEncoding(reader)

	utf8reader := transform.NewReader(reader, e.NewDecoder())

	if response.StatusCode != http.StatusOK {
		log.Printf("url: %s, error http code: %d", url, response.StatusCode)
		return nil, fmt.Errorf("wrong http code: %d", response.StatusCode)
	}

	return ioutil.ReadAll(utf8reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {

	buffer, err := r.Peek(1024)
	if err != nil {
		log.Printf("error ocurred when determining encoding, utf8 will set as default. %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(buffer, "")

	return e
}
