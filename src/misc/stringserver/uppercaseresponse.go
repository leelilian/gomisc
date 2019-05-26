package stringserver

type uppercaseResponse struct {
	Response string `json:"response"`
	ErrMsg   string `json:"errMsg,omitempty"` // errors don't JSON-marshal, so we use a string
}
