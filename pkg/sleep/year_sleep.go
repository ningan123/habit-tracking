package sleep


type YearSleep struct {
	YearNum string
	RawInfo  map[string]*DaySleep  // string表示周几
	IsFinish bool
	TargetFinishDays int
	ActualFinishDays int
}

func NewYearSleep(yearNum string, rawInfo map[string]*DaySleep ) (*YearSleep, error) {
  return &YearSleep{
    YearNum: yearNum,
		RawInfo: rawInfo,
		TargetFinishDays: TargetYearFinishDays,
  }, nil
}


// 只要早起天数大于TargetFinishDays
func (y *YearSleep) CheckFinish() error {
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