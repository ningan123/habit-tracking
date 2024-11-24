package sleep

import (
	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
	hString "ningan.com/habit-tracking/pkg/string"
)


type DaySleep struct {
	Day *hDate.Day
  RawInfo  string
	IsFinish bool
	TargetTime string
}

func NewDaySleep(date string, weekday string, weekNum string, monthNum string, yearNum string, dayOfMonth int, dayOfYear int, rawInfo string) (*DaySleep, error) {
	return &DaySleep{
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
		TargetTime: TargetDaySleepTime,
	}, nil	
}


func (d *DaySleep) CheckFinish() error {
	if d.RawInfo == "" || d.RawInfo == "×" {
	  return nil
	}
	klog.V(2).InfoS("check day sleep", "date", d.Day.Date, "rawInfo", d.RawInfo)

	res, err := hDate.IsStr1BeforeOrEqualStr2(hString.StrTrimPrefix(d.RawInfo, "~"), d.TargetTime)
	if err != nil {
		return err
	}
	d.IsFinish = res
	return nil
}

