package date

import (
	"fmt"
	"strconv"
	"strings"
)

// parseYearWeek 解析"年份-周数"格式的字符串，返回年份和周数  
func parseYearWeek(s string) (int, int, error) {  
	parts := strings.Split(s, "-")  
	if len(parts) != 2 {  
		return 0, 0, fmt.Errorf("invalid year-week format: %s", s)  
	}  
  
	year, err := strconv.Atoi(parts[0])  
	if err != nil {  
		return 0, 0, err  
	}  
  
	week, err := strconv.Atoi(parts[1])  
	if err != nil {  
		return 0, 0, err  
	}  
  
	return year, week, nil  
}  
  
// byYearWeek 实现了sort.Interface用于对字符串键进行排序  
type byYearWeek []string  
  
func (a byYearWeek) Len() int           { return len(a) }  
func (a byYearWeek) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }  
func (a byYearWeek) Less(i, j int) bool { return compareYearWeek(a[i], a[j]) < 0 }  
  
// compareYearWeek 比较两个"年份-周数"格式的字符串  
func compareYearWeek(a, b string) int {  
	yearA, weekA, _ := parseYearWeek(a)  
	yearB, weekB, _ := parseYearWeek(b)  
  
	if yearA != yearB {  
		return yearA - yearB  
	}  
	return weekA - weekB  
}  