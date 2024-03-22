package getup


type MonthGetup struct {
	MonthNum string
	RawInfo  map[int]*DayGetup   // int表示几号
	IsFinish bool
	TargetFinishDays int
	ActualFinishDays int
}

func NewMonthGetup(weekNum string, rawInfo map[int]*DayGetup ) (*MonthGetup, error) {
  return &MonthGetup{
    MonthNum: weekNum,
		RawInfo: rawInfo,
		TargetFinishDays: TargetMonthFinishDays,
  }, nil
}


// 只要早起天数大于TargetFinishDays
func (m *MonthGetup) CheckFinish() error {
	
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