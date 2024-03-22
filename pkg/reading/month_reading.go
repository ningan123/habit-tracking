package reading

import (
	"fmt"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type MonthReading struct {
	MonthNum string
	ReadingTime string 
	DaysInMonth int
	ReadingTimeOfDifferentContent map[string]string // 不同内容的阅读时间
	ReadingTimeOfDifferentContentStr string
	RawInfo  map[int]*DayReading  // int表示几号
	IsFinish bool
	TargetReadingTime string
}


func NewMonthReading(month string, rawInfo map[int]*DayReading, daysInMonth int) (*MonthReading, error) {
	tReadingTime, err := hDate.FormatDurationMultiply(TargetDayReadingTime, daysInMonth)
	if err != nil {
		klog.Errorf("format duration error: %v", err)
		return nil, err
	}

	return &MonthReading{
		MonthNum: month, 
		ReadingTime: "0min",
		ReadingTimeOfDifferentContent: make(map[string]string),
		RawInfo: rawInfo,
		DaysInMonth: daysInMonth,
		TargetReadingTime: tReadingTime,
	},nil
}

func (m *MonthReading) ComputeReadingTime() error {
	for _, dayReading := range m.RawInfo {
		err := dayReading.ComputeReadingTime()
		if err != nil {
			return err
		}

		// 计算ReadingTimeOfDifferentContent
		for content, conReadingTime := range dayReading.ReadingTimeOfDifferentContent {
			if _, ok := m.ReadingTimeOfDifferentContent[content]; !ok {
				m.ReadingTimeOfDifferentContent[content] = conReadingTime
			} else {
				conSum, err := hDate.FormatDurationSum(m.ReadingTimeOfDifferentContent[content], conReadingTime)
				if err != nil {
					return err 
				}
				m.ReadingTimeOfDifferentContent[content] = conSum
			}
		}
	  sum, err := hDate.FormatDurationSum(m.ReadingTime, dayReading.ReadingTime)
		if err != nil {
			return err 
		}
		m.ReadingTime = sum
	}

	for k,v := range m.ReadingTimeOfDifferentContent {
	  m.ReadingTimeOfDifferentContentStr += fmt.Sprintf("%s: %s	", k, v)
	}

	return nil
}	


func (m *MonthReading) Print() {
	for content, conReadingTime := range m.ReadingTimeOfDifferentContent {
		klog.InfoS("month reading info", "monthNum", m.MonthNum, "readingTime", m.ReadingTime, "content", content, "contentReadingTime", conReadingTime)
	}
	
}


// 只要阅读时长>=target时长，就认为完成
func (m *MonthReading) CheckFinish() error {
  res, err :=  hDate.IsActualDurationLongerOrEqualToTargetDuration(m.ReadingTime, m.TargetReadingTime)
	if err != nil {
		return err
	}
	m.IsFinish = res
	return nil
}