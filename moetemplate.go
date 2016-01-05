package moetemplate

import (
	"log"
	"regexp"
	"strconv"
	"strings"
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
	return replace_regex(regexstr, 15, len(regexstr)-4-14, input, key, value)
}

func Replace_keyvalue(input string, key string, value string) string {
	var regexstr string = `{{=.{0,20}}}`
	return replace_regex(regexstr, 3, len(regexstr)-4-1, input, key, value)
}
func replace_regex(regexstr string, namestart int, namelength int, input string, key string, value string) string {
	reg := regexp.MustCompile(regexstr)
	var mc []string
	mc = reg.FindAllString(input, -1)
	var result string = input
	for i := 0; i < len(mc); i++ {
		var tag string = mc[i]
		rs := []rune(tag)
		log.Println(tag + "\t" + strconv.Itoa(namestart) + "\t" + strconv.Itoa(namelength))
		var str = (rs[namestart:namelength])
		log.Println(string(str))
		var pagename = strings.TrimSpace(string(str))
		if key == pagename {
			result = strings.Replace(result, key, value, -1)
		}
	}
	return result
}
