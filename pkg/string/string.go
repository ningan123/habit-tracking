package string

import (
	"sort"
	"strings"
)

func StrTrimPrefix(str string, prefix string) string {
	var res string

	if strings.HasPrefix(str, prefix) {
		res = strings.TrimPrefix(str, prefix)
	} else {
		res = str
	}
	return res
}

// 输入："abc<br>bcd<br>"
func SortString(str string) string {
	var res string
	strList := strings.Split(strings.TrimSuffix(str, "<br>"), "<br>")

	// 使用 sort.Strings 对字符串数组进行排序
	sort.Strings(strList)

	for _, item := range strList {
		res += (item + "<br>")
	}
	return res
}
