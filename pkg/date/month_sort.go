package date

import (
	"fmt"
	"strconv"
	"strings"
)

// parseYearMonth 解析"年份-月份"格式的字符串，返回年份和月份
func parseYearMonth(s string) (int, int, error) {  
	parts := strings.Split(s, "-")  
	if len(parts) != 2 {  
		return 0, 0, fmt.Errorf("invalid year-month format: %s", s)  
	}  
  
	year, err := strconv.Atoi(parts[0])  
	if err != nil {  
		return 0, 0, err  
	}  
  
	month, err := strconv.Atoi(parts[1])  
	if err != nil {  
		return 0, 0, err  
	}  
  
	if month < 1 || month > 12 {  
		return 0, 0, fmt.Errorf("invalid month: %d", month)  
	}  
  
	return year, month, nil  
}  

// byYearMonth 实现了sort.Interface用于对字符串键进行排序  
type byYearMonth []string  
  
func (a byYearMonth) Len() int           { return len(a) }  
func (a byYearMonth) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }  
func (a byYearMonth) Less(i, j int) bool { return compareYearWeek(a[i], a[j]) < 0 }  
  
  
// compareYearMonth 比较两个"年份-月份"格式的字符串  
func compareYearMonth(a, b string) int {  
	yearA, monthA, _ := parseYearMonth(a)  
	yearB, monthB, _ := parseYearMonth(b)  
  
	if yearA != yearB {  
		return yearA - yearB  
	}  
	return monthA - monthB  
}  