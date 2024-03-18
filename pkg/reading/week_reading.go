package reading

import (
	"fmt"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type WeekReading struct {
	WeekNum int
	WeekReadingTime string 
	WeekReadingTimeOfDifferentContent map[string]string 
	WeekReadingTimeOfDifferentContentStr string
	WeekRawInfo  map[string]*DayReading  // string表示周几
}

func NewWeekReading(weekNum int, weekRawInfo map[string]*DayReading ) (*WeekReading, error) {
  return &WeekReading{
    WeekNum: weekNum,
		WeekReadingTime: "0min",
		WeekReadingTimeOfDifferentContent : make(map[string]string),
		WeekRawInfo: weekRawInfo,
  }, nil
}

func (w *WeekReading) ComputeReadingTime() error {
	for _, dayReading := range w.WeekRawInfo {
		err := dayReading.ComputeReadingTime()
		if err != nil {
			return err
		}

		// 计算WeekReadingTimeOfDifferentContent
		for content, conReadingTime := range dayReading.DayReadingTimeOfDifferentContent {
			if _, ok := w.WeekReadingTimeOfDifferentContent[content]; !ok {
				// klog.InfoS("Week-Day contentReadingTime", "date", dayReading.DayDate, "contentReadingTime", conReadingTime)
				w.WeekReadingTimeOfDifferentContent[content] = conReadingTime
			} else {
				conSum, err := hDate.FormatDurationSum(w.WeekReadingTimeOfDifferentContent[content], conReadingTime)
				if err != nil {
					return err 
				}
				w.WeekReadingTimeOfDifferentContent[content] = conSum
			}
			// klog.InfoS("Week-Day contentReadingTime", "WeekReadingTimeOfDifferentContent", w.WeekReadingTimeOfDifferentContent)
		}

		// 计算WeekReadingTime
	  sum, err := hDate.FormatDurationSum(w.WeekReadingTime, dayReading.DayReadingTime)
		if err != nil {
			return err 
		}
		w.WeekReadingTime = sum
	}

	for k, v := range w.WeekReadingTimeOfDifferentContent {
		w.WeekReadingTimeOfDifferentContentStr += fmt.Sprintf("%s: %s	", k, v)
	}
	
	return nil
}

func (w *WeekReading) Print() {
	for content, conReadingTime := range w.WeekReadingTimeOfDifferentContent {
		klog.InfoS("week reading info", "weekNum", w.WeekNum, "readingTime", w.WeekReadingTime, "content", content, "contentReadingTime", conReadingTime)
	}
}