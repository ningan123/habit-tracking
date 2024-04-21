package sport

import (
	"fmt"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type YearSport struct {
	Year                            *hDate.Year
	RawInfo                         map[string]*DaySport
	SportTimes                      int
	SportTimesOfDifferentContent    map[string]int
	SportTimesOfDifferentContentStr string
	IsFinish                        bool
	TargetFinishDays                int
	ActualFinishDays                int
}

func NewYearSport(yearNum string, daysInYear int, yearRawInfo map[string]*DaySport) (*YearSport, error) {
	return &YearSport{
		Year: &hDate.Year{
			YearNum:    yearNum,
			DaysInYear: daysInYear,
		},
		RawInfo:                      yearRawInfo,
		SportTimesOfDifferentContent: make(map[string]int),
		TargetFinishDays:             TargetYearSportDays,
	}, nil
}

func (y *YearSport) ComputeSportTimes() error {
	for _, DaySport := range y.RawInfo {
		err := DaySport.ComputeSportTimes()
		if err != nil {
			return err
		}

		// 计算SportTimesOfDifferentContent
		for content, sTime := range DaySport.SportTimesOfDifferentContent {
			if _, ok := y.SportTimesOfDifferentContent[content]; !ok {
				y.SportTimesOfDifferentContent[content] = sTime
			} else {
				y.SportTimesOfDifferentContent[content] += sTime
			}
		}

		y.SportTimes += DaySport.SportTimes
	}

	for k, v := range y.SportTimesOfDifferentContent {
		y.SportTimesOfDifferentContentStr += fmt.Sprintf("%s: %d<br>", k, v)
	}

	return nil
}

func (y *YearSport) Print() {
	for content, consportTime := range y.SportTimesOfDifferentContent {
		klog.InfoS("year sport info", "yearNum", y.Year.YearNum, "sportTimes", y.SportTimes, "content", content, "contentsportTime", consportTime)
	}
}

func (y *YearSport) CheckFinish() error {
	for _, v := range y.RawInfo {
		err := v.CheckFinish()
		if err != nil {
			return err
		}
		if v.IsFinish {
			y.ActualFinishDays++
		}
	}

	if y.ActualFinishDays >= y.TargetFinishDays {
		y.IsFinish = true
	}
	return nil
}
