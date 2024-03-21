package piano

import (
	"fmt"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type WeekPiano struct {
	WeekNum string
	PianoTime string 
	PianoTimeOfDifferentContent map[string]string 
	PianoTimeOfDifferentContentStr string
	RawInfo  map[string]*DayPiano  // string表示周几
	IsFinish bool
	TargetPianoTime string
	ExtraPianoTime string
}

func NewWeekPiano(weekNum string, rawInfo map[string]*DayPiano ) (*WeekPiano, error) {
	tPianoTime, err := hDate.FormatDurationMultiply(TargetDayPianoTime, 7)
	if err != nil {
		klog.Errorf("format duration error: %v", err)
		return nil, err
	}

  return &WeekPiano{
    WeekNum: weekNum,
		PianoTime: "0min",
		PianoTimeOfDifferentContent : make(map[string]string),
		RawInfo: rawInfo,
		TargetPianoTime: tPianoTime,
  }, nil
}

func (w *WeekPiano) ComputePianoTime() error {
	for _, dayPiano := range w.RawInfo {
		err := dayPiano.ComputePianoTime()
		if err != nil {
			return err
		}

		// 计算PianoTimeOfDifferentContent
		for content, conPianoTime := range dayPiano.PianoTimeOfDifferentContent {
			if _, ok := w.PianoTimeOfDifferentContent[content]; !ok {
				// klog.InfoS("Week-Day contentPianoTime", "date", dayPiano.DayDate, "contentPianoTime", conPianoTime)
				w.PianoTimeOfDifferentContent[content] = conPianoTime
			} else {
				conSum, err := hDate.FormatDurationSum(w.PianoTimeOfDifferentContent[content], conPianoTime)
				if err != nil {
					return err 
				}
				w.PianoTimeOfDifferentContent[content] = conSum
			}
			// klog.InfoS("Week-Day contentPianoTime", "PianoTimeOfDifferentContent", w.PianoTimeOfDifferentContent)
		}

		// 计算PianoTime
	  sum, err := hDate.FormatDurationSum(w.PianoTime, dayPiano.PianoTime)
		if err != nil {
			return err 
		}
		w.PianoTime = sum		
	}

	for k, v := range w.PianoTimeOfDifferentContent {
		w.PianoTimeOfDifferentContentStr += fmt.Sprintf("%s: %s	", k, v)
	}
	
	return nil
}

func (w *WeekPiano) Print() {
	for content, conPianoTime := range w.PianoTimeOfDifferentContent {
		klog.InfoS("week piano info", "weekNum", w.WeekNum, "pianoTime", w.PianoTime, "content", content, "contentPianoTime", conPianoTime)
	}
}


// 只要阅读时长>=target时长，就认为完成
func (w *WeekPiano) CheckFinish() error {
  res, err :=  hDate.IsActualDurationLongerOrEqualToTargetDuration(w.PianoTime, w.TargetPianoTime)
	if err != nil {
		return err
	}
	// klog.InfoS("week piano info", "weekNum", w.WeekNum, "pianoTime", w.PianoTime, "targetPianoTime", w.TargetPianoTime, "isFinish", res)
	w.IsFinish = res
	return nil
}

func (w *WeekPiano) ComputeExtraPianoTime() error {
  sub, err := hDate.FormatDurationSub(w.PianoTime, w.TargetPianoTime)
	if err != nil {
	  return err
	}
	w.ExtraPianoTime = sub
	return nil
}