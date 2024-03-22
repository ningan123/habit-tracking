package reading

import (
	"fmt"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type WeekReading struct {
	WeekNum string
	ReadingTime string 
	ReadingTimeOfDifferentContent map[string]string 
	ReadingTimeOfDifferentContentStr string
	RawInfo  map[string]*DayReading  // string表示周几
	IsFinish bool
	TargetReadingTime string
	ExtraReadingTime string
}

func NewWeekReading(weekNum string, rawInfo map[string]*DayReading ) (*WeekReading, error) {
	tReadingTime, err := hDate.FormatDurationMultiply(TargetDayReadingTime, 7)
	if err != nil {
		klog.Errorf("format duration error: %v", err)
		return nil, err
	}

  return &WeekReading{
    WeekNum: weekNum,
		ReadingTime: "0min",
		ReadingTimeOfDifferentContent : make(map[string]string),
		RawInfo: rawInfo,
		TargetReadingTime: tReadingTime,
  }, nil
}

func (w *WeekReading) ComputeReadingTime() error {
	for _, dayReading := range w.RawInfo {
		err := dayReading.ComputeReadingTime()
		if err != nil {
			return err
		}

		// 计算ReadingTimeOfDifferentContent
		for content, conReadingTime := range dayReading.ReadingTimeOfDifferentContent {
			if _, ok := w.ReadingTimeOfDifferentContent[content]; !ok {
				// klog.InfoS("Week-Day contentReadingTime", "date", dayReading.DayDate, "contentReadingTime", conReadingTime)
				w.ReadingTimeOfDifferentContent[content] = conReadingTime
			} else {
				conSum, err := hDate.FormatDurationSum(w.ReadingTimeOfDifferentContent[content], conReadingTime)
				if err != nil {
					return err 
				}
				w.ReadingTimeOfDifferentContent[content] = conSum
			}
			// klog.InfoS("Week-Day contentReadingTime", "ReadingTimeOfDifferentContent", w.ReadingTimeOfDifferentContent)
		}

		// 计算ReadingTime
	  sum, err := hDate.FormatDurationSum(w.ReadingTime, dayReading.ReadingTime)
		if err != nil {
			return err 
		}
		w.ReadingTime = sum		
	}

	for k, v := range w.ReadingTimeOfDifferentContent {
		w.ReadingTimeOfDifferentContentStr += fmt.Sprintf("%s: %s	", k, v)
	}
	
	return nil
}

func (w *WeekReading) Print() {
	for content, conReadingTime := range w.ReadingTimeOfDifferentContent {
		klog.InfoS("week reading info", "weekNum", w.WeekNum, "readingTime", w.ReadingTime, "content", content, "contentReadingTime", conReadingTime)
	}
}


// 只要阅读时长>=target时长，就认为完成
func (w *WeekReading) CheckFinish() error {
  res, err :=  hDate.IsActualDurationLongerOrEqualToTargetDuration(w.ReadingTime, w.TargetReadingTime)
	if err != nil {
		return err
	}
	// klog.InfoS("week reading info", "weekNum", w.WeekNum, "readingTime", w.ReadingTime, "targetReadingTime", w.TargetReadingTime, "isFinish", res)
	w.IsFinish = res
	return nil
}

func (w *WeekReading) ComputeExtraReadingTime() error {
  sub, err := hDate.FormatDurationSub(w.ReadingTime, w.TargetReadingTime)
	if err != nil {
	  return err
	}
	w.ExtraReadingTime = sub
	return nil
}