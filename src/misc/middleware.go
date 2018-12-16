package main

import (
	"log"
	"net/http"
)

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("OK"))
}

func main() {

	handler := http.HandlerFunc(final)
	http.Handler
}
