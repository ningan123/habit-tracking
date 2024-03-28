package audiobook

import (
	"strings"

	"k8s.io/klog/v2"
)

type WeekAudiobook struct {
	WeekNum string
	RawInfo map[string]*DayAudiobook 
	TargetFinishBooks int

	IsFinish bool
	FinishBooks int
}

func NewWeekAudiobook(weekNum string, rawInfo map[string]*DayAudiobook) (*WeekAudiobook, error) {
  return &WeekAudiobook{
    WeekNum: weekNum,
		RawInfo: rawInfo,

		TargetFinishBooks: TargetWeekFinishBooks,
  }, nil
}


func (w *WeekAudiobook) ComputeFinishBooks() error {
  for _, item := range w.RawInfo {
		// 判断字符串是否以“(完)”结尾
		if strings.HasSuffix(item.RawInfo, "(完)") {
			w.FinishBooks++
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
  klog.InfoS("WeekAudiobook" , "weekNum", w.WeekNum, "finishBooks", w.FinishBooks, "targetFinishBooks", w.TargetFinishBooks, "isFinish", w.IsFinish)
}