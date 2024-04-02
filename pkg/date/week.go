package date

type Week struct {
	WeekNum string
}

func NewWeek(date string,) (*Week, error) {
	_, _, _, _, _, _,_, weekNum, _, _, _, _, err := GetDateDetails(date)
	if err != nil {
		return nil, err
	}

  return &Week{
    WeekNum: weekNum,
  }, nil
}


