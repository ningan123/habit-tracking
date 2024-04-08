package reading

import (
	"errors"
	"fmt"
	"strings"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type DayReading struct {
	Day *hDate.Day
  RawInfo  string
	ReadingTime string // 这一天总共的阅读时长
	ReadingTimeOfDifferentContent map[string]string
	ReadingTimeOfDifferentContentStr string
	ContentInfoList []ContentInfo 
	IsFinish bool
	TargetReadingTime string
}


type ContentInfo struct {
	Content string 
	ReadingTime string
}


func NewDayReading(date string, weekday string, weekNum string, monthNum string, yearNum string, dayOfMonth int, dayOfYear int, rawInfo string) (*DayReading, error) {
	contentInfoList, err := SplitRawInfo(rawInfo)
	if err != nil {
		return nil, err
	}

	return &DayReading{
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
		ReadingTime: "0min",
		ReadingTimeOfDifferentContent: make(map[string]string),
		ContentInfoList: contentInfoList,
		TargetReadingTime: TargetDayReadingTime,
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

	if rawInfo == "" || rawInfo == "×" {
	  return nil, nil
	}

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
		if _, ok := d.ReadingTimeOfDifferentContent[contentInfo.Content]; !ok {
			d.ReadingTimeOfDifferentContent[contentInfo.Content] = contentInfo.ReadingTime
		} else {
			conSum, err := hDate.FormatDurationSum(d.ReadingTimeOfDifferentContent[contentInfo.Content], contentInfo.ReadingTime)
			if err != nil {
				return err
			}
			d.ReadingTimeOfDifferentContent[contentInfo.Content] = conSum
		}
		
		// 计算DayReadingTime
	  sum, err := hDate.FormatDurationSum(d.ReadingTime, contentInfo.ReadingTime)
		if err != nil {
			return err 
		}
		d.ReadingTime = sum
	}

	for k,v := range d.ReadingTimeOfDifferentContent {
	  d.ReadingTimeOfDifferentContentStr += fmt.Sprintf("%s: %s	", k, v)
	}
	return nil
}


func (d *DayReading) Print() {
	for _, conInfo := range d.ContentInfoList {
		klog.InfoS("day reading info", "date", d.Day.Date,"readingTime", d.ReadingTime,  "content", conInfo.Content, "contentReadingTime", conInfo.ReadingTime)
	}
}

// 只要阅读时长>=target时长，就认为完成
func (d *DayReading) CheckFinish() error {
  res, err :=  hDate.IsActualDurationLongerOrEqualToTargetDuration(d.ReadingTime, d.TargetReadingTime)
	if err != nil {
		return err
	}
	d.IsFinish = res
	return nil
}