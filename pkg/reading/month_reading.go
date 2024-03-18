package reading

import (
	"fmt"
	"time"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type MonthReading struct {
	MonthNum time.Month
	MonthReadingTime string 
	MonthReadingTimeOfDifferentContent map[string]string // 不同内容的阅读时间
	MonthReadingTimeOfDifferentContentStr string
	MonthRawInfo  map[int]*DayReading  // int表示几号
}


func NewMonthReading(month time.Month, monthRawInfo map[int]*DayReading) (*MonthReading, error) {
	return &MonthReading{
		MonthNum: month, 
		MonthReadingTime: "0min",
		MonthReadingTimeOfDifferentContent: make(map[string]string),
		MonthRawInfo: monthRawInfo,
	},nil
}

func (m *MonthReading) ComputeReadingTime() error {
	for _, dayReading := range m.MonthRawInfo {
		err := dayReading.ComputeReadingTime()
		if err != nil {
			return err
		}

		// 计算MonthReadingTimeOfDifferentContent
		for content, conReadingTime := range dayReading.DayReadingTimeOfDifferentContent {
			if _, ok := m.MonthReadingTimeOfDifferentContent[content]; !ok {
				m.MonthReadingTimeOfDifferentContent[content] = conReadingTime
			} else {
				conSum, err := hDate.FormatDurationSum(m.MonthReadingTimeOfDifferentContent[content], conReadingTime)
				if err != nil {
					return err 
				}
				m.MonthReadingTimeOfDifferentContent[content] = conSum
			}
		}
	  sum, err := hDate.FormatDurationSum(m.MonthReadingTime, dayReading.DayReadingTime)
		if err != nil {
			return err 
		}
		m.MonthReadingTime = sum
	}

	for k,v := range m.MonthReadingTimeOfDifferentContent {
	  m.MonthReadingTimeOfDifferentContentStr += fmt.Sprintf("%s: %s	", k, v)
	}

	return nil
}	


func (m *MonthReading) Print() {
	for content, conReadingTime := range m.MonthReadingTimeOfDifferentContent {
		klog.InfoS("month reading info", "monthNum", m.MonthNum, "readingTime", m.MonthReadingTime, "content", content, "contentReadingTime", conReadingTime)
	}
	
}
