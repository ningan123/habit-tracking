package reading

import (
	"time"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)



type Reading struct {
	RawInfo map[string]string // 原始数据
	YearReadingInfo map[int]*YearReading
	MonthReadingInfo map[time.Month]*MonthReading
	MonthOrderReadingInfo []*MonthReading
	WeekReadingInfo map[int]*WeekReading
	WeekOrderReadingInfo []*WeekReading
  DayReadingInfo map[string]*DayReading 
	DayOrderReadingInfo []*DayReading
}


func NewReading(rawInfo map[string]string) *Reading {
	return &Reading{
	  RawInfo: rawInfo,
		YearReadingInfo: make(map[int]*YearReading),
		MonthReadingInfo: make(map[time.Month]*MonthReading),
		MonthOrderReadingInfo: make([]*MonthReading, 12),
	  WeekReadingInfo: make(map[int]*WeekReading),
		WeekOrderReadingInfo: make([]*WeekReading, 53),
    DayReadingInfo: make(map[string]*DayReading),
		DayOrderReadingInfo: make([]*DayReading, 365),
	}
}



func(r *Reading) GenYearAndMonthAndWeekAndDayReadingInfo() error {
	klog.InfoS("GenYearAndMonthAndWeekAndDayReadingInfo")
  for date, info := range r.RawInfo {
		year, dayOfYear, month, dayOfMonth, weekNum, weekday, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.InfoS("date detail", "date", date, "year", year, "dayOfYear", dayOfYear, "month", month, "dayOfMonth", dayOfMonth, "weekNum", weekNum, "weekday", weekday)

		if r.MonthReadingInfo[month] == nil {
		  monthRawInfo := make(map[int]*DayReading)
			r.MonthReadingInfo[month], err = NewMonthReading(month, monthRawInfo)
			if err != nil {
			  return err
			}
		}

		if r.WeekReadingInfo[weekNum] == nil {
			weekRawInfo := make(map[string]*DayReading)
			r.WeekReadingInfo[weekNum], err = NewWeekReading(weekNum, weekRawInfo)
			if err != nil {
			  return err
			}
		}

		if r.YearReadingInfo[year] == nil {
			yearRawInfo := make(map[string]*DayReading)
			r.YearReadingInfo[year], err = NewYearReading(year, yearRawInfo)
			if err != nil {
			  return err
			}
		}


		dReading, err := NewDayReading(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
    if err != nil {
      return err
    }
		
		dReading2, err := NewDayReading(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
    if err != nil {
      return err
    }

		dReading3, err := NewDayReading(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
    if err != nil {
      return err
    }

		dReading4, err := NewDayReading(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
		if err != nil {
			return err
		}
		r.DayReadingInfo[date] = dReading
		r.WeekReadingInfo[weekNum].WeekRawInfo[weekday.String()] = dReading2
		r.MonthReadingInfo[month].MonthRawInfo[dayOfMonth] = dReading3
		r.YearReadingInfo[year].YearRawInfo[date] = dReading4
	   
  }  
  return nil
}


func(r *Reading) ComputeDayReadingTime() error {
	klog.InfoS("ComputDayReadingTime")
	for _, dReading := range r.DayReadingInfo {
		err := dReading.ComputeReadingTime()
		if err != nil {
		  return err
		}
	}
  return nil
}


func(r *Reading) ComputeWeekReadingTime() error {
	klog.InfoS("ComputWeekReadingTime")
	for _, wReading := range r.WeekReadingInfo {
		err := wReading.ComputeReadingTime()
		if err != nil {
		  return err
		}
	}
  return nil
}

func(r *Reading) ComputeMonthReadingTime() error {
	klog.InfoS("ComputMonthReadingTime")
	for _, mReading := range r.MonthReadingInfo {
		err := mReading.ComputeReadingTime()
		if err != nil {
		  return err
		}
	}
  return nil
}


func(r *Reading) ComputeYearReadingTime() error {
	klog.InfoS("ComputYearReadingTime")
	for _, mReading := range r.YearReadingInfo {
		err := mReading.ComputeReadingTime()
		if err != nil {
		  return err
		}
	}
  return nil
}

func (r *Reading) ConverDayReadingInfoToDayOrderReadingInfo() error {
	for _, dReading := range r.DayReadingInfo {
	  r.DayOrderReadingInfo[dReading.DayOfYear-1] = dReading
	}
	return nil
}


func (r *Reading) ConvertWeekReadingInfoToWeekOrderReadingInfo() error {
	for weekNum, mReading := range r.WeekReadingInfo {
		// klog.InfoS("ConvertWeekReadingTimeToWeekOrderReadingTime", "weekNum", weekNum, "mReading", mReading)
	  r.WeekOrderReadingInfo[weekNum-1] = mReading		
	}
  return nil
}


func (r *Reading) ConvertMonthReadingInfoToMonthOrderReadingInfo() error {
	monthsMap := map[string]int{  
		"January": 1,  
		"February": 2,  
		"March": 3,  
		"April": 4,  
		"May": 5,  
		"June": 6,  
		"July": 7,  
		"August": 8,  
		"September": 9,  
		"October": 10,  
		"November": 11,  
		"December": 12,  
	}  

	for monthNum, mReading := range r.MonthReadingInfo {
		monStr := monthNum.String()
		// klog.InfoS("ConvertMonthReadingTimeToMonthOrderReadingTime", "monthNum", monthNum, "mReading", mReading)
	  r.MonthOrderReadingInfo[monthsMap[monStr]-1] = mReading		
	}
  return nil
}


func (r *Reading) CheckFinish() error {
	klog.InfoS("CheckFinish")
	err := r.CheckDayFinish()
	if err != nil {
	  return err
	}

  return nil
}

func (r *Reading) CheckDayFinish() error {
	klog.InfoS("CheckDayFinish")
	for _, dReading := range r.DayReadingInfo {
		err := dReading.CheckFinish()
		if err != nil {
		  return err
		}
	}
  return nil
}



func(r *Reading) Print() {
	klog.InfoS("**************** Begin to print reading statistic data ****************")
	// ri
	for _, dReading := range r.DayReadingInfo {
		dReading.Print()
	}

	// 星期
	for _, wReading := range r.WeekReadingInfo {
		wReading.Print()
	}

	// 月
	for _, mReading := range r.MonthReadingInfo {
		mReading.Print()
	}

	for _, yReading := range r.YearReadingInfo {
		yReading.Print()
	}
}




