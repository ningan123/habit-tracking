package reading

import (
	"fmt"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type WeekReading struct {
	WeekNum string
	WeekReadingTime string 
	WeekReadingTimeOfDifferentContent map[string]string 
	WeekReadingTimeOfDifferentContentStr string
	WeekRawInfo  map[string]*DayReading  // string表示周几
	IsFinish bool
	TargetReadingTime string
	ExtraReadingTime string
}

func NewWeekReading(weekNum string, weekRawInfo map[string]*DayReading ) (*WeekReading, error) {
	tReadingTime, err := hDate.FormatDurationMultiply(TargetDayReadingTime, 7)
	if err != nil {
		klog.Errorf("format duration error: %v", err)
		return nil, err
	}

  return &WeekReading{
    WeekNum: weekNum,
		WeekReadingTime: "0min",
		WeekReadingTimeOfDifferentContent : make(map[string]string),
		WeekRawInfo: weekRawInfo,
		TargetReadingTime: tReadingTime,
  }, nil
}

func (w *WeekReading) ComputeReadingTime() error {
	for _, dayReading := range w.WeekRawInfo {
		err := dayReading.ComputeReadingTime()
		if err != nil {
			return err
		}

		// 计算WeekReadingTimeOfDifferentContent
		for content, conReadingTime := range dayReading.ReadingTimeOfDifferentContent {
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
	  sum, err := hDate.FormatDurationSum(w.WeekReadingTime, dayReading.ReadingTime)
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


// 只要阅读时长>=target时长，就认为完成
func (w *WeekReading) CheckFinish() error {
  res, err :=  hDate.IsActualDurationLongerOrEqualToTargetDuration(w.WeekReadingTime, w.TargetReadingTime)
	if err != nil {
		return err
	}
	// klog.InfoS("week reading info", "weekNum", w.WeekNum, "readingTime", w.WeekReadingTime, "targetReadingTime", w.TargetReadingTime, "isFinish", res)
	w.IsFinish = res
	return nil
}

func (w *WeekReading) ComputeExtraReadingTime() error {
  sub, err := hDate.FormatDurationSub(w.WeekReadingTime, w.TargetReadingTime)
	if err != nil {
	  return err
	}
	w.ExtraReadingTime = sub
	return nil
}