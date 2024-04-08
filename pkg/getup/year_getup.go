package getup

import (
	hDate "ningan.com/habit-tracking/pkg/date"
)

type YearGetup struct {
	Year *hDate.Year
	YearNum string
	RawInfo  map[string]*DayGetup  // string表示周几
	IsFinish bool
	TargetFinishDays int
	ActualFinishDays int
}

func NewYearGetup(yearNum string, daysInYear int, rawInfo map[string]*DayGetup ) (*YearGetup, error) {
  return &YearGetup{
		Year: &hDate.Year{
			YearNum: yearNum,
			DaysInYear: daysInYear,
		},
    YearNum: yearNum,
		RawInfo: rawInfo,
		TargetFinishDays: TargetYearFinishDays,
  }, nil
}


// 只要早起天数大于TargetFinishDays
func (y *YearGetup) CheckFinish() error {
	for _, v := range y.RawInfo {
		err := v.CheckFinish()
		if err != nil {
			return err
		}
	  if v.IsFinish {
			y.ActualFinishDays++
		}
	}
	
	if y.ActualFinishDays >= y.TargetFinishDays {
		y.IsFinish = true
	}

	return nil
}