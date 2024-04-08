package sleep

import (
	hDate "ningan.com/habit-tracking/pkg/date"
)

type MonthSleep struct {
	Month *hDate.Month
	RawInfo  map[int]*DaySleep   // int表示几号
	IsFinish bool
	TargetFinishDays int
	ActualFinishDays int
}

func NewMonthSleep(monthNum string, daysInMonth int,  rawInfo map[int]*DaySleep ) (*MonthSleep, error) {
  return &MonthSleep{
    Month: &hDate.Month{
			MonthNum: monthNum,
			DaysInMonth: daysInMonth, 
		},
		RawInfo: rawInfo,
		TargetFinishDays: TargetMonthFinishDays,
  }, nil
}


// 只要早起天数大于TargetFinishDays
func (m *MonthSleep) CheckFinish() error {
	
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