package reading

import (
	"fmt"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type YearReading struct {
  YearNum string // 年份
	DaysInYear int // 一年多少天
  ReadingTime string  // 年总阅读时长
	YearRawInfo map[string]*DayReading  
  ReadingTimeOfDifferentContent map[string]string // 不同内容的阅读时间
	ReadingTimeOfDifferentContentStr string	
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
    ReadingTimeOfDifferentContent: make(map[string]string),
    ReadingTime: "0min",
		TargetReadingTime: tReadingTime,
  },nil
}


func (y *YearReading) ComputeReadingTime() error {
	for _, dayReading := range y.YearRawInfo {
		err := dayReading.ComputeReadingTime()
		if err != nil {
			return err
		}

		// 计算ReadingTimeOfDifferentContent
		for content, conReadingTime := range dayReading.ReadingTimeOfDifferentContent {
			if _, ok := y.ReadingTimeOfDifferentContent[content]; !ok {
				y.ReadingTimeOfDifferentContent[content] = conReadingTime
			} else {
				conSum, err := hDate.FormatDurationSum(y.ReadingTimeOfDifferentContent[content], conReadingTime)
				if err != nil {
					return err 
				}
				y.ReadingTimeOfDifferentContent[content] = conSum
			}
		}
	  sum, err := hDate.FormatDurationSum(y.ReadingTime, dayReading.ReadingTime)
		if err != nil {
			return err 
		}
		y.ReadingTime = sum
	}

	for k, v := range y.ReadingTimeOfDifferentContent {
		y.ReadingTimeOfDifferentContentStr += fmt.Sprintf("%s: %s	", k, v)
	}

	return nil
}	



func (y *YearReading) Print() {
	for content, conReadingTime := range y.ReadingTimeOfDifferentContent {
		klog.InfoS("year reading info", "yearNum", y.YearNum, "readingTime", y.ReadingTime, "content", content, "contentReadingTime", conReadingTime)
	}	
}


// 只要阅读时长>=target时长，就认为完成
func (y *YearReading) CheckFinish() error {
  res, err :=  hDate.IsActualDurationLongerOrEqualToTargetDuration(y.ReadingTime, y.TargetReadingTime)
	if err != nil {
		return err
	}
	y.IsFinish = res
	return nil
}