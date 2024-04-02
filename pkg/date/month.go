package date


type Month struct {
	MonthNum string
	DaysInMonth int
}


func NewMonth(monthNum string, daysInMonth int) (*Month, error) {
	return &Month{
		MonthNum: monthNum, 
		DaysInMonth: daysInMonth,
	}, nil
}