package tick_type

import (
	hDate "ningan.com/habit-tracking/pkg/date"
)

type WeekTickType struct {
	Week *hDate.Week
	RawInfo  map[string]*DayTickType  // string表示周几
	IsFinish bool
	TargetFinishDays int
	ActualFinishDays int
}

func NewWeekTickType(weekNum string, rawInfo map[string]*DayTickType, targetWeekFinishDays int ) (*WeekTickType, error) {
  return &WeekTickType{
    Week: &hDate.Week{
			WeekNum: weekNum,
		},
		RawInfo: rawInfo,
		TargetFinishDays: targetWeekFinishDays,
  }, nil
}


// 只要完成天数大于TargetFinishDays
func (w *WeekTickType) CheckFinish() error {
	
	for _, v := range w.RawInfo {
		err := v.CheckFinish()
		if err != nil {
			return err
		}
	  if v.IsFinish {
			w.ActualFinishDays++
		}
	}

	if w.ActualFinishDays >= w.TargetFinishDays {
		w.IsFinish = true
	}

	return nil
}