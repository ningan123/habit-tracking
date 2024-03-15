package reading

import "time"

type MonthReading struct {
	MonthNum time.Month
	MonthReadingTime string 
	MothRawInfo  map[string]*DayReading  // string表示几号
}