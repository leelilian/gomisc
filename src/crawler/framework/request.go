package framework

type Request struct {
	Url    string
	Parser func(contents []byte) *ParseResult
}
