package main

import (
	"log"
	"time"

	"crawler/engine"
	"crawler/framework"
	"crawler/zhenai/parser"
)

func main() {

	start := time.Now()
	engine.Run(framework.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: parser.ParseCityList,
	})

	end := time.Now()

	log.Printf("crawler ellapsed: %s", end.Sub(start))

}
