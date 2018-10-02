package main


import (
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	/*
		mx := http.ServeMux{}
		http.HandleFunc("/", fuck)
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println(err)
		}*/
	mx := http.NewServeMux()
	mx.Handle("/", &myhandler{})
	
	log.Println("server starting....")
	err := http.ListenAndServe(":8080", mx)
	if err != nil {
		log.Fatal(err)
	}
	
	strings.Split()
	
}

type myhandler struct {
}

func (this *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "fuck and yourself")
}
