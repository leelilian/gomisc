package worker

import (
	"log"

	"crawler/framework"
	"crawler/utils"
)

func Handle(r framework.Request) (*framework.ParseResult, error) {
	log.Printf("fetching url: %s", r.Url)
	resp, err := utils.Fetch(r.Url)
	if err != nil {
		log.Printf("error fetch url: %s, %v", r.Url, err)
		return nil, err

	}
	return r.Parser(resp), nil
}

func HandleAsync(in chan framework.Request, out chan *framework.ParseResult) {

	go func() {
		for {

			request := <-in
			result, err := Handle(request)
			if err != nil {
				continue
			}
			out <- result

		}
	}()

}
