package date

import "time"

// 比较两个形如"6:30"的时间字符串
// 返回bool值，如果str1早于或等于str2则返回true，否则返回false
func IsStr1BeforeOrEqualStr2(str1, str2 string) (bool, error) {  
	// 获取当前日期和时间  
	now := time.Now()  
	dateStr := now.Format("2006-01-02") // 获取当前日期  
  
	// 构建完整的RFC3339格式的时间字符串  
	fullStr1 := dateStr + "T" + str1 + ":00Z" // 假设秒和纳秒为0，并添加时区信息Z  
	fullStr2 := dateStr + "T" + str2 + ":00Z"  
  
	// 解析时间  
	time1, err := time.Parse(time.RFC3339, fullStr1)  
	if err != nil {  
		return false, err  
	}  
	time2, err := time.Parse(time.RFC3339, fullStr2)  
	if err != nil {  
		return false, err  
	}  
  
	// 比较时间  
	return time1.Before(time2) || time1.Equal(time2), nil   
}  