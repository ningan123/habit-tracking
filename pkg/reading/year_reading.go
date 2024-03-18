package reading

import (
	"fmt"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type YearReading struct {
  YearNum string // 年份
  YearReadingTime string  // 年总阅读时长
	YearRawInfo map[string]*DayReading  
  YearReadingTimeOfDifferentContent map[string]string // 不同内容的阅读时间
	YearReadingTimeOfDifferentContentStr string	
}

func NewYearReading(yearNum string, yearRawInfo map[string]*DayReading) (*YearReading, error) {
  return &YearReading{
    YearRawInfo: yearRawInfo,
		YearNum: yearNum,
    YearReadingTimeOfDifferentContent: make(map[string]string),
    YearReadingTime: "0min",
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
