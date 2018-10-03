package engine

import (
	"log"

	"crawler/utils"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, m := range seeds {
		requests = append(requests, m)
	}

	for len(requests) > 0 {

		request := requests[0]
		requests = requests[1:]
		log.Printf("fetching url: %s", request.Url)
		resp, err := utils.Fetch(request.Url)
		if err != nil {
			log.Printf("error fetch url: %s, %v", request.Url, err)
			continue

		}
		result := request.Parser(resp)
		requests = append(requests, result.RequestList...)

		for _, item := range result.Items {
			log.Printf("city: %v", item)
		}

	}
}
