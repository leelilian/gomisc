package stringserver

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeUppercaseEndpoint(svr StringService) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(uppercaseRequest)
		rsp, err := svr.ToUpper(req.Request)
		if err != nil {
			return uppercaseResponse{"", err.Error()}, err
		}
		return uppercaseResponse{rsp, ""}, nil
	}
}
