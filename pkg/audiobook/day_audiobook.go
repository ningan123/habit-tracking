package audiobook

import (
	"strings"
	"time"

	"k8s.io/klog/v2"
)

type DayAudiobook struct {
  RawInfo  string
	Date string // 具体日期
	Weekday string // 星期几
	Month time.Month // 几月
	WeekNum string // 几周
	Year int // 哪一年
	DayOfYear int // 一年中的第几天
	DayOfMonth int // 一个月中的第几天
	FinishBooks int // 完成书籍数量

	IsFinish bool
}




func NewDayAudiobook(date string, year int, dayOfYear int, month time.Month, dayOfMonth int, weekNum string, weekday string, rawInfo string) (*DayAudiobook, error) {
	return &DayAudiobook{
		RawInfo: rawInfo,
		Date: date,
		Weekday: weekday,
		DayOfYear: dayOfYear,
		DayOfMonth: dayOfMonth,
		WeekNum: weekNum,
		Month: month,
		Year: year,

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
  klog.InfoS("DayAudiobook", "date", d.Date, "rawInfo", d.RawInfo, "finishBooks", d.FinishBooks, "isFinish", d.IsFinish)
}