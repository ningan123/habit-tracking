package sport

import (
	"fmt"
	"strings"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

type DaySport struct {
	Day                             *hDate.Day
	RawInfo                         string
	SportTimes                      int
	SportTimesOfDifferentContent    map[string]int
	SportTimesOfDifferentContentStr string
	ContentInfoList                 []ContentInfo
	IsFinish                        bool
}

type ContentInfo struct {
	Content    string
	SportTimes int // 次数
}

func NewDaySport(date string, weekday string, weekNum string, monthNum string, yearNum string, dayOfMonth int, dayOfYear int, rawInfo string) (*DaySport, error) {
	contentInfoList, err := SplitRawInfo(rawInfo)
	if err != nil {
		klog.ErrorS(err, "date raw info error", "date", date)
		return nil, err
	}

	return &DaySport{
		Day: &hDate.Day{
			Date:       date,
			Weekday:    weekday,
			WeekNum:    weekNum,
			MonthNum:   monthNum,
			YearNum:    yearNum,
			DayOfMonth: dayOfMonth,
			DayOfYear:  dayOfYear,
		},
		RawInfo:                      rawInfo,
		SportTimesOfDifferentContent: make(map[string]int),
		ContentInfoList:              contentInfoList,
	}, nil
}

func SplitRawInfo(rawInfo string) ([]ContentInfo, error) {
	contentInfoList := make([]ContentInfo, 0)

	if rawInfo == "" || rawInfo == "×" {
		return nil, nil
	}

	rawInfoList := strings.Split(rawInfo, "；\n")
	for _, str := range rawInfoList {
		contentInfoList = append(contentInfoList, ContentInfo{Content: str, SportTimes: 1})
	}
	return contentInfoList, nil
}

func (d *DaySport) ComputeSportTimes() error {
	for _, contentInfo := range d.ContentInfoList {
		// 计算DaySportTimeOfDifferentContent
		if _, ok := d.SportTimesOfDifferentContent[contentInfo.Content]; !ok {
			d.SportTimesOfDifferentContent[contentInfo.Content] = contentInfo.SportTimes
		} else {
			d.SportTimesOfDifferentContent[contentInfo.Content] += contentInfo.SportTimes
		}
		d.SportTimes += contentInfo.SportTimes
	}

	for k, v := range d.SportTimesOfDifferentContent {
		d.SportTimesOfDifferentContentStr += fmt.Sprintf("%s: %d<br>", k, v)
	}
	return nil
}

func (d *DaySport) Print() {
	for _, conInfo := range d.ContentInfoList {
		klog.InfoS("day Sport info", "date", d.Day.Date, "content", conInfo.Content, "contentSportTimes", conInfo.SportTimes)
	}
}

func (d *DaySport) CheckFinish() error {
	if d.RawInfo != "" && d.RawInfo != "×" {
		d.IsFinish = true
	}
	return nil
}
