package tick_type

import (
	hDate "ningan.com/habit-tracking/pkg/date"
)

type YearTickType struct {
	Year *hDate.Year
	RawInfo  map[string]*DayTickType  // string表示周几
	IsFinish bool
	TargetFinishDays int
	ActualFinishDays int
}

func NewYearTickType(yearNum string, daysInYear int, rawInfo map[string]*DayTickType, targetYearFinishDays int ) (*YearTickType, error) {
  return &YearTickType{
    Year: &hDate.Year{
			YearNum: yearNum,
			DaysInYear: daysInYear,
		},
		RawInfo: rawInfo,
		TargetFinishDays: targetYearFinishDays,
  }, nil
}


// 只要完成天数大于TargetFinishDays
func (y *YearTickType) CheckFinish() error {
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