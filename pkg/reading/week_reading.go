package reading

import (
	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type WeekReading struct {
	WeekNum int
	WeekReadingTime string 
	WeekRawInfo  map[string]*DayReading  // string表示周几
}

func NewWeekReading(weekNum int, weekRawInfo map[string]*DayReading ) (*WeekReading, error) {
  return &WeekReading{
    WeekNum: weekNum,
		WeekReadingTime: "0min",
		WeekRawInfo: weekRawInfo,
  }, nil
}

func (w *WeekReading) ComputeReadingTime() error {
	for _, dayReading := range w.WeekRawInfo {
	  sum, err := hDate.FormatDurationSum(w.WeekReadingTime, dayReading.DayReadingTime)
		if err != nil {
			return err 
		}
		w.WeekReadingTime = sum
	}
	 return nil
}

func (w *WeekReading) Print() {
  klog.InfoS("week reading info", "weekNum", w.WeekNum, "readingTime", w.WeekReadingTime)
}