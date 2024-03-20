package getup

import "time"


type DayGetup struct {
  RawInfo  string

	Date string // 具体日期
	Weekday time.Weekday // 星期几
	Month time.Month // 几月
	Year int // 哪一年
	DayOfYear int // 一年中的第几天
	DayOfMonth int // 一个月中的第几天

	IsFinish bool
	TargetTime string
}

func NewDayGetup(date string, year int, dayOfYear int, month time.Month, dayOfMonth int, weekNum string, weekday time.Weekday, rawInfo string) (*DayGetup, error) {
	return &DayGetup{
		RawInfo: rawInfo,
		Date: date,
		Weekday: weekday,
		DayOfYear: dayOfYear,
		DayOfMonth: dayOfMonth,
		Month: month,
		Year: year,

	}, nil	
}


func (d *DayGetup) CheckFinish() error {
	return nil
}