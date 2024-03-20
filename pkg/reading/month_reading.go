package reading

import (
	"fmt"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type MonthReading struct {
	MonthNum string
	MonthReadingTime string 
	DaysInMonth int
	MonthReadingTimeOfDifferentContent map[string]string // 不同内容的阅读时间
	MonthReadingTimeOfDifferentContentStr string
	MonthRawInfo  map[int]*DayReading  // int表示几号
	IsFinish bool
	TargetReadingTime string
}


func NewMonthReading(month string, monthRawInfo map[int]*DayReading, daysInMonth int) (*MonthReading, error) {
	tReadingTime, err := hDate.FormatDurationMultiply(TargetDayReadingTime, daysInMonth)
	if err != nil {
		klog.Errorf("format duration error: %v", err)
		return nil, err
	}

	return &MonthReading{
		MonthNum: month, 
		MonthReadingTime: "0min",
		MonthReadingTimeOfDifferentContent: make(map[string]string),
		MonthRawInfo: monthRawInfo,
		DaysInMonth: daysInMonth,
		TargetReadingTime: tReadingTime,
	},nil
}

func (m *MonthReading) ComputeReadingTime() error {
	for _, dayReading := range m.MonthRawInfo {
		err := dayReading.ComputeReadingTime()
		if err != nil {
			return err
		}

		// 计算MonthReadingTimeOfDifferentContent
		for content, conReadingTime := range dayReading.ReadingTimeOfDifferentContent {
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
	  sum, err := hDate.FormatDurationSum(m.MonthReadingTime, dayReading.ReadingTime)
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


// 只要阅读时长>=target时长，就认为完成
func (m *MonthReading) CheckFinish() error {
  res, err :=  hDate.IsActualDurationLongerOrEqualToTargetDuration(m.MonthReadingTime, m.TargetReadingTime)
	if err != nil {
		return err
	}
	m.IsFinish = res
	return nil
}