package getup

import (
	"time"

	hDate "ningan.com/habit-tracking/pkg/date"
)


type DayGetup struct {
  RawInfo  string

	Date string // 具体日期
	Weekday string // 星期几
	Month time.Month // 几月
	WeekNum string // 几周
	Year int // 哪一年
	DayOfYear int // 一年中的第几天
	DayOfMonth int // 一个月中的第几天

	IsFinish bool
	TargetTime string
}

func NewDayGetup(date string, year int, dayOfYear int, month time.Month, dayOfMonth int, weekNum string, weekday string, rawInfo string) (*DayGetup, error) {
	return &DayGetup{
		RawInfo: rawInfo,
		Date: date,
		Weekday: weekday,
		DayOfYear: dayOfYear,
		DayOfMonth: dayOfMonth,
		WeekNum: weekNum,
		Month: month,
		Year: year,
	
		TargetTime: TargetDayGetupTime,
	}, nil	
}


func (d *DayGetup) CheckFinish() error {
	res, err := hDate.IsStr1BeforeOrEqualStr2(d.RawInfo, d.TargetTime)
	if err != nil {
		return err
	}
	d.IsFinish = res
	return nil
}