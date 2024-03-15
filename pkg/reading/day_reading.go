package reading

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type DayReading struct {
  DayRawInfo  string
	DayDate string 
	DayWeek time.Weekday
	DayMonth time.Month
	DayYear int
	DayReadingTime string
	DayReadingContent string 
}




func NewDayReading(date string, year int, month time.Month, weekNum int, weekday time.Weekday, dayRawInfo string) (*DayReading, error) {
	strList := strings.Split(dayRawInfo, ",")
	if len(strList) != 2 {
		errMsg := fmt.Sprintf("error split raw info: %s", dayRawInfo)
		return nil, errors.New(errMsg)
	}

	return &DayReading{
		DayRawInfo: dayRawInfo,
		DayDate: date,
		DayWeek: weekday,
		DayMonth: month,
		DayYear: year,
		DayReadingContent: strList[0],
		DayReadingTime: strList[1],
	}, nil

}
