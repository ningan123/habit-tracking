package audiobook

import (
	"math"
	"strings"

	"k8s.io/klog/v2"
)

import (
	hDate "ningan.com/habit-tracking/pkg/date"
)



type YearAudiobook struct {
  Year *hDate.Year
	RawInfo map[string]*DayAudiobook
	TargetFinishBooks int
	IsFinish bool
	FinishBooks int
}

func NewYearAudiobook(yearNum string, rawInfo map[string]*DayAudiobook, daysInYear int) (*YearAudiobook, error) {

  return &YearAudiobook{
		Year: &hDate.Year{
			YearNum: yearNum,
			DaysInYear: daysInYear,
		},
    RawInfo: rawInfo,
		TargetFinishBooks: int(math.Ceil(float64(daysInYear) / 3.5)),
  },nil
}

func (y *YearAudiobook) ComputeFinishBooks() error {
  for _, item := range y.RawInfo {
		strList := strings.Split(item.RawInfo, ",")

		for _, str := range strList {
			// 判断字符串是否以“(完)”结尾
			if strings.HasSuffix(str, "(完)") {
				y.FinishBooks++
			}
		} 
	}	
  return nil
}


func (y *YearAudiobook) CheckFinish() error {
	if y.FinishBooks >= y.TargetFinishBooks {
		y.IsFinish = true
	}
	return nil
}

func (y *YearAudiobook) Print() {
	klog.InfoS("YearAudiobook" , "yearNum", y.Year.YearNum, "daysInYear", y.Year.DaysInYear, "targetFinishBooks", y.TargetFinishBooks, "finishBooks", y.FinishBooks, "isFinish", y.IsFinish)
}