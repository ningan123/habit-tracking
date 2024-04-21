package sport

import (
	"fmt"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type MonthSport struct {
	Month                           *hDate.Month
	SportTimes                      int
	SportTimesOfDifferentContent    map[string]int
	SportTimesOfDifferentContentStr string
	RawInfo                         map[int]*DaySport // int表示几号
	IsFinish                        bool
	TargetFinishDays                int
	ActualFinishDays                int
}

func NewMonthSport(monthNum string, daysInMonth int, rawInfo map[int]*DaySport) (*MonthSport, error) {
	return &MonthSport{
		Month: &hDate.Month{
			MonthNum:    monthNum,
			DaysInMonth: daysInMonth,
		},
		SportTimesOfDifferentContent: make(map[string]int),
		RawInfo:                      rawInfo,
		TargetFinishDays:             TargetMonthSportDays,
	}, nil
}

func (m *MonthSport) ComputeSportTimes() error {
	for _, DaySport := range m.RawInfo {
		err := DaySport.ComputeSportTimes()
		if err != nil {
			return err
		}

		// 计算sportTimeOfDifferentContent
		for content, sTimes := range DaySport.SportTimesOfDifferentContent {
			if _, ok := m.SportTimesOfDifferentContent[content]; !ok {
				m.SportTimesOfDifferentContent[content] = sTimes
			} else {
				m.SportTimesOfDifferentContent[content] += sTimes
			}

		}

		m.SportTimes += DaySport.SportTimes
	}

	for k, v := range m.SportTimesOfDifferentContent {
		m.SportTimesOfDifferentContentStr += fmt.Sprintf("%s: %d<br>", k, v)
	}

	return nil
}

func (m *MonthSport) Print() {
	for content, consportTime := range m.SportTimesOfDifferentContent {
		klog.InfoS("month sport info", "monthNum", m.Month.MonthNum, "content", content, "contentsportTime", consportTime)
	}
}

func (m *MonthSport) CheckFinish() error {
	for _, v := range m.RawInfo {
		err := v.CheckFinish()
		if err != nil {
			return err
		}
		if v.IsFinish {
			m.ActualFinishDays++
		}
	}

	if m.ActualFinishDays >= m.TargetFinishDays {
		m.IsFinish = true
	}

	return nil
}
