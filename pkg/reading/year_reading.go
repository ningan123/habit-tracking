package reading

type YearReading struct {
	YearRawInfo []string 
  YearReadingTime string // 年总阅读时长
	
}

func NewYearReading(yearRawInfo []string) *YearReading {
  return &YearReading{
    YearRawInfo: yearRawInfo,
		
  }
}



