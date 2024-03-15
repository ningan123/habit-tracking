package date

import (
	"fmt"
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