package piano

import (
	"fmt"
	"sort"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

var (
	TargetDayPianoTime string = "20min"
)

type Piano struct {
	RawInfo map[string]string // 原始数据
	YearPianoInfo map[string]*YearPiano
	YearOrderPianoInfo []*YearPiano
	MonthPianoInfo map[string]*MonthPiano
	MonthOrderPianoInfo []*MonthPiano
	WeekPianoInfo map[string]*WeekPiano
	WeekOrderPianoInfo []*WeekPiano
  DayPianoInfo map[string]*DayPiano 
	DayOrderPianoInfo []*DayPiano
}


func NewPiano(rawInfo map[string]string) *Piano {
	return &Piano{
	  RawInfo: rawInfo,
		YearPianoInfo: make(map[string]*YearPiano),
		MonthPianoInfo: make(map[string]*MonthPiano),
	  WeekPianoInfo: make(map[string]*WeekPiano),
    DayPianoInfo: make(map[string]*DayPiano),
	}
}


// ==============================================
// ==============================================

func(r *Piano) GenPianoInfo() error {
	klog.Info("Generate Piano Info")
	err := r.GenDayPianoInfo()
	if err != nil {
		return err
	}

	err = r.GenWeekPianoInfo()
	if err != nil {
		return err
	}

	err = r.GenMonthPianoInfo()
	if err != nil {
		return err
	}

	err = r.GenYearPianoInfo()
	if err != nil {
		return err
	}

	return nil
}


func(r *Piano) GenDayPianoInfo() error {
	klog.InfoS("GenDayPianoInfo")
  for date, info := range r.RawInfo {
		year, month, weekyear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.InfoS("date detail", "date", date, "year", year, "month", month, "weekyear", weekyear, "week", week, "weekday", weekday, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)
		
		weekNum := fmt.Sprintf("%d-%02d", weekyear, week)
		dPiano, err := NewDayPiano(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
    if err != nil {
      return err
    }
		
		r.DayPianoInfo[date] = dPiano
  }  
  return nil
}


func(r *Piano) GenWeekPianoInfo() error {
	klog.InfoS("GenWeekPianoInfo")
  for date, info := range r.RawInfo {
		year, month, weekyear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.InfoS("date detail", "date", date, "year", year, "month", month, "weekyear", weekyear, "week", week, "weekday", weekday, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		weekNum := fmt.Sprintf("%d-%02d", weekyear, week) 
		if r.WeekPianoInfo[weekNum] == nil {
			weekRawInfo := make(map[string]*DayPiano)
			r.WeekPianoInfo[weekNum], err = NewWeekPiano(weekNum, weekRawInfo)
			if err != nil {
			  return err
			}
		}
		
		dPiano, err := NewDayPiano(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
    if err != nil {
      return err
		}

		r.WeekPianoInfo[weekNum].RawInfo[weekday] = dPiano
	   
  }  
  return nil
}


func(r *Piano) GenMonthPianoInfo() error {
	klog.InfoS("GenMonthPianoInfo")
  for date, info := range r.RawInfo {
		year, month, weekyear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.InfoS("date detail", "date", date, "year", year, "month", month, "weekyear", weekyear, "week", week, "weekday", weekday, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)
		
		monthNum := fmt.Sprintf("%d-%02d", year, month)
		if r.MonthPianoInfo[monthNum] == nil {
		  monthRawInfo := make(map[int]*DayPiano)
			r.MonthPianoInfo[monthNum], err = NewMonthPiano(monthNum, monthRawInfo, daysInMonth)
			if err != nil {
			  return err
			}
		}
		weekNum := fmt.Sprintf("%d-%02d", weekyear, week) 

		dPiano, err := NewDayPiano(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
    if err != nil {
      return err
    }

		r.MonthPianoInfo[monthNum].RawInfo[dayOfMonth] = dPiano	   
  }  
  return nil
}


func(r *Piano) GenYearPianoInfo() error {
	klog.InfoS("GenYearPianoInfo")
  for date, info := range r.RawInfo {
		year, month, weekyear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.InfoS("date detail", "date", date, "year", year, "month", month, "weekyear", weekyear, "week", week, "weekday", weekday, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		yearNum := fmt.Sprintf("%d", year)
		if r.YearPianoInfo[yearNum] == nil {
			yearRawInfo := make(map[string]*DayPiano)
			r.YearPianoInfo[yearNum], err = NewYearPiano(yearNum, yearRawInfo, daysInYear)
			if err != nil {
			  return err
			}
		}
		weekNum := fmt.Sprintf("%d-%02d", weekyear, week)

		dPiano, err := NewDayPiano(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
		if err != nil {
			return err
		}

		r.YearPianoInfo[yearNum].YearRawInfo[date] = dPiano
	   
  }  
  return nil
}

// ==============================================
// ==============================================

func(r *Piano) ComputePianoTime() error {
	klog.InfoS("ComputePianoTime")

	err := r.ComputeDayPianoTime()
	if err != nil {
		return err
	}

	err = r.ComputeWeekPianoTime()
	if err != nil {
		return err
	}

	err = r.ComputeMonthPianoTime()
	if err != nil {
		return err
	}

	err = r.ComputeYearPianoTime()
	if err != nil {
		return err
	}

	return nil
}



func(r *Piano) ComputeDayPianoTime() error {
	klog.InfoS("ComputDayPianoTime")
	for _, dPiano := range r.DayPianoInfo {
		err := dPiano.ComputePianoTime()
		if err != nil {
		  return err
		}
	}
  return nil
}


func(r *Piano) ComputeWeekPianoTime() error {
	klog.InfoS("ComputWeekPianoTime")
	for _, wPiano := range r.WeekPianoInfo {
		err := wPiano.ComputePianoTime()
		if err != nil {
		  return err
		}
	}
  return nil
}

func(r *Piano) ComputeMonthPianoTime() error {
	klog.InfoS("ComputMonthPianoTime")
	for _, mPiano := range r.MonthPianoInfo {
		err := mPiano.ComputePianoTime()
		if err != nil {
		  return err
		}
	}
  return nil
}


func(r *Piano) ComputeYearPianoTime() error {
	klog.InfoS("ComputYearPianoTime")
	for _, mPiano := range r.YearPianoInfo {
		err := mPiano.ComputePianoTime()
		if err != nil {
		  return err
		}
	}
  return nil
}


// ==================================================
// ==================================================

func (r *Piano) ConvertPianoInfoToOrderPianoInfo() error {
	klog.InfoS("ConverPianoInfoToOrderPianoInfo")

	err := r.ConvertDayPianoInfoToDayOrderPianoInfo() 
	if err != nil {
		return err
	}

	err = r.ConvertWeekPianoInfoToWeekOrderPianoInfo()
	if err != nil {
		return err
	}

	err = r.ConvertMonthPianoInfoToMonthOrderPianoInfo()
	if err != nil {
		return err
	}

	err = r.ConvertYearPianoInfoToYearOrderPianoInfo()
	if err != nil {
		return err
	}

	return nil
}


func (r *Piano) ConvertDayPianoInfoToDayOrderPianoInfo() error {
	klog.InfoS("ConvertDayPianoInfoToDayOrderPianoInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.DayPianoInfo))
	for k := range r.DayPianoInfo {
	  keys = append(keys, k)
	}
	sort.Sort(hDate.ByDate(keys))

	// 按照排序后的键顺序提取值到切片 
	r.DayOrderPianoInfo = make([]*DayPiano, len(keys))
	for i, k := range keys {
	  r.DayOrderPianoInfo[i] = r.DayPianoInfo[k]
	}
	return nil
}


func (r *Piano) ConvertWeekPianoInfoToWeekOrderPianoInfo() error {
	klog.InfoS("ConvertWeekPianoInfoToWeekOrderPianoInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.WeekPianoInfo))
	for k := range r.WeekPianoInfo {
	  keys = append(keys, k)
	}
	sort.Sort(hDate.ByYearWeek(keys))

	// 按照排序后的键顺序提取值到切片 
	r.WeekOrderPianoInfo = make([]*WeekPiano, len(keys))
	for i, k := range keys {
	  r.WeekOrderPianoInfo[i] = r.WeekPianoInfo[k]
	}

  return nil
}


func (r *Piano) ConvertMonthPianoInfoToMonthOrderPianoInfo() error {
	klog.InfoS("ConvertMonthPianoInfoToMonthOrderPianoInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.MonthPianoInfo))
	for k := range r.MonthPianoInfo {
		keys = append(keys, string(k))
	}
	sort.Sort(hDate.ByYearMonth(keys))

	// 按照排序后的键顺序提取值到切片 
	r.MonthOrderPianoInfo = make([]*MonthPiano, len(keys))
	for i, k := range keys {
		r.MonthOrderPianoInfo[i] = r.MonthPianoInfo[k]
	}

  return nil
}


func (r *Piano) ConvertYearPianoInfoToYearOrderPianoInfo() error {
	klog.InfoS("ConvertYearPianoInfoToYearOrderPianoInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.YearPianoInfo))
	for k := range r.YearPianoInfo {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 按照排序后的键顺序提取值到切片 
	r.YearOrderPianoInfo = make([]*YearPiano, len(keys))
	for i, k := range keys {
		r.YearOrderPianoInfo[i] = r.YearPianoInfo[k]
	}

  return nil
}

// ==============================================
// ==============================================


func (r *Piano) CheckFinish() error {
	klog.InfoS("CheckFinish")
	err := r.CheckDayFinish()
	if err != nil {
	  return err
	}

	err = r.CheckWeekFinish()
	if err != nil {
	  return err
	}

  return nil
}

func (r *Piano) CheckDayFinish() error {
	klog.InfoS("CheckDayFinish")
	for _, dPiano := range r.DayPianoInfo {
		err := dPiano.CheckFinish()
		if err != nil {
		  return err
		}
	}
  return nil
}

func (r *Piano) CheckWeekFinish() error {
  klog.InfoS("CheckWeekFinish")
	for _, wPiano := range r.WeekPianoInfo {
	  err := wPiano.CheckFinish()
		if err != nil {
		  return err
		}
	}
	return nil
}

// ==============================================
// ==============================================


func(r *Piano) ComputeExtraPianoTime() error {
  klog.InfoS("ComputeExtraPianoTime")
	err := r.ComputeWeekExtraPianoTime()
	if err != nil {
	  return err
	}
	return nil
}


func(r *Piano) ComputeWeekExtraPianoTime() error {
  klog.InfoS("ComputeWeekExtraPianoTime")
	for _, wPiano := range r.WeekPianoInfo {
	  err := wPiano.ComputeExtraPianoTime()
		if err != nil {
		  return err
		}
	}
	return nil
}


// ==============================================
// ==============================================



func(r *Piano) PrintPianoInfo() {
	klog.InfoS("PrintPianoInfo")

	r.PrintDayPianoInfo()
	r.PrintWeekPianoInfo()
	r.PrintMonthPianoInfo()
	r.PrintYearPianoInfo()
}



func(r *Piano) PrintDayPianoInfo() {
	klog.InfoS("PrintDayPianoInfo")
	for _, dPiano := range r.DayPianoInfo {
		dPiano.Print()
	}
}

func(r *Piano) PrintWeekPianoInfo() {
	klog.InfoS("PrintWeekPianoInfo")
	for _, wPiano := range r.WeekPianoInfo {
		wPiano.Print()
	}
}

func(r *Piano) PrintMonthPianoInfo() {
	klog.InfoS("PrintMonthPianoInfo")
	for _, mPiano := range r.MonthPianoInfo {
		mPiano.Print()
	}
}

func(r *Piano) PrintYearPianoInfo() {
	klog.InfoS("PrintYearPianoInfo")
	for _, yPiano := range r.YearPianoInfo {
		yPiano.Print()
	}
}


