package tick_type

import (
	hDate "ningan.com/habit-tracking/pkg/date"
)


type DayTickType struct {
	Day *hDate.Day
  RawInfo  string
	IsFinish bool
}

func NewDayTickType(date string, weekday string, weekNum string, monthNum string, yearNum string, dayOfMonth int, dayOfYear int, rawInfo string) (*DayTickType, error) {
	return &DayTickType{
		Day: &hDate.Day{
			Date: date,
			Weekday: weekday,
			WeekNum: weekNum,
			MonthNum: monthNum,
			YearNum: yearNum,
			DayOfMonth: dayOfMonth,
			DayOfYear: dayOfYear,
		},
		RawInfo: rawInfo,
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

