package htmlparts

import (
	"os"
	"io/ioutil"
	"regexp"
	"strings"
 
)

const (
	regex_Loadtemplate = `<!--{{LoadTemplate(.{0,20})}}-->`
	regex_V            = `<!--{{=.{0,20}}}-->`
)

func Render(mainpage string, htmlparts map[string]string, contents map[string]string) string {
	return replaceAll(mainpage, htmlparts, contents)
}

func LoadPartFile(Viewpath string) (htmlparts map[string]string) {
	htmlparts = map[string]string{}
	f,err:=os.Open(Viewpath) 
	if(err!=nil){
		return 
	}
	fs,err:=f.Readdir(0)
 
	for i, part := range fs {
		if(part.IsDir()){
			continue
		}
		if m,err:=regexp.Match( ".+.part$",[]byte(part.Name()));m{
			i,er:=ioutil.ReadFile(part)
			if(er==nil){
				htmlparts[partnames[i]]  =  string(i)
			}
		}
		
		
		
	}
	return htmlparts
}
func replaceAll(input string, htmlparts map[string]string, contents map[string]string) string {
	var result string = input

	for key, value := range htmlparts {
		result = replace_LoadtemplatePart(result, key, value)
	}

	for key, value := range contents {
		result = replace_keyvalue(result, key, value)
	}

	return result
}
func replace_LoadtemplatePart(input string, key string, value string) string {

	return replace_regex(regex_Loadtemplate, 19, 6, input, key, value)
}

func replace_keyvalue(input string, key string, value string) string {

	return replace_regex(regex_V, 7, 5, input, key, value)
}

func replace_regex(regexstr string, namestart int, endlength int, input string, key string, value string) string {
	reg := regexp.MustCompile(regexstr)
	var mc []string
	mc = reg.FindAllString(input, -1)
	var result string = input
	for i := 0; i < len(mc); i++ {
		var tag string = mc[i]
		var str = tag[namestart : len(tag)-endlength]
		var pagename = strings.TrimSpace(str)
		if key == pagename {
			result = strings.Replace(result, tag, value, -1)
		}
	}
	return result
}
