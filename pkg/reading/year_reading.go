package reading

import (
	"fmt"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type YearReading struct {
  YearNum string // 年份
	DaysInYear int // 一年多少天
  YearReadingTime string  // 年总阅读时长
	YearRawInfo map[string]*DayReading  
  YearReadingTimeOfDifferentContent map[string]string // 不同内容的阅读时间
	YearReadingTimeOfDifferentContentStr string	
	IsFinish bool
	TargetReadingTime string
}

func NewYearReading(yearNum string, yearRawInfo map[string]*DayReading, daysInYear int) (*YearReading, error) {
	tReadingTime, err := hDate.FormatDurationMultiply(TargetDayReadingTime, daysInYear)
	if err != nil {
		klog.Errorf("format duration error: %v", err)
		return nil, err
	}

  return &YearReading{
    YearRawInfo: yearRawInfo,
		YearNum: yearNum,
    YearReadingTimeOfDifferentContent: make(map[string]string),
    YearReadingTime: "0min",
		TargetReadingTime: tReadingTime,
  },nil
}


func (y *YearReading) ComputeReadingTime() error {
	for _, dayReading := range y.YearRawInfo {
		err := dayReading.ComputeReadingTime()
		if err != nil {
			return err
		}

		// 计算YearReadingTimeOfDifferentContent
		for content, conReadingTime := range dayReading.DayReadingTimeOfDifferentContent {
			if _, ok := y.YearReadingTimeOfDifferentContent[content]; !ok {
				y.YearReadingTimeOfDifferentContent[content] = conReadingTime
			} else {
				conSum, err := hDate.FormatDurationSum(y.YearReadingTimeOfDifferentContent[content], conReadingTime)
				if err != nil {
					return err 
				}
				y.YearReadingTimeOfDifferentContent[content] = conSum
			}
		}
	  sum, err := hDate.FormatDurationSum(y.YearReadingTime, dayReading.DayReadingTime)
		if err != nil {
			return err 
		}
		y.YearReadingTime = sum
	}

	for k, v := range y.YearReadingTimeOfDifferentContent {
		y.YearReadingTimeOfDifferentContentStr += fmt.Sprintf("%s: %s	", k, v)
	}

	return nil
}	



func (y *YearReading) Print() {
	for content, conReadingTime := range y.YearReadingTimeOfDifferentContent {
		klog.InfoS("year reading info", "yearNum", y.YearNum, "readingTime", y.YearReadingTime, "content", content, "contentReadingTime", conReadingTime)
	}	
}


// 只要阅读时长>=target时长，就认为完成
func (y *YearReading) CheckFinish() error {
  res, err :=  hDate.IsActualDurationLongerOrEqualToTargetDuration(y.YearReadingTime, y.TargetReadingTime)
	if err != nil {
		return err
	}
	y.IsFinish = res
	return nil
}