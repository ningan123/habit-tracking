package audiobook

import (
	"strings"

	"k8s.io/klog/v2"

	hDate "ningan.com/habit-tracking/pkg/date"
)

type DayAudiobook struct {
	Day *hDate.Day
  RawInfo  string
	FinishBooks int // 完成书籍数量
	IsFinish bool
}




func NewDayAudiobook(date string, weekday string, weekNum string, monthNum string, yearNum string, dayOfMonth int, dayOfYear int, rawInfo string) (*DayAudiobook, error) {
	return &DayAudiobook{
		Day: &hDate.Day{
			Date: date,
			Weekday: weekday,
			WeekNum: weekNum,
			MonthNum: monthNum,
			YearNum: yearNum,
			DayOfMonth: dayOfMonth,
			DayOfYear: dayOfYear,
		},
		RawInfo: rawInfo,
	}, nil
}


func (d *DayAudiobook) ComputeFinishBooks() error {
	strList := strings.Split(d.RawInfo, ",")

	for _, str := range strList {
		// 判断字符串是否以“(完)”结尾
		if strings.HasSuffix(str, "(完)") {
			d.FinishBooks++
		}
	}

  return nil
}


// 只要阅读时长>=target时长，就认为完成
func (d *DayAudiobook) CheckFinish() error {
	if d.RawInfo == "" || d.RawInfo == "×" {
	  return nil
	}

	d.IsFinish = true
	return nil
}


func (d *DayAudiobook) Print() {
  klog.InfoS("DayAudiobook", "date", d.Day.Date, "rawInfo", d.RawInfo, "finishBooks", d.FinishBooks, "isFinish", d.IsFinish)
}