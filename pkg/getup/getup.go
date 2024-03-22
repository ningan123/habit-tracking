package getup

import (
	"fmt"
	"sort"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

var (
	TargetDayGetupTime string = "7:30"
	TargetWeekFinishDays int = 5
	TargetMonthFinishDays int = 21
	TargetYearFinishDays int = TargetMonthFinishDays*12
)

type Getup struct {
	RawInfo map[string]string // 原始数据
	DayGetupInfo map[string]*DayGetup
	DayOrderGetupInfo []*DayGetup
	WeekGetupInfo map[string]*WeekGetup
	WeekOrderGetupInfo []*WeekGetup
	MonthGetupInfo map[string]*MonthGetup
	MonthOrderGetupInfo []*MonthGetup
	YearGetupInfo map[string]*YearGetup
	YearOrderGetupInfo []*YearGetup
}


func NewGetup(rawInfo map[string]string) *Getup {
	return &Getup{
		RawInfo: rawInfo,
		DayGetupInfo: make(map[string]*DayGetup),
		WeekGetupInfo: make(map[string]*WeekGetup),
		MonthGetupInfo: make(map[string]*MonthGetup),
		YearGetupInfo: make(map[string]*YearGetup),
	}
}

func(g *Getup) GenGetupInfo() error {
	klog.Info("GenGetupInfo")
	err := g.GenDayGetupInfo()
	if err != nil {
		return err
	}

	err = g.GenWeekGetupInfo()
	if err != nil {
		return err
	}

	err = g.GenMonthGetupInfo()
	if err != nil {
		return err
	}

	err = g.GenYearGetupInfo()
	if err != nil {
		return err
	}


	return nil
}

func(g *Getup) GenDayGetupInfo() error {
	klog.InfoS("GenDayGetupInfo")
  for date, info := range g.RawInfo {
		year, month, weekyear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.InfoS("date detail", "date", date, "year", year, "month", month, "weekyear", weekyear, "week", week, "weekday", weekday, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)
		
		weekNum := fmt.Sprintf("%d-%02d", weekyear, week)
		dGetup, err := NewDayGetup(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
    if err != nil {
      return err
    }
		
		g.DayGetupInfo[date] = dGetup
  }  
  return nil
}




func(r *Getup) GenWeekGetupInfo() error {
	klog.InfoS("GenWeekGetupInfo")
  for date, info := range r.RawInfo {
		year, month, weekyear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.InfoS("date detail", "date", date, "year", year, "month", month, "weekyear", weekyear, "week", week, "weekday", weekday, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		weekNum := fmt.Sprintf("%d-%02d", weekyear, week) 
		if r.WeekGetupInfo[weekNum] == nil {
			weekRawInfo := make(map[string]*DayGetup)
			r.WeekGetupInfo[weekNum], err = NewWeekGetup(weekNum, weekRawInfo)
			if err != nil {
			  return err
			}
		}
		
		dGetup, err := NewDayGetup(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
    if err != nil {
      return err
		}

		r.WeekGetupInfo[weekNum].RawInfo[weekday.String()] = dGetup
	   
  }  
  return nil
}



func(g *Getup) GenMonthGetupInfo() error {
	klog.InfoS("GenMonthGetupInfo")
  for date, info := range g.RawInfo {
		year, month, weekyear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.InfoS("date detail", "date", date, "year", year, "month", month, "weekyear", weekyear, "week", week, "weekday", weekday, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)
		
		monthNum := fmt.Sprintf("%d-%02d", year, month)
		if g.MonthGetupInfo[monthNum] == nil {
		  monthRawInfo := make(map[int]*DayGetup)
			g.MonthGetupInfo[monthNum], err = NewMonthGetup(monthNum, monthRawInfo)
			if err != nil {
			  return err
			}
		}
		weekNum := fmt.Sprintf("%d-%02d", weekyear, week) 

		dGetup, err := NewDayGetup(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
    if err != nil {
      return err
    }

		g.MonthGetupInfo[monthNum].RawInfo[dayOfMonth] = dGetup	   
  }  
  return nil
}




func(r *Getup) GenYearGetupInfo() error {
	klog.InfoS("GenYearGetupInfo")
  for date, info := range r.RawInfo {
		year, month, weekyear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.InfoS("date detail", "date", date, "year", year, "month", month, "weekyear", weekyear, "week", week, "weekday", weekday, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		yearNum := fmt.Sprintf("%d", year)
		if r.YearGetupInfo[yearNum] == nil {
			yearRawInfo := make(map[string]*DayGetup)
			r.YearGetupInfo[yearNum], err = NewYearGetup(yearNum, yearRawInfo)
			if err != nil {
			  return err
			}
		}
		weekNum := fmt.Sprintf("%d-%02d", weekyear, week)

		dGetup, err := NewDayGetup(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
		if err != nil {
			return err
		}

		r.YearGetupInfo[yearNum].RawInfo[date] = dGetup
	   
  }  
  return nil
}



func (g *Getup) CheckFinish() error {
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


func (g *Getup) CheckDayFinish() error {
	klog.InfoS("CheckDayFinish")
	for _, dGetup := range g.DayGetupInfo {
		err := dGetup.CheckFinish()
		if err != nil {
		  return err
		}
	}
  return nil
}


func (g *Getup) CheckWeekFinish() error {
  klog.InfoS("CheckWeekFinish")
	for _, wGetup := range g.WeekGetupInfo {
	  err := wGetup.CheckFinish()
		if err != nil {
		  return err
		}
	}
	return nil
}


func (g *Getup) CheckMonthFinish() error {
  klog.InfoS("CheckMonthFinish")
	for _, mGetup := range g.MonthGetupInfo {
	  err := mGetup.CheckFinish()
		if err != nil {
		  return err
		}
	}
	return nil
}


func (g *Getup) CheckYearFinish() error {
  klog.InfoS("CheckYearFinish")
	for _, gGetup := range g.YearGetupInfo {
	  err := gGetup.CheckFinish()
		if err != nil {
		  return err
		}
	}
	return nil
}


// ==================================================
// ==================================================

func (g *Getup) ConvertGetupInfoToOrderGetupInfo() error {
	klog.InfoS("ConverGetupInfoToOrderGetupInfo")

	err := g.ConvertDayGetupInfoToDayOrderGetupInfo() 
	if err != nil {
		return err
	}

	err = g.ConvertWeekGetupInfoToWeekOrderGetupInfo()
	if err != nil {
		return err
	}

	err = g.ConvertMonthGetupInfoToMonthOrderGetupInfo()
	if err != nil {
		return err
	}

	err = g.ConvertYearGetupInfoToYearOrderGetupInfo()
	if err != nil {
		return err
	}

	return nil
}


func (g *Getup) ConvertDayGetupInfoToDayOrderGetupInfo() error {
	klog.InfoS("ConvertDayGetupInfoToDayOrderGetupInfo")
	// 提取key并排序
	keys := make([]string, 0, len(g.DayGetupInfo))
	for k := range g.DayGetupInfo {
	  keys = append(keys, k)
	}
	sort.Sort(hDate.ByDate(keys))

	// 按照排序后的键顺序提取值到切片 
	g.DayOrderGetupInfo = make([]*DayGetup, len(keys))
	for i, k := range keys {
	  g.DayOrderGetupInfo[i] = g.DayGetupInfo[k]
	}
	return nil
}

func (r *Getup) ConvertWeekGetupInfoToWeekOrderGetupInfo() error {
	klog.InfoS("ConvertWeekGetupInfoToWeekOrderGetupInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.WeekGetupInfo))
	for k := range r.WeekGetupInfo {
	  keys = append(keys, k)
	}
	sort.Sort(hDate.ByYearWeek(keys))

	// 按照排序后的键顺序提取值到切片 
	r.WeekOrderGetupInfo = make([]*WeekGetup, len(keys))
	for i, k := range keys {
	  r.WeekOrderGetupInfo[i] = r.WeekGetupInfo[k]
	}

	// for index, weekGetupInfo := range r.WeekOrderGetupInfo {
	//   klog.InfoS("ConvertWeekGetupInfoToWeekOrderGetupInfo", "index", index, "weekGetupInfo", weekGetupInfo)
	// }

  return nil
}

func (r *Getup) ConvertMonthGetupInfoToMonthOrderGetupInfo() error {
  klog.InfoS("ConvertMonthGetupInfoToMonthOrderGetupInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.MonthGetupInfo))
	for k := range r.MonthGetupInfo {
	  keys = append(keys, k)
	}
	sort.Sort(hDate.ByYearMonth(keys))

	// 按照排序后的键顺序提取值到切片
	r.MonthOrderGetupInfo = make([]*MonthGetup, len(keys))
	for i, k := range keys {
	  r.MonthOrderGetupInfo[i] = r.MonthGetupInfo[k]
	}

	return nil
}

func (r *Getup) ConvertYearGetupInfoToYearOrderGetupInfo() error {
  klog.InfoS("ConvertYearGetupInfoToYearOrderGetupInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.YearGetupInfo))
	for k := range r.YearGetupInfo {	  
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 按照排序后的键顺序提取值到切片
	r.YearOrderGetupInfo = make([]*YearGetup, len(keys))
	for i, k := range keys {
	  r.YearOrderGetupInfo[i] = r.YearGetupInfo[k]
	}

	return nil
}