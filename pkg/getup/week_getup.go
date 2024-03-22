package getup


type WeekGetup struct {
	WeekNum string
	RawInfo  map[string]*DayGetup  // string表示周几
	IsFinish bool
	TargetFinishDays int
	ActualFinishDays int
}

func NewWeekGetup(weekNum string, rawInfo map[string]*DayGetup ) (*WeekGetup, error) {
  return &WeekGetup{
    WeekNum: weekNum,
		RawInfo: rawInfo,
		TargetFinishDays: TargetWeekFinishDays,
  }, nil
}


// 只要早起天数大于TargetFinishDays
func (w *WeekGetup) CheckFinish() error {
	
	for _, v := range w.RawInfo {
		err := v.CheckFinish()
		if err != nil {
			return err
		}
	  if v.IsFinish {
			w.ActualFinishDays++
		}
	}
	// fmt.Println("w.ActualFinishDays", w.ActualFinishDays)

	if w.ActualFinishDays >= w.TargetFinishDays {
		w.IsFinish = true
	}

	return nil
}