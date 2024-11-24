package sport

import (
	"fmt"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type WeekSport struct {
	Week                            *hDate.Week
	SportTimes                      int
	SportTimesOfDifferentContent    map[string]int
	SportTimesOfDifferentContentStr string
	RawInfo                         map[string]*DaySport // string表示周几
	IsFinish                        bool
	TargetFinishDays                int
	ActualFinishDays                int
}

func NewWeekSport(weekNum string, rawInfo map[string]*DaySport) (*WeekSport, error) {
	return &WeekSport{
		Week: &hDate.Week{
			WeekNum: weekNum,
		},
		SportTimesOfDifferentContent: make(map[string]int),
		RawInfo:                      rawInfo,
		TargetFinishDays:             TargetWeekSportDays,
	}, nil
}

func (w *WeekSport) ComputeSportTimes() error {
	for _, DaySport := range w.RawInfo {
		err := DaySport.ComputeSportTimes()
		if err != nil {
			return err
		}

		// 计算sportTimeOfDifferentContent
		for content, sTimes := range DaySport.SportTimesOfDifferentContent {
			if _, ok := w.SportTimesOfDifferentContent[content]; !ok {
				// klog.InfoS("Week-Day contentsportTime", "date", DaySport.DayDate, "contentsportTime", consportTime)
				w.SportTimesOfDifferentContent[content] = sTimes
			} else {
				w.SportTimesOfDifferentContent[content] += sTimes
			}
			klog.V(2).InfoS("Week-Day contentsportTime", "weekNum", w.Week.WeekNum, "date", DaySport.Day.Date, "sportTimeOfDifferentContent", w.SportTimesOfDifferentContent)
		}

		w.SportTimes += DaySport.SportTimes
	}

	for k, v := range w.SportTimesOfDifferentContent {
		w.SportTimesOfDifferentContentStr += fmt.Sprintf("%s: %d<br>", k, v)
	}

	return nil
}

func (w *WeekSport) Print() {
	for content, consportTime := range w.SportTimesOfDifferentContent {
		klog.InfoS("week sport info", "weekNum", w.Week.WeekNum, "content", content, "contentsportTime", consportTime)
	}
}

func (w *WeekSport) CheckFinish() error {
	for _, v := range w.RawInfo {
		err := v.CheckFinish()
		if err != nil {
			return err
		}
		if v.IsFinish {
			w.ActualFinishDays++
		}
	}

	if w.ActualFinishDays >= w.TargetFinishDays {
		w.IsFinish = true
	}

	return nil
}
