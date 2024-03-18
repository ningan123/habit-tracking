package reading

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type DayReading struct {
  DayRawInfo  string
	DayDate string // 具体日期
	DayWeek time.Weekday // 星期几
	DayMonth time.Month // 几月
	DayYear int // 哪一年
	DayOfYear int // 一年中的第几天
	DayOfMonth int // 一个月中的第几天
	DayReadingTime string // 这一天总共的阅读时长
	DayReadingTimeOfDifferentContent map[string]string
	DayReadingTimeOfDifferentContentStr string
	ContentInfoList []ContentInfo 
}


type ContentInfo struct {
	Content string 
	ReadingTime string
}


func NewDayReading(date string, year int, dayOfYear int, month time.Month, dayOfMonth int, weekNum int, weekday time.Weekday, dayRawInfo string) (*DayReading, error) {
	contentInfoList, err := SplitRawInfo(dayRawInfo)
	if err != nil {
		return nil, err
	}

	return &DayReading{
		DayRawInfo: dayRawInfo,
		DayDate: date,
		DayWeek: weekday,
		DayOfYear: dayOfYear,
		DayOfMonth: dayOfMonth,
		DayMonth: month,
		DayYear: year,
		DayReadingTime: "0min",
		DayReadingTimeOfDifferentContent: make(map[string]string),
		ContentInfoList: contentInfoList,
	}, nil
}

func SplitRawInfo(rawInfo string) ([]ContentInfo, error) {
	contentInfoList := make([]ContentInfo, 0)
	// 假设rawInfo的格式为： 内容1,时长1;内容2,时长2;
	// 则需要将这个字符串按照分号进行分割，然后再按照逗号进行分割
	// 最终得到一个二维数组，每个元素都是一个字符串数组，其中第一个元素是内容，第二个元素是时长
	// 然后将这个二维数组转换为ContentInfo结构体数组
	// 例如：[["内容1", "时长1"], ["内容2", "时长2"]]
	// 转换为[]ContentInfo{ContentInfo{Content: "内容1", ReadingTime: "时长1"}, ContentInfo{Content: "内容2", ReadingTime: "时长2"}}
	// 然后返回这个数组

	rawInfoList := strings.Split(rawInfo, ";")
	for _, str := range rawInfoList {
		strList := strings.Split(str, ",")
		if len(strList) != 2 {
			errMsg := fmt.Sprintf("error split raw info: %s", str)
			return nil, errors.New(errMsg)
		}
		contentInfoList = append(contentInfoList, ContentInfo{Content: strList[0], ReadingTime: strList[1]})
		
	}
	return contentInfoList, nil
}


func (d *DayReading) ComputeReadingTime () error {
	// 假设阅读时间等于ContentInfoList中每个ContentInfo的ReadingTime的和
	for _, contentInfo := range d.ContentInfoList {
		// 计算DayReadingTimeOfDifferentContent
		if _, ok := d.DayReadingTimeOfDifferentContent[contentInfo.Content]; !ok {
			d.DayReadingTimeOfDifferentContent[contentInfo.Content] = contentInfo.ReadingTime
		} else {
			conSum, err := hDate.FormatDurationSum(d.DayReadingTimeOfDifferentContent[contentInfo.Content], contentInfo.ReadingTime)
			if err != nil {
				return err
			}
			d.DayReadingTimeOfDifferentContent[contentInfo.Content] = conSum
		}
		
		// 计算DayReadingTime
	  sum, err := hDate.FormatDurationSum(d.DayReadingTime, contentInfo.ReadingTime)
		if err != nil {
			return err 
		}
		d.DayReadingTime = sum
	}

	for k,v := range d.DayReadingTimeOfDifferentContent {
	  d.DayReadingTimeOfDifferentContentStr += fmt.Sprintf("%s: %s	", k, v)
	}
	return nil
}


func (d *DayReading) Print() {
	for _, conInfo := range d.ContentInfoList {
		klog.InfoS("day reading info", "date", d.DayDate,"readingTime", d.DayReadingTime,  "content", conInfo.Content, "contentReadingTime", conInfo.ReadingTime)
	}
}
