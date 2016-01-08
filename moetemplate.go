package moetemplate

import (
	"regexp"

	"github.com/golangframework/xstring"
)

func Render(mainpage string, htmlparts map[string]string, contents map[string]string) string {
	return replaceAll(mainpage, htmlparts, contents)
}
func replaceAll(input string, htmlparts map[string]string, contents map[string]string) string {
	var result string = input

	for key, value := range htmlparts {
		result = replace_LoadtemplatePart(result, key, value)
	}
	for key, value := range contents {
		result = Replace_keyvalue(result, key, value)
	}

	return result
}
func replace_LoadtemplatePart(input string, key string, value string) string {
	var regexstr string = `{{LoadTemplate(.{0,20})}}`
	return replace_regex(regexstr, 15, 15+3, input, key, value)
}

func Replace_keyvalue(input string, key string, value string) string {
	var regexstr string = `{{=.{0,20}}}`
	return replace_regex(regexstr, 3, 3+2, input, key, value)
}

func Replace_Repeat(input string, key string, value string) string {
	var regexstr string = `{{Repeat(.{0,20})}}`
	return replace_regex(regexstr, 9, 9+3, input, key, value)
}

func replace_regex(regexstr string, namestart int, otherlength int, input string, key string, value string) string {
	reg := regexp.MustCompile(regexstr)
	var mc []string
	mc = reg.FindAllString(input, -1)
	var result string = input
	for i := 0; i < len(mc); i++ {
		var tag string = mc[i]
		var str = xstring.Substring(tag, namestart, len(tag)-otherlength)
		var pagename = xstring.Trim(str)
		if key == pagename {
			result = xstring.Replace(result, tag, value)
		}
	}
	return result
}
