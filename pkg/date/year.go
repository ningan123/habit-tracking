package date

type Year struct {
  YearNum string // 年份
	DaysInYear int // 一年多少天
}

func NewYear(yearNum string, daysInYear int) (*Year, error) {
  return &Year{
		YearNum: yearNum,
		DaysInYear: daysInYear,
  }, nil
}
