package date

type Year struct {
  YearNum string // 年份
	DaysInYear int // 一年多少天
}

func NewYear(date string) (*Year, error) {
	_, yearNum,_, _, _,  _, _, _, _, _, _, daysInYear, err := GetDateDetails(date)
	if err != nil {
		return nil, err
	}	

  return &Year{
		YearNum: yearNum,
		DaysInYear: daysInYear,
  }, nil
}
