package date

type Week struct {
	WeekNum string
}

func NewWeek(weekNum string,) (*Week, error) {
  return &Week{
    WeekNum: weekNum,
  }, nil
}
