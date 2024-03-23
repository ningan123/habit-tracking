package piano

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type DayPiano struct {
  RawInfo  string
	Date string // 具体日期
	Weekday string // 星期几
	Month time.Month // 几月
	WeekNum string // 几周
	Year int // 哪一年
	DayOfYear int // 一年中的第几天
	DayOfMonth int // 一个月中的第几天
	PianoTime string 
	PianoTimeOfDifferentContent map[string]string
	PianoTimeOfDifferentContentStr string
	ContentInfoList []ContentInfo 
	IsFinish bool
	TargetPianoTime string
}


type ContentInfo struct {
	Content string 
	PianoTime string
}


func NewDayPiano(date string, year int, dayOfYear int, month time.Month, dayOfMonth int, weekNum string, weekday string, rawInfo string) (*DayPiano, error) {
	contentInfoList, err := SplitRawInfo(rawInfo)
	if err != nil {
		return nil, err
	}

	return &DayPiano{
		RawInfo: rawInfo,
		Date: date,
		Weekday: weekday,
		DayOfYear: dayOfYear,
		DayOfMonth: dayOfMonth,
		WeekNum: weekNum,
		Month: month,
		Year: year,
		PianoTime: "0min",
		PianoTimeOfDifferentContent: make(map[string]string),
		ContentInfoList: contentInfoList,
		TargetPianoTime: TargetDayPianoTime,
	}, nil
}

func SplitRawInfo(rawInfo string) ([]ContentInfo, error) {
	contentInfoList := make([]ContentInfo, 0)
	// 假设rawInfo的格式为： 内容1,时长1;内容2,时长2;
	// 则需要将这个字符串按照分号进行分割，然后再按照逗号进行分割
	// 最终得到一个二维数组，每个元素都是一个字符串数组，其中第一个元素是内容，第二个元素是时长
	// 然后将这个二维数组转换为ContentInfo结构体数组
	// 例如：[["内容1", "时长1"], ["内容2", "时长2"]]
	// 转换为[]ContentInfo{ContentInfo{Content: "内容1", PianoTime: "时长1"}, ContentInfo{Content: "内容2", PianoTime: "时长2"}}
	// 然后返回这个数组

	rawInfoList := strings.Split(rawInfo, ";")
	for _, str := range rawInfoList {
		strList := strings.Split(str, ",")
		if len(strList) != 2 {
			errMsg := fmt.Sprintf("error split raw info: %s", str)
			return nil, errors.New(errMsg)
		}
		contentInfoList = append(contentInfoList, ContentInfo{Content: strList[0], PianoTime: strList[1]})
		
	}
	return contentInfoList, nil
}


func (d *DayPiano) ComputePianoTime () error {
	// 假设总时间等于ContentInfoList中每个ContentInfo的PianoTime的和
	for _, contentInfo := range d.ContentInfoList {
		// 计算DayPianoTimeOfDifferentContent
		if _, ok := d.PianoTimeOfDifferentContent[contentInfo.Content]; !ok {
			d.PianoTimeOfDifferentContent[contentInfo.Content] = contentInfo.PianoTime
		} else {
			conSum, err := hDate.FormatDurationSum(d.PianoTimeOfDifferentContent[contentInfo.Content], contentInfo.PianoTime)
			if err != nil {
				return err
			}
			d.PianoTimeOfDifferentContent[contentInfo.Content] = conSum
		}
		
		// 计算DayPianoTime
	  sum, err := hDate.FormatDurationSum(d.PianoTime, contentInfo.PianoTime)
		if err != nil {
			return err 
		}
		d.PianoTime = sum
	}

	for k,v := range d.PianoTimeOfDifferentContent {
	  d.PianoTimeOfDifferentContentStr += fmt.Sprintf("%s: %s	", k, v)
	}
	return nil
}


func (d *DayPiano) Print() {
	for _, conInfo := range d.ContentInfoList {
		klog.InfoS("day piano info", "date", d.Date,"pianoTime", d.PianoTime,  "content", conInfo.Content, "contentPianoTime", conInfo.PianoTime)
	}
}

// 只要时长>=target时长，就认为完成
func (d *DayPiano) CheckFinish() error {
  res, err :=  hDate.IsActualDurationLongerOrEqualToTargetDuration(d.PianoTime, d.TargetPianoTime)
	if err != nil {
		return err
	}
	d.IsFinish = res
	return nil
}