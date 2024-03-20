package date

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

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

func parseDurationToString(duration time.Duration) string {
	// 提取小时和剩余的分钟数  
	hours := int(duration.Hours())  
	remainingMinutes := int(duration.Minutes()) % 60 
	// klog.InfoS("parseDurationToString", "hours", hours, "remainingMinutes", remainingMinutes) 
	
	// 返回格式化后的字符串  
	if hours > 0 {
		if remainingMinutes == 0 {
			return fmt.Sprintf("%dh", hours)
		} else {
			return fmt.Sprintf("%dh%dmin", hours, remainingMinutes)
		}
	} else if hours == 0 {		
		if remainingMinutes == 0 {
			return "0"
		} else {
			return fmt.Sprintf("%dmin", remainingMinutes)
		}		
	} else {
		if remainingMinutes == 0 {
			return fmt.Sprintf("%dh", hours) 
		} else {
			return fmt.Sprintf("%dh%dmin", hours, -remainingMinutes)
		}
	}
}


func FormatDurationSub(durationStr1, durationStr2 string) (string, error) {
	// 解析两个时长字符串  
	duration1, err := parseDuration(durationStr1)  
	if err != nil {  
		return "", err  
	}  
	duration2, err := parseDuration(durationStr2)  
	if err != nil {  
		return "", err  
	}  

	// 将两个时长相减
	subDuration := duration1 - duration2  
	// klog.InfoS("FormatDurationSub", "subDuration", subDuration)

	return parseDurationToString(subDuration), nil
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
  
	return parseDurationToString(totalDuration),nil
}  


func FormatDurationMultiply(durationStr string, multiplier int) (string, error) {
	duration, err := parseDuration(durationStr)
	if err != nil {
		return "", err
	}

	resDuration := duration * time.Duration(multiplier)

	return parseDurationToString(resDuration),nil
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