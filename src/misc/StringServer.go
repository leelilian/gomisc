package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"misc/stringserver"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)
	var svc stringserver.StringService
	svc = stringserver.StringSvr{}
	svc = stringserver.LoggingMiddleware{logger, svc}
	handler := httptransport.NewServer(stringserver.MakeUppercaseEndpoint(svc),
		stringserver.HttpDecodeUppercaseRequest,
		stringserver.HttpEncodeResponse)

	http.Handle("/uppercase", handler)

	http.ListenAndServe(":9090", nil)
}
