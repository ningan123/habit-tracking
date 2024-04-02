package date

type Day struct {
	Date string // 具体日期
	Weekday string // 星期几

	WeekNum string // 哪周
	MonthNum string // 哪月
	YearNum string // 哪年

	DayOfMonth int // 一个月中的第几天	
	DayOfYear int // 一年中的第几天
}


func NewDay(date string) (*Day, error) {
	_, yearNum, _, monthNum, _, _, weekday, weekNum, dayOfMonth, dayOfYear, _, _, err := GetDateDetails(date)
	if err != nil {
		return nil, err
	}

	return &Day{
		Date: date,
		Weekday: weekday,

		WeekNum: weekNum,
		MonthNum: monthNum,
		YearNum: yearNum,

		DayOfYear: dayOfYear,
		DayOfMonth: dayOfMonth,
	}, nil	
}

