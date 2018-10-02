package misc

import (
	"fmt"
	"net/http"
)

func main() {
	
	rsp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	
	defer rsp.Body.Close()
	
	header := rsp.Header
	for k, v := range header {
		fmt.Printf("key=%s, value=%s\n", k, v)
	}
	
	fmt.Printf("resp status %s,statusCode %d\n", rsp.Status, rsp.StatusCode)
	
	fmt.Printf("resp Proto %s\n", rsp.Proto)
	
	fmt.Printf("resp content length %d\n", rsp.ContentLength)
	
}
