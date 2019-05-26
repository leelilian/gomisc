package stringserver

import (
	"errors"
	"strings"
)

type StringService interface {
	ToUpper(s string) (string, error)
}

type StringSvr struct {
}

func (svr StringSvr) ToUpper(s string) (string, error) {
	if s == "" {
		return "", errors.New("input is empty")

	}
	return strings.ToUpper(s), nil
}
