package audiobook

import (
	"math"
	"strings"

	"k8s.io/klog/v2"
)

type MonthAudiobook struct {
	MonthNum string
	DaysInMonth int
	RawInfo map[int]*DayAudiobook
	TargetFinishBooks int

	IsFinish bool
	FinishBooks int
}

/*
math.Floor 向下取整，返回不大于参数的最大整数。
math.Ceil 向上取整，返回不小于参数的最小整数。
math.Round 四舍五入取整，返回最接近参数的整数。
使用类型转换（如int(num / 3.5)）会截断小数部分，这类似于向零取整。
*/
func NewMonthAudiobook(month string, rawInfo map[int]*DayAudiobook, daysInMonth int) (*MonthAudiobook, error) {
	return &MonthAudiobook{
		MonthNum: month, 
		RawInfo: rawInfo,
		DaysInMonth: daysInMonth,

		TargetFinishBooks: int(math.Ceil(float64(daysInMonth) / 3.5)),
	},nil
}


func (m *MonthAudiobook) ComputeFinishBooks() error {
  for _, item := range m.RawInfo {
		// 判断字符串是否以“(完)”结尾
		if strings.HasSuffix(item.RawInfo, "(完)") {
			m.FinishBooks++
		}
	}	
  return nil
}


func (m *MonthAudiobook) CheckFinish() error {
	if m.FinishBooks >= m.TargetFinishBooks {
		m.IsFinish = true
	}
	return nil
}

func (m *MonthAudiobook) Print()  {
  klog.InfoS("MonthAudiobook", "month", m.MonthNum, "finish", m.IsFinish, "finishBooks", m.FinishBooks, "targetFinishBooks", m.TargetFinishBooks)
}