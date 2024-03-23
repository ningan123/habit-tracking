package sleep


type MonthSleep struct {
	MonthNum string
	RawInfo  map[int]*DaySleep   // int表示几号
	IsFinish bool
	TargetFinishDays int
	ActualFinishDays int
}

func NewMonthSleep(weekNum string, rawInfo map[int]*DaySleep ) (*MonthSleep, error) {
  return &MonthSleep{
    MonthNum: weekNum,
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