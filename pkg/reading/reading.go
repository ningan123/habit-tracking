package reading

import (
	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)



type Reading struct {
	RawInfo map[string]string // 原始数据
	WeekReadingInfo map[int]*WeekReading
  DayReadingInfo map[string]*DayReading 
}


func NewReading(rawInfo map[string]string) *Reading {
	return &Reading{
	  RawInfo: rawInfo,
	  WeekReadingInfo: make(map[int]*WeekReading),
    DayReadingInfo: make(map[string]*DayReading),
	}
}

func(r *Reading) GenMonthAndWeekAndDayReadingInfo() error {
	klog.InfoS("GenMonthAndWeekAndDayReadingInfo")
  for date, info := range r.RawInfo {
		year, month, weekNum, weekday, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.InfoS("date detail", "date", date, "year", year, "month", month, "weekNum", weekNum, "weekday", weekday)

		if r.WeekReadingInfo[weekNum] == nil {
			weekRawInfo := make(map[string]*DayReading)
			r.WeekReadingInfo[weekNum], err = NewWeekReading(weekNum, weekRawInfo)
			if err != nil {
			  return nil
			}
		}
		dReading, err := NewDayReading(date, year, month, weekNum, weekday, info)
    if err != nil {
      return err
    }
		
		r.DayReadingInfo[date] = dReading
		r.WeekReadingInfo[weekNum].WeekRawInfo[weekday.String()] =  dReading
		  
	   
  }  
  return nil
}

func(r *Reading) ComputReadingTime() error {
	klog.InfoS("ComputReadingTime")
	for _, wReading := range r.WeekReadingInfo {
		err := wReading.ComputeReadingTime()
		if err != nil {
		  return err
		}
	}
  return nil
}


func(r *Reading) Print() {
	// 星期
	for _, wReading := range r.WeekReadingInfo {
		wReading.Print()
	}
}




