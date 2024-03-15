package date

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// TODO：2024-12-30的输出结果不太对
func GetDateDetails(inputDate string) (int, time.Month, int, time.Weekday, error) {
	// 将输入的日期解析为time.Time类型
	date, err := time.Parse("2006-01-02", inputDate)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("无效的日期格式: %s", err)
	}

	// 获取年份、月份、周数和星期几
	year := date.Year()
	month := date.Month()
	_, weekNumber := date.ISOWeek()
	weekday := date.Weekday()
	

	// 如果日期在第一周，但是年份不同，需要修正周数
	if weekNumber == 1 && date.Month() == 12 {
		year++
	}

	return year, month, weekNumber, weekday, nil
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
	return fmt.Sprintf("%dh%dmin", hours, remainingMinutes), nil  
}  