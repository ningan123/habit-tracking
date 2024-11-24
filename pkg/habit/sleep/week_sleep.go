package sleep

import (
	hDate "ningan.com/habit-tracking/pkg/date"
)


type WeekSleep struct {
	Week *hDate.Week
	RawInfo  map[string]*DaySleep  // string表示周几
	IsFinish bool
	TargetFinishDays int
	ActualFinishDays int
}

func NewWeekSleep(weekNum string, rawInfo map[string]*DaySleep ) (*WeekSleep, error) {
  return &WeekSleep{
    Week: &hDate.Week{
			WeekNum: weekNum,
		},
		RawInfo: rawInfo,
		TargetFinishDays: TargetWeekFinishDays,
  }, nil
}


// 只要早起天数大于TargetFinishDays
func (w *WeekSleep) CheckFinish() error {
	
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