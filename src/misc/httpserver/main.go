package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.Handle("/", &sampleserver{})

	http.ListenAndServe(":3000", nil)

}

type sampleserver struct {
}

func (s *sampleserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	go func() {
		time.Sleep(4 * time.Second)
		fmt.Println("done")
	}()

	w.Write([]byte("Hello"))

}

type wr struct {
}

func (w *wr) Header() http.Header {

	head := http.Header{}
	head.Add("content-type", "json")
	return head
}

func (w *wr) Write(content []byte) (int, error) {
	return 0, nil
}

func (w *wr) WriteHeader(statusCode int) {

}
