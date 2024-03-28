package audiobook

import (
	"math"
	"strings"

	"k8s.io/klog/v2"
)

type YearAudiobook struct {
  YearNum string // 年份
	DaysInYear int // 一年多少天
	RawInfo map[string]*DayAudiobook
	TargetFinishBooks int
	  
	IsFinish bool
	FinishBooks int
}

func NewYearAudiobook(yearNum string, rawInfo map[string]*DayAudiobook, daysInYear int) (*YearAudiobook, error) {

  return &YearAudiobook{
    RawInfo: rawInfo,
		YearNum: yearNum,
		DaysInYear: daysInYear,
    
		TargetFinishBooks: int(math.Ceil(float64(daysInYear) / 3.5)),
  },nil
}

func (y *YearAudiobook) ComputeFinishBooks() error {
  for _, item := range y.RawInfo {
		// 判断字符串是否以“(完)”结尾
		if strings.HasSuffix(item.RawInfo, "(完)") {
			y.FinishBooks++
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
	klog.InfoS("YearAudiobook" , "yearNum", y.YearNum, "daysInYear", y.DaysInYear, "targetFinishBooks", y.TargetFinishBooks, "finishBooks", y.FinishBooks, "isFinish", y.IsFinish)
}