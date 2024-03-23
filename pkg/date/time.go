package date

import (
	"fmt"
	"strings"
	"time"

	"k8s.io/klog/v2"
)

// 比较两个形如"6:30"的时间字符串
// 返回bool值，如果str1早于或等于str2则返回true，否则返回false
func IsStr1BeforeOrEqualStr2(str1, str2 string) (bool, error) {  
	// // 获取当前日期和时间  
	// now := time.Now()  
	// dateStr := now.Format("2006-01-02") // 获取当前日期  
  
	// // 构建完整的RFC3339格式的时间字符串  
	// fullStr1 := dateStr + "T" + str1 + ":00Z" // 假设秒和纳秒为0，并添加时区信息Z  
	// fullStr2 := dateStr + "T" + str2 + ":00Z"  
  
	// // 解析时间  
	// time1, err := time.Parse(time.RFC3339, fullStr1)  
	// if err != nil {  
	// 	return false, err  
	// }  
	// time2, err := time.Parse(time.RFC3339, fullStr2)  
	// if err != nil {  
	// 	return false, err  
	// }  

	time1, err := ParseStr(str1)
	if err != nil {  
		return false, err  
	}  
	time2, err := ParseStr(str2)
	if err != nil {  
		return false, err  
	}    
	// 比较时间  
	return time1.Before(time2) || time1.Equal(time2), nil   
}  


func ParseStr(str string)(time.Time, error)  {
	now := time.Now()  
	dateStr := now.Format("2006-01-02") // 获取当前日期

	strlist := strings.Split(str, "+")
	if len(strlist) == 2 && strlist[1] == "1" {	
		// 构建完整的RFC3339格式的时间字符串  
		fullStr := dateStr + "T" + strlist[0] + ":00Z" // 假设秒和纳秒为0，并添加时区信息Z
		parsedTime, err := time.Parse(time.RFC3339, fullStr)
		parsedTime = parsedTime.AddDate(0, 0, 1) // 增加一天 
		if err != nil {
			return parsedTime, err
		}
		fmt.Println(parsedTime)
		klog.V(2).InfoS("ParseStr +1", "parsedTime", parsedTime)
		return parsedTime, nil
	} else {
		fullStr := dateStr + "T" + str + ":00Z" // 假设秒和纳秒为0，并添加时区信息Z
		parsedTime, err := time.Parse(time.RFC3339, fullStr)
		if err != nil {
			return parsedTime, err
		}
		klog.V(2).InfoS("ParseStr", "parsedTime", parsedTime)
		return parsedTime, nil
	}
}