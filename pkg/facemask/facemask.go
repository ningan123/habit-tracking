package facemask

import (
	"fmt"
	"sort"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
	hTickType "ningan.com/habit-tracking/pkg/tick_type"
)

var (
	FaceMaskTargetWeekFinishDays int = 2
	FaceMaskTargetMonthFinishDays int = 8
	FaceMaskTargetYearFinishDays int = FaceMaskTargetMonthFinishDays*12
)

type FaceMask struct {
	RawInfo map[string]string // 原始数据
	DayInfo map[string]*hTickType.DayTickType
	DayOrderInfo []*hTickType.DayTickType
	WeekInfo map[string]*hTickType.WeekTickType
	WeekOrderInfo []*hTickType.WeekTickType
	MonthInfo map[string]*hTickType.MonthTickType
	MonthOrderInfo []*hTickType.MonthTickType
	YearInfo map[string]*hTickType.YearTickType
	YearOrderInfo []*hTickType.YearTickType
}


func NewFaceMask(rawInfo map[string]string) *FaceMask {
	return &FaceMask{
		RawInfo: rawInfo,
		DayInfo: make(map[string]*hTickType.DayTickType),
		WeekInfo: make(map[string]*hTickType.WeekTickType),
		MonthInfo: make(map[string]*hTickType.MonthTickType),
		YearInfo: make(map[string]*hTickType.YearTickType),
	}
}

func(g *FaceMask) GenInfo() error {
	klog.Info("GenInfo")
	err := g.GenDayInfo()
	if err != nil {
		return err
	}

	err = g.GenWeekInfo()
	if err != nil {
		return err
	}

	err = g.GenMonthInfo()
	if err != nil {
		return err
	}

	err = g.GenYearInfo()
	if err != nil {
		return err
	}


	return nil
}

func(g *FaceMask) GenDayInfo() error {
	klog.InfoS("GenDayInfo")
  for date, info := range g.RawInfo {
		year, month, weekyear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "month", month, "weekyear", weekyear, "week", week, "weekday", weekday, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)
		
		weekNum := fmt.Sprintf("%d-%02d", weekyear, week)
		dTickType, err := hTickType.NewDayTickType(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
    if err != nil {
      return err
    }
		
		g.DayInfo[date] = dTickType
  }  
  return nil
}




func(r *FaceMask) GenWeekInfo() error {
	klog.InfoS("GenWeekInfo")
  for date, info := range r.RawInfo {
		year, month, weekyear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "month", month, "weekyear", weekyear, "week", week, "weekday", weekday, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		weekNum := fmt.Sprintf("%d-%02d", weekyear, week) 
		if r.WeekInfo[weekNum] == nil {
			weekRawInfo := make(map[string]*hTickType.DayTickType)
			r.WeekInfo[weekNum], err = hTickType.NewWeekTickType(weekNum, weekRawInfo, FaceMaskTargetWeekFinishDays)
			if err != nil {
			  return err
			}
		}
		
		dTickType, err := hTickType.NewDayTickType(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
    if err != nil {
      return err
		}

		r.WeekInfo[weekNum].RawInfo[weekday] = dTickType
	   
  }  
  return nil
}



func(g *FaceMask) GenMonthInfo() error {
	klog.InfoS("GenMonthInfo")
  for date, info := range g.RawInfo {
		year, month, weekyear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "month", month, "weekyear", weekyear, "week", week, "weekday", weekday, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)
		
		monthNum := fmt.Sprintf("%d-%02d", year, month)
		if g.MonthInfo[monthNum] == nil {
		  monthRawInfo := make(map[int]*hTickType.DayTickType)
			g.MonthInfo[monthNum], err = hTickType.NewMonthTickType(monthNum, monthRawInfo, FaceMaskTargetMonthFinishDays)
			if err != nil {
			  return err
			}
		}
		weekNum := fmt.Sprintf("%d-%02d", weekyear, week) 

		dTickType, err := hTickType.NewDayTickType(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
    if err != nil {
      return err
    }

		g.MonthInfo[monthNum].RawInfo[dayOfMonth] = dTickType	   
  }  
  return nil
}




func(r *FaceMask) GenYearInfo() error {
	klog.InfoS("GenYearInfo")
  for date, info := range r.RawInfo {
		year, month, weekyear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "month", month, "weekyear", weekyear, "week", week, "weekday", weekday, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		yearNum := fmt.Sprintf("%d", year)
		if r.YearInfo[yearNum] == nil {
			yearRawInfo := make(map[string]*hTickType.DayTickType)
			r.YearInfo[yearNum], err = hTickType.NewYearTickType(yearNum, yearRawInfo, FaceMaskTargetYearFinishDays)
			if err != nil {
			  return err
			}
		}
		weekNum := fmt.Sprintf("%d-%02d", weekyear, week)

		dTickType, err := hTickType.NewDayTickType(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
		if err != nil {
			return err
		}

		r.YearInfo[yearNum].RawInfo[date] = dTickType
	   
  }  
  return nil
}



func (g *FaceMask) CheckFinish() error {
	klog.InfoS("CheckFinish")
	err := g.CheckDayFinish()
	if err != nil {
	  return err
	}

	err = g.CheckWeekFinish()
	if err != nil {
	  return err
	}

	err = g.CheckMonthFinish()
	if err != nil {
	  return err
	}

	err = g.CheckYearFinish()
	if err != nil {
	  return err
	}
	return nil
}


func (g *FaceMask) CheckDayFinish() error {
	klog.InfoS("CheckDayFinish")
	for _, dTickType := range g.DayInfo {
		err := dTickType.CheckFinish()
		if err != nil {
		  return err
		}
	}
  return nil
}


func (g *FaceMask) CheckWeekFinish() error {
  klog.InfoS("CheckWeekFinish")
	for _, wTickType := range g.WeekInfo {
	  err := wTickType.CheckFinish()
		if err != nil {
		  return err
		}
	}
	return nil
}


func (g *FaceMask) CheckMonthFinish() error {
  klog.InfoS("CheckMonthFinish")
	for _, mTickType := range g.MonthInfo {
	  err := mTickType.CheckFinish()
		if err != nil {
		  return err
		}
	}
	return nil
}


func (g *FaceMask) CheckYearFinish() error {
  klog.InfoS("CheckYearFinish")
	for _, gTickType := range g.YearInfo {
	  err := gTickType.CheckFinish()
		if err != nil {
		  return err
		}
	}
	return nil
}


// ==================================================
// ==================================================

func (g *FaceMask) ConvertInfoToOrderInfo() error {
	klog.InfoS("ConvertInfoToOrderInfo")

	err := g.ConvertDayInfoToDayOrderInfo() 
	if err != nil {
		return err
	}

	err = g.ConvertWeekInfoToWeekOrderInfo()
	if err != nil {
		return err
	}

	err = g.ConvertMonthInfoToMonthOrderInfo()
	if err != nil {
		return err
	}

	err = g.ConvertYearInfoToYearOrderInfo()
	if err != nil {
		return err
	}

	return nil
}


func (g *FaceMask) ConvertDayInfoToDayOrderInfo() error {
	klog.InfoS("ConvertDayInfoToDayOrderInfo")
	// 提取key并排序
	keys := make([]string, 0, len(g.DayInfo))
	for k := range g.DayInfo {
	  keys = append(keys, k)
	}
	sort.Sort(hDate.ByDate(keys))

	// 按照排序后的键顺序提取值到切片 
	g.DayOrderInfo = make([]*hTickType.DayTickType, len(keys))
	for i, k := range keys {
	  g.DayOrderInfo[i] = g.DayInfo[k]
	}
	return nil
}

func (r *FaceMask) ConvertWeekInfoToWeekOrderInfo() error {
	klog.InfoS("ConvertWeekInfoToWeekOrderInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.WeekInfo))
	for k := range r.WeekInfo {
	  keys = append(keys, k)
	}
	sort.Sort(hDate.ByYearWeek(keys))

	// 按照排序后的键顺序提取值到切片 
	r.WeekOrderInfo = make([]*hTickType.WeekTickType, len(keys))
	for i, k := range keys {
	  r.WeekOrderInfo[i] = r.WeekInfo[k]
	}

	// for index, weekTickTypeInfo := range r.WeekOrderTickTypeInfo {
	//   klog.InfoS("ConvertWeekInfoToWeekOrderTickTypeInfo", "index", index, "weekTickTypeInfo", weekTickTypeInfo)
	// }

  return nil
}

func (r *FaceMask) ConvertMonthInfoToMonthOrderInfo() error {
  klog.InfoS("ConvertMonthInfoToMonthOrderTickTypeInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.MonthInfo))
	for k := range r.MonthInfo {
	  keys = append(keys, k)
	}
	sort.Sort(hDate.ByYearMonth(keys))

	// 按照排序后的键顺序提取值到切片
	r.MonthOrderInfo = make([]*hTickType.MonthTickType, len(keys))
	for i, k := range keys {
	  r.MonthOrderInfo[i] = r.MonthInfo[k]
	}

	return nil
}

func (r *FaceMask) ConvertYearInfoToYearOrderInfo() error {
  klog.InfoS("ConvertYearInfoToYearOrderInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.YearInfo))
	for k := range r.YearInfo {	  
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 按照排序后的键顺序提取值到切片
	r.YearOrderInfo = make([]*hTickType.YearTickType, len(keys))
	for i, k := range keys {
	  r.YearOrderInfo[i] = r.YearInfo[k]
	}

	return nil
}