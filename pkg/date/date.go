package date

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// dayOfYear 计算给定时间是一年中的第几天
func dayOfYear(t time.Time) int {  
	jan1 := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())  
	return int(t.Sub(jan1).Hours() / 24) +1 
} 

// dayOfYearFromString 计算给定日期字符串是一年中的第几天  
func dayOfYearFromString(dateStr string, layout string) (int, error) {  
	// 解析日期字符串为time.Time对象  
	t, err := time.Parse(layout, dateStr)  
	if err != nil {  
		return 0, err  
	}  
	// 使用之前定义的dayOfYear函数计算是一年中的第几天  
	return dayOfYear(t), nil  
}  


// TODO：2024-12-30的输出结果不太对
func GetDateDetails(inputDate string) (int, int, time.Month, int, int, time.Weekday, error) {
	// 将输入的日期解析为time.Time类型
	date, err := time.Parse("2006-01-02", inputDate)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, fmt.Errorf("无效的日期格式: %s", err)
	}

	// 获取年份、月份、周数和星期几
	year := date.Year()
	dayOfYear := dayOfYear(date)
	month := date.Month()
	dayOfMonth := date.Day()
	_, weekNumber := date.ISOWeek()
	weekday := date.Weekday()
	

	// 如果日期在第一周，但是年份不同，需要修正周数
	if weekNumber == 1 && date.Month() == 12 {
		year++
	}

	return year, dayOfYear, month, dayOfMonth, weekNumber, weekday, nil
}


// parseDuration 解析形如 "XhYmin" 或 "Ymin" 的字符串，并返回对应的 time.Duration  
func parseDuration(s string) (time.Duration, error) {  
	// 尝试先按照 "XhYmin" 格式解析  
	parts := strings.SplitN(s, "h", 2)  
	if len(parts) == 2 {  
		hours, err := strconv.Atoi(parts[0])  
		if err != nil {  
			return 0, fmt.Errorf("invalid hours in duration: %s", s)  
		} 
		if parts[1] == "" {
			return time.Duration(hours)*time.Hour, nil
		}
		minutes, err := strconv.Atoi(strings.TrimSuffix(parts[1], "min"))  
		if err != nil {  
			return 0, fmt.Errorf("invalid minutes in duration: %s", s)  
		}  
		return time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute, nil  
	}  
  
	// 如果不是 "XhYmin" 格式，则尝试按照 "Ymin" 格式解析  
	minutes, err := strconv.Atoi(strings.TrimSuffix(s, "min"))  
	if err != nil {  
		return 0, fmt.Errorf("invalid duration: %s", s)  
	}  
	return time.Duration(minutes) * time.Minute, nil  
}  


// formatDurationSum 接受两个表示时间的字符串，返回格式化后的字符串，格式为XhYmin  
func FormatDurationSum(durationStr1, durationStr2 string) (string, error) {  
	// 解析两个时长字符串  
	duration1, err := parseDuration(durationStr1)  
	if err != nil {  
		return "", err  
	}  
	duration2, err := parseDuration(durationStr2)  
	if err != nil {  
		return "", err  
	}  
  
	// 将两个时长相加  
	totalDuration := duration1 + duration2  
  
	// 提取小时和剩余的分钟数  
	hours := int(totalDuration.Hours())  
	remainingMinutes := int(totalDuration.Minutes()) % 60  
  
	// 返回格式化后的字符串  
	if hours > 0 {
		if remainingMinutes == 0 {
		  return fmt.Sprintf("%dh", hours), nil
		} else {
			return fmt.Sprintf("%dh%dmin", hours, remainingMinutes), nil
		}
	} else {		
		return fmt.Sprintf("%dmin", remainingMinutes), nil
	} 
}  


func FormatDurationMultiply(durationStr string, multiplier int) (string, error) {
	duration, err := parseDuration(durationStr)
	if err != nil {
		return "", err
	}

	resDuration := duration * time.Duration(multiplier)

	// 提取小时和剩余的分钟数  
	hours := int(resDuration.Hours())  
	remainingMinutes := int(resDuration.Minutes()) % 60  

	// 返回格式化后的字符串  
	if hours > 0 {
		if remainingMinutes == 0 {
		  return fmt.Sprintf("%dh", hours), nil
		} else {
			return fmt.Sprintf("%dh%dmin", hours, remainingMinutes), nil
		}
	} else {		
		return fmt.Sprintf("%dmin", remainingMinutes), nil
	} 

}


func IsActualDurationLongerOrEqualToTargetDuration(actualDurationStr, targetDurationStr string) (bool, error) {  
	// 解析两个时长字符串  
	actualDuration, err := parseDuration(actualDurationStr)  
	if err != nil {  
		return false, err  
	}  
	targetDuration, err := parseDuration(targetDurationStr)  
	if err != nil {  
		return false, err  
	}  	
	return actualDuration >= targetDuration, nil
}  