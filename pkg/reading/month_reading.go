package reading

import (
	"time"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type MonthReading struct {
	MonthNum time.Month
	MonthReadingTime string 
	MonthRawInfo  map[int]*DayReading  // int表示几号
}


func NewMonthReading(month time.Month, monthRawInfo map[int]*DayReading) (*MonthReading, error) {
	return &MonthReading{
		MonthNum: month, 
		MonthReadingTime: "0min",
		MonthRawInfo: monthRawInfo,
	},nil
}

func (m *MonthReading) ComputeReadingTime() error {
	for _, dayReading := range m.MonthRawInfo {
	  sum, err := hDate.FormatDurationSum(m.MonthReadingTime, dayReading.DayReadingTime)
		if err != nil {
			return err 
		}
		m.MonthReadingTime = sum
	}
	 return nil
}	


func (m *MonthReading) Print() {
	klog.InfoS("month reading info", "monthNum", m.MonthNum, "readingTime", m.MonthReadingTime)
}
