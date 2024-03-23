package string

import "strings"


func StrTrimPrefix(str string, prefix string) string {
	var res string
	// 判断字符串是否以"~"开头
	if strings.HasPrefix(str, prefix) {  
		res = strings.TrimPrefix(str, prefix)  
	} else {
		res = str
	}	
	return res
}