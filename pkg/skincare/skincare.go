package facemask

import (
	"sort"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
	hTickType "ningan.com/habit-tracking/pkg/tick_type"
)

var (
	SkinCareTargetWeekFinishDays int = 7
)

type SkinCare struct {
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


func NewSkinCare(rawInfo map[string]string) *SkinCare {
	return &SkinCare{
		RawInfo: rawInfo,
		DayInfo: make(map[string]*hTickType.DayTickType),
		WeekInfo: make(map[string]*hTickType.WeekTickType),
		MonthInfo: make(map[string]*hTickType.MonthTickType),
		YearInfo: make(map[string]*hTickType.YearTickType),
	}
}

func(g *SkinCare) GenInfo() error {
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

func(g *SkinCare) GenDayInfo() error {
	klog.InfoS("GenDayInfo")
  for date, rawInfo := range g.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)
		
		dTickType, err := hTickType.NewDayTickType(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
    if err != nil {
      return err
    }
		
		g.DayInfo[date] = dTickType
  }  
  return nil
}




func(r *SkinCare) GenWeekInfo() error {
	klog.InfoS("GenWeekInfo")
  for date, rawInfo := range r.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		if r.WeekInfo[weekNum] == nil {
			weekRawInfo := make(map[string]*hTickType.DayTickType)
			r.WeekInfo[weekNum], err = hTickType.NewWeekTickType(weekNum, weekRawInfo, SkinCareTargetWeekFinishDays)
			if err != nil {
			  return err
			}
		}
		
		dTickType, err := hTickType.NewDayTickType(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
    if err != nil {
      return err
		}

		r.WeekInfo[weekNum].RawInfo[weekday] = dTickType
	   
  }  
  return nil
}



func(g *SkinCare) GenMonthInfo() error {
	klog.InfoS("GenMontInfo")
  for date, rawInfo := range g.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)
		
		if g.MonthInfo[monthNum] == nil {
		  monthRawInfo := make(map[int]*hTickType.DayTickType)
			g.MonthInfo[monthNum], err = hTickType.NewMonthTickType(monthNum, daysInMonth, monthRawInfo, daysInMonth-2)
			if err != nil {
			  return err
			}
		}
		
		dTickType, err := hTickType.NewDayTickType(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
    if err != nil {
      return err
    }

		g.MonthInfo[monthNum].RawInfo[dayOfMonth] = dTickType	   
  }  
  return nil
}




func(r *SkinCare) GenYearInfo() error {
	klog.InfoS("GenYearInfo")
  for date, rawInfo := range r.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		if r.YearInfo[yearNum] == nil {
			yearRawInfo := make(map[string]*hTickType.DayTickType)
			r.YearInfo[yearNum], err = hTickType.NewYearTickType(yearNum, daysInYear, yearRawInfo, daysInYear-20)
			if err != nil {
			  return err
			}
		}
		
		dTickType, err := hTickType.NewDayTickType(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
		if err != nil {
			return err
		}

		r.YearInfo[yearNum].RawInfo[date] = dTickType
	   
  }  
  return nil
}



func (g *SkinCare) CheckFinish() error {
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


func (g *SkinCare) CheckDayFinish() error {
	klog.InfoS("CheckDayFinish")
	for _, dTickType := range g.DayInfo {
		err := dTickType.CheckFinish()
		if err != nil {
		  return err
		}
	}
  return nil
}


func (g *SkinCare) CheckWeekFinish() error {
  klog.InfoS("CheckWeekFinish")
	for _, wTickType := range g.WeekInfo {
	  err := wTickType.CheckFinish()
		if err != nil {
		  return err
		}
	}
	return nil
}


func (g *SkinCare) CheckMonthFinish() error {
  klog.InfoS("CheckMonthFinish")
	for _, mTickType := range g.MonthInfo {
	  err := mTickType.CheckFinish()
		if err != nil {
		  return err
		}
	}
	return nil
}


func (g *SkinCare) CheckYearFinish() error {
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

func (g *SkinCare) ConvertInfoToOrderInfo() error {
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


func (g *SkinCare) ConvertDayInfoToDayOrderInfo() error {
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

func (r *SkinCare) ConvertWeekInfoToWeekOrderInfo() error {
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

func (r *SkinCare) ConvertMonthInfoToMonthOrderInfo() error {
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

func (r *SkinCare) ConvertYearInfoToYearOrderInfo() error {
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