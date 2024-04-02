package date

import "time"

type Day struct {
	Date string // 具体日期
	Weekday string // 星期几

	WeekNum string // 哪周
	Month time.Month // 哪月
	Year int // 哪年

	DayOfYear int // 一年中的第几天
	DayOfMonth int // 一个月中的第几天	
}


func NewDay(date string, weekday string, weekNum string, month time.Month, year int, dayOfYear int,  dayOfMonth int, ) (*Day, error) {
	return &Day{
		Date: date,
		Weekday: weekday,

		WeekNum: weekNum,
		Month: month,
		Year: year,

		DayOfYear: dayOfYear,
		DayOfMonth: dayOfMonth,
	}, nil	
}