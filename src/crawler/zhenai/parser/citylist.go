package parser

import (
	"log"
	"regexp"

	"crawler/framework"
)

var cityListReg = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(contents []byte) *framework.ParseResult {

	// compile := regexp.MustCompile(cityListReg)
	result := &framework.ParseResult{}
	all := cityListReg.FindAllSubmatch(contents, -1)

	for _, m := range all {
		log.Printf("city: %s, url: %s", string(m[2]), string(m[1]))
		result.Items = append(result.Items, string(m[2]))
		result.RequestList = append(result.RequestList, framework.Request{
			Url:    string(m[1]),
			Parser: ParseCity,
		})

	}
	return result
}
