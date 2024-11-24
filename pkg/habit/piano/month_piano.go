package piano

import (
	"fmt"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
	hString "ningan.com/habit-tracking/pkg/string"
)

type MonthPiano struct {
	Month                          *hDate.Month
	PianoTime                      string
	PianoTimeOfDifferentContent    map[string]string
	PianoTimeOfDifferentContentStr string
	RawInfo                        map[int]*DayPiano // int表示几号
	IsFinish                       bool
	TargetPianoTime                string
	ExtraPianoTime                 string
}

func NewMonthPiano(monthNum string, daysInMonth int, rawInfo map[int]*DayPiano) (*MonthPiano, error) {
	tPianoTime, err := hDate.FormatDurationMultiply(TargetDayPianoTime, daysInMonth)
	if err != nil {
		klog.Errorf("format duration error: %v", err)
		return nil, err
	}

	return &MonthPiano{
		Month: &hDate.Month{
			MonthNum:    monthNum,
			DaysInMonth: daysInMonth,
		},
		PianoTime:                   "0min",
		PianoTimeOfDifferentContent: make(map[string]string),
		RawInfo:                     rawInfo,
		TargetPianoTime:             tPianoTime,
	}, nil
}

func (m *MonthPiano) ComputePianoTime() error {
	for _, dayPiano := range m.RawInfo {
		err := dayPiano.ComputePianoTime()
		if err != nil {
			return err
		}

		// 计算PianoTimeOfDifferentContent
		for content, conPianoTime := range dayPiano.PianoTimeOfDifferentContent {
			if _, ok := m.PianoTimeOfDifferentContent[content]; !ok {
				m.PianoTimeOfDifferentContent[content] = conPianoTime
			} else {
				conSum, err := hDate.FormatDurationSum(m.PianoTimeOfDifferentContent[content], conPianoTime)
				if err != nil {
					return err
				}
				m.PianoTimeOfDifferentContent[content] = conSum
			}
		}
		sum, err := hDate.FormatDurationSum(m.PianoTime, dayPiano.PianoTime)
		if err != nil {
			return err
		}
		m.PianoTime = sum
	}

	for k, v := range m.PianoTimeOfDifferentContent {
		m.PianoTimeOfDifferentContentStr += fmt.Sprintf("%s: %s<br>", k, v)
	}
	m.PianoTimeOfDifferentContentStr = hString.SortString(m.PianoTimeOfDifferentContentStr)

	return nil
}

func (m *MonthPiano) Print() {
	for content, conPianoTime := range m.PianoTimeOfDifferentContent {
		klog.InfoS("month piano info", "monthNum", m.Month.MonthNum, "pianoTime", m.PianoTime, "content", content, "contentPianoTime", conPianoTime)
	}

}

// 只要时长>=target时长，就认为完成
func (m *MonthPiano) CheckFinish() error {
	res, err := hDate.IsActualDurationLongerOrEqualToTargetDuration(m.PianoTime, m.TargetPianoTime)
	if err != nil {
		return err
	}
	m.IsFinish = res
	return nil
}

func (m *MonthPiano) ComputeExtraPianoTime() error {
	sub, err := hDate.FormatDurationSub(m.PianoTime, m.TargetPianoTime)
	if err != nil {
		return err
	}
	m.ExtraPianoTime = sub
	return nil
}
