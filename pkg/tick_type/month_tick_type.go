package tick_type

import (
	hDate "ningan.com/habit-tracking/pkg/date"
)

type MonthTickType struct {
	Month *hDate.Month
	RawInfo  map[int]*DayTickType   // int表示几号
	IsFinish bool
	TargetFinishDays int
	ActualFinishDays int
}

func NewMonthTickType(monthNum string, daysInMonth int, rawInfo map[int]*DayTickType, targetMonthFinishDays int ) (*MonthTickType, error) {
  return &MonthTickType{
    Month: &hDate.Month{
			MonthNum: monthNum,
			DaysInMonth: daysInMonth, 
		},
		RawInfo: rawInfo,
		TargetFinishDays: targetMonthFinishDays,
  }, nil
}


// 只要完成天数大于TargetFinishDays
func (m *MonthTickType) CheckFinish() error {
	
	for _, v := range m.RawInfo {
		err := v.CheckFinish()
		if err != nil {
			return err
		}
	  if v.IsFinish {
			m.ActualFinishDays++
		}
	}
	
	if m.ActualFinishDays >= m.TargetFinishDays {
		m.IsFinish = true
	}

	return nil
}