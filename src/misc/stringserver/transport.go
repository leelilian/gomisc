package stringserver

import (
	"context"
	"encoding/json"
	"net/http"
)

func HttpDecodeUppercaseRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func HttpEncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func GrpcDecodeUppercaseRequest(ctx context.Context, req interface{}) (request interface{}, err error) {
	return req, nil
}

func GrpcEncodeResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}
