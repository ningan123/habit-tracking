package audiobook

import (
	"strings"

	"k8s.io/klog/v2"
)

import (
	hDate "ningan.com/habit-tracking/pkg/date"
)


type WeekAudiobook struct {
	Week *hDate.Week
	RawInfo map[string]*DayAudiobook 
	TargetFinishBooks int

	IsFinish bool
	FinishBooks int
}

func NewWeekAudiobook(weekNum string, rawInfo map[string]*DayAudiobook) (*WeekAudiobook, error) {
  return &WeekAudiobook{
    Week: &hDate.Week{
			WeekNum: weekNum,
		},
		RawInfo: rawInfo,
		TargetFinishBooks: TargetWeekFinishBooks,
  }, nil
}


func (w *WeekAudiobook) ComputeFinishBooks() error {
  for _, item := range w.RawInfo {
		strList := strings.Split(item.RawInfo, ",")

		for _, str := range strList {
			// 判断字符串是否以“(完)”结尾
			if strings.HasSuffix(str, "(完)") {
				w.FinishBooks++
			}
		}  
	}
		
	return nil
}


func (w *WeekAudiobook) CheckFinish() error {
	if w.FinishBooks >= w.TargetFinishBooks {
		w.IsFinish = true
	}
	return nil
}

func (w *WeekAudiobook) Print() {
  klog.InfoS("WeekAudiobook" , "weekNum", w.Week.WeekNum, "finishBooks", w.FinishBooks, "targetFinishBooks", w.TargetFinishBooks, "isFinish", w.IsFinish)
}