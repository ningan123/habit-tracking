package getup


type WeekGetup struct {
	WeekNum string
	WeekRawInfo  map[string]*DayGetup  // string表示周几
	IsFinish bool
	TargetFinishDays int
	ActualFinishDays int
}

func NewWeekGetup(weekNum string, weekRawInfo map[string]*DayGetup ) (*WeekGetup, error) {
  return &WeekGetup{
    WeekNum: weekNum,
		WeekRawInfo: weekRawInfo,
		TargetFinishDays: 5,
  }, nil
}


// 只要早起天数大于TargetFinishDays
func (w *WeekGetup) CheckFinish() error {
	
	for _, v := range w.WeekRawInfo {
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