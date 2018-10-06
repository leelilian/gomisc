package parser

import (
	"log"
	"regexp"

	"crawler/framework"
)

var cityReg = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

func ParseCity(contents []byte) *framework.ParseResult {

	result := &framework.ParseResult{}
	all := cityReg.FindAllSubmatch(contents, -1)

	for _, m := range all {
		log.Printf("user: %s, url: %s", string(m[2]), string(m[1]))
		name := string(m[2])
		result.Items = append(result.Items, name)
		result.RequestList = append(result.RequestList, framework.Request{
			Url: string(m[1]),
			Parser: func(contents []byte) *framework.ParseResult {
				return ParseProfile(contents, name)
			},
		})
	}
	return result
}
