package reading

import (
	"fmt"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type YearReading struct {
	Year *hDate.Year
  ReadingTime string  // 年总阅读时长
	YearRawInfo map[string]*DayReading  
  ReadingTimeOfDifferentContent map[string]string // 不同内容的阅读时间
	ReadingTimeOfDifferentContentStr string	
	IsFinish bool
	TargetReadingTime string
}

func NewYearReading(yearNum string, daysInYear int, yearRawInfo map[string]*DayReading) (*YearReading, error) {
	tReadingTime, err := hDate.FormatDurationMultiply(TargetDayReadingTime, daysInYear)
	if err != nil {
		klog.Errorf("format duration error: %v", err)
		return nil, err
	}

  return &YearReading{
		Year: &hDate.Year{
			YearNum: yearNum,
			DaysInYear: daysInYear,
		},
    YearRawInfo: yearRawInfo,
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
		klog.InfoS("year reading info", "yearNum", y.Year.YearNum, "readingTime", y.ReadingTime, "content", content, "contentReadingTime", conReadingTime)
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