package date


type Month struct {
	MonthNum string
	DaysInMonth int
}

func NewMonth(date string) (*Month, error) {
	_, _, _,monthNum, _,  _, _, _, _, _, daysInMonth, _, err := GetDateDetails(date)
	if err != nil {
		return nil, err
	}	

	return &Month{
		MonthNum: monthNum, 
		DaysInMonth: daysInMonth,
	}, nil
}