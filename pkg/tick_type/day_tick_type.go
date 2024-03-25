package tick_type

import (
	"time"
)


type DayTickType struct {
  RawInfo  string

	Date string // 具体日期
	Weekday string // 星期几
	Month time.Month // 几月
	WeekNum string // 几周
	Year int // 哪一年
	DayOfYear int // 一年中的第几天
	DayOfMonth int // 一个月中的第几天

	IsFinish bool
}

func NewDayTickType(date string, year int, dayOfYear int, month time.Month, dayOfMonth int, weekNum string, weekday string, rawInfo string) (*DayTickType, error) {

	return &DayTickType{
		RawInfo: rawInfo,
		Date: date,
		Weekday: weekday,
		DayOfYear: dayOfYear,
		DayOfMonth: dayOfMonth,
		WeekNum: weekNum,
		Month: month,
		Year: year,
	}, nil	
}


func (d *DayTickType) CheckFinish() error {
	if d.RawInfo == "" || d.RawInfo == "×" {
	  return nil
	}

	if d.RawInfo == "√" {
		d.IsFinish = true
		return nil
	}
	return nil
}

