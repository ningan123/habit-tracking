package reading

import (
	"fmt"
	"sort"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

var (
	TargetDayReadingTime string = "20min"
)

type Reading struct {
	RawInfo map[string]string // 原始数据
	YearReadingInfo map[int]*YearReading
	MonthReadingInfo map[string]*MonthReading
	MonthOrderReadingInfo []*MonthReading
	WeekReadingInfo map[string]*WeekReading
	WeekOrderReadingInfo []*WeekReading
  DayReadingInfo map[string]*DayReading 
	DayOrderReadingInfo []*DayReading
}


func NewReading(rawInfo map[string]string) *Reading {
	return &Reading{
	  RawInfo: rawInfo,
		YearReadingInfo: make(map[int]*YearReading),
		MonthReadingInfo: make(map[string]*MonthReading),
	  WeekReadingInfo: make(map[string]*WeekReading),
    DayReadingInfo: make(map[string]*DayReading),
	}
}



func(r *Reading) GenYearAndMonthAndWeekAndDayReadingInfo() error {
	klog.InfoS("GenYearAndMonthAndWeekAndDayReadingInfo")
  for date, info := range r.RawInfo {
		year, month, weekyear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.InfoS("date detail", "date", date, "year", year, "month", month, "weekyear", weekyear, "week", week, "weekday", weekday, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)
		
		monthNum := fmt.Sprintf("%d-%02d", year, month)
		if r.MonthReadingInfo[monthNum] == nil {
		  monthRawInfo := make(map[int]*DayReading)
			r.MonthReadingInfo[monthNum], err = NewMonthReading(monthNum, monthRawInfo)
			if err != nil {
			  return err
			}
		}

		weekNum := fmt.Sprintf("%d-%02d", weekyear, week) 
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
		r.MonthReadingInfo[monthNum].MonthRawInfo[dayOfMonth] = dReading3
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
	// 提取key并排序
	keys := make([]string, 0, len(r.DayReadingInfo))
	for k := range r.DayReadingInfo {
	  keys = append(keys, k)
	}
	sort.Sort(hDate.ByDate(keys))

	// 按照排序后的键顺序提取值到切片 
	r.DayOrderReadingInfo = make([]*DayReading, len(keys))
	for i, k := range keys {
	  r.DayOrderReadingInfo[i] = r.DayReadingInfo[k]
	}
	return nil
}


func (r *Reading) ConvertWeekReadingInfoToWeekOrderReadingInfo() error {
	// 提取key并排序
	keys := make([]string, 0, len(r.WeekReadingInfo))
	for k := range r.WeekReadingInfo {
	  keys = append(keys, k)
	}
	sort.Sort(hDate.ByYearWeek(keys))

	// 按照排序后的键顺序提取值到切片 
	r.WeekOrderReadingInfo = make([]*WeekReading, len(keys))
	for i, k := range keys {
	  r.WeekOrderReadingInfo[i] = r.WeekReadingInfo[k]
	}

  return nil
}


func (r *Reading) ConvertMonthReadingInfoToMonthOrderReadingInfo() error {
	// 提取key并排序
	keys := make([]string, 0, len(r.MonthReadingInfo))
	for k := range r.MonthReadingInfo {
		keys = append(keys, string(k))
	}
	sort.Sort(hDate.ByYearMonth(keys))

	// 按照排序后的键顺序提取值到切片 
	r.MonthOrderReadingInfo = make([]*MonthReading, len(keys))
	for i, k := range keys {
		r.MonthOrderReadingInfo[i] = r.MonthReadingInfo[k]
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




