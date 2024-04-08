package getup

import (
	hDate "ningan.com/habit-tracking/pkg/date"
	hString "ningan.com/habit-tracking/pkg/string"
)


type DayGetup struct {
	Day *hDate.Day
  RawInfo  string
	IsFinish bool
	TargetTime string
}

func NewDayGetup(date string, weekday string, weekNum string, monthNum string, yearNum string, dayOfMonth int, dayOfYear int, rawInfo string) (*DayGetup, error) {

	return &DayGetup{
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
		TargetTime: TargetDayGetupTime,
	}, nil	
}


func (d *DayGetup) CheckFinish() error {
	if d.RawInfo == "" || d.RawInfo == "×" {
	  return nil
	}
	
	res, err := hDate.IsStr1BeforeOrEqualStr2(hString.StrTrimPrefix(d.RawInfo, "~"), d.TargetTime)
	if err != nil {
		return err
	}
	d.IsFinish = res
	return nil
}