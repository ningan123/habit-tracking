package piano

import (
	"fmt"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type YearPiano struct {
	Year *hDate.Year
  PianoTime string   
	RawInfo map[string]*DayPiano  
  PianoTimeOfDifferentContent map[string]string
	PianoTimeOfDifferentContentStr string	
	IsFinish bool
	TargetPianoTime string
}

func NewYearPiano(yearNum string, daysInYear int, rawInfo map[string]*DayPiano) (*YearPiano, error) {
	tPianoTime, err := hDate.FormatDurationMultiply(TargetDayPianoTime, daysInYear)
	if err != nil {
		klog.Errorf("format duration error: %v", err)
		return nil, err
	}

  return &YearPiano{
		Year: &hDate.Year{
			YearNum: yearNum,
			DaysInYear: daysInYear,
		},
    RawInfo: rawInfo,
    PianoTimeOfDifferentContent: make(map[string]string),
    PianoTime: "0min",
		TargetPianoTime: tPianoTime,
  },nil
}


func (y *YearPiano) ComputePianoTime() error {
	for _, dayPiano := range y.RawInfo {
		err := dayPiano.ComputePianoTime()
		if err != nil {
			return err
		}

		// 计算PianoTimeOfDifferentContent
		for content, conPianoTime := range dayPiano.PianoTimeOfDifferentContent {
			if _, ok := y.PianoTimeOfDifferentContent[content]; !ok {
				y.PianoTimeOfDifferentContent[content] = conPianoTime
			} else {
				conSum, err := hDate.FormatDurationSum(y.PianoTimeOfDifferentContent[content], conPianoTime)
				if err != nil {
					return err 
				}
				y.PianoTimeOfDifferentContent[content] = conSum
			}
		}
	  sum, err := hDate.FormatDurationSum(y.PianoTime, dayPiano.PianoTime)
		if err != nil {
			return err 
		}
		y.PianoTime = sum
	}

	for k, v := range y.PianoTimeOfDifferentContent {
		y.PianoTimeOfDifferentContentStr += fmt.Sprintf("%s: %s	", k, v)
	}

	return nil
}	



func (y *YearPiano) Print() {
	for content, conPianoTime := range y.PianoTimeOfDifferentContent {
		klog.InfoS("year piano info", "yearNum", y.Year.YearNum, "pianoTime", y.PianoTime, "content", content, "contentPianoTime", conPianoTime)
	}	
}


// 只要阅读时长>=target时长，就认为完成
func (y *YearPiano) CheckFinish() error {
  res, err :=  hDate.IsActualDurationLongerOrEqualToTargetDuration(y.PianoTime, y.TargetPianoTime)
	if err != nil {
		return err
	}
	y.IsFinish = res
	return nil
}