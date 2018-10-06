package parser

import (
	"regexp"
	"strconv"

	"crawler/framework"
	"crawler/model"
)

var ageRegx = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var incomeRegx = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var heightRegx = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)

func ParseProfile(contents []byte, name string) *framework.ParseResult {

	user := model.UserProfile{
		Name:   name,
		Age:    findInt(contents, ageRegx),
		Height: findInt(contents, heightRegx),
		Income: findString(contents, incomeRegx),
	}

	result := &framework.ParseResult{}
	result.Items = append(result.Items, user)

	return result
}

func findString(contents []byte, reg *regexp.Regexp) string {
	result := reg.FindSubmatch(contents)
	if len(result) >= 2 {
		return string(result[1])
	} else {
		return ""
	}

}

func findInt(contents []byte, reg *regexp.Regexp) int {
	result := reg.FindSubmatch(contents)
	var val = 0
	if len(result) >= 2 {
		val, _ = strconv.Atoi(string(result[1]))
	}
	return val

}
