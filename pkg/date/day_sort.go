package date

import (
	"fmt"
	"time"
)

// parseDate 解析"年-月-日"格式的字符串，返回time.Time类型
func parseDate(s string) (time.Time, error) {  
	const layout = "2006-01-02" // Go的布局字符串，代表年-月-日  
	t, err := time.Parse(layout, s)  
	if err != nil {  
		return time.Time{}, err  
	}  
	return t, nil  
}

 
type ByDate []string  

func (a ByDate) Len() int           { return len(a) }  
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }  
func (a ByDate) Less(i, j int) bool { return compareDates(a[i], a[j]) < 0 }  
  
// compareDates 比较两个"年-月-日"格式的字符串  
func compareDates(a, b string) int {  
	dateA, errA := parseDate(a)  
	dateB, errB := parseDate(b)  
	if errA != nil || errB != nil {  
		// 如果解析出错，可以根据需要处理错误  
		panic(fmt.Sprintf("Invalid date format: %s or %s", a, b))  
	}  
	return dateA.Compare(dateB)  
}  
  