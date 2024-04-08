package sleep

import (
	"sort"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

var (
	TargetDaySleepTime string = "23:30"
	TargetWeekFinishDays int = 5
	TargetMonthFinishDays int = 21
	TargetYearFinishDays int = TargetMonthFinishDays*12
)

type Sleep struct {
	RawInfo map[string]string // 原始数据
	DaySleepInfo map[string]*DaySleep
	DayOrderSleepInfo []*DaySleep
	WeekSleepInfo map[string]*WeekSleep
	WeekOrderSleepInfo []*WeekSleep
	MonthSleepInfo map[string]*MonthSleep
	MonthOrderSleepInfo []*MonthSleep
	YearSleepInfo map[string]*YearSleep
	YearOrderSleepInfo []*YearSleep
}


func NewSleep(rawInfo map[string]string) *Sleep {
	return &Sleep{
		RawInfo: rawInfo,
		DaySleepInfo: make(map[string]*DaySleep),
		WeekSleepInfo: make(map[string]*WeekSleep),
		MonthSleepInfo: make(map[string]*MonthSleep),
		YearSleepInfo: make(map[string]*YearSleep),
	}
}

func(g *Sleep) GenSleepInfo() error {
	klog.Info("GenSleepInfo")
	err := g.GenDaySleepInfo()
	if err != nil {
		return err
	}

	err = g.GenWeekSleepInfo()
	if err != nil {
		return err
	}

	err = g.GenMonthSleepInfo()
	if err != nil {
		return err
	}

	err = g.GenYearSleepInfo()
	if err != nil {
		return err
	}


	return nil
}

func(g *Sleep) GenDaySleepInfo() error {
	klog.InfoS("GenDaySleepInfo")
  for date, rawInfo := range g.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)
		
		dSleep, err := NewDaySleep(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
    if err != nil {
      return err
    }
		
		g.DaySleepInfo[date] = dSleep
  }  
  return nil
}




func(r *Sleep) GenWeekSleepInfo() error {
	klog.InfoS("GenWeekSleepInfo")
  for date, rawInfo := range r.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		if r.WeekSleepInfo[weekNum] == nil {
			weekRawInfo := make(map[string]*DaySleep)
			r.WeekSleepInfo[weekNum], err = NewWeekSleep(weekNum, weekRawInfo)
			if err != nil {
			  return err
			}
		}
		
		dSleep, err := NewDaySleep(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
    if err != nil {
      return err
		}

		r.WeekSleepInfo[weekNum].RawInfo[weekday] = dSleep
	   
  }  
  return nil
}



func(g *Sleep) GenMonthSleepInfo() error {
	klog.InfoS("GenMonthSleepInfo")
  for date, rawInfo := range g.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)
		
		if g.MonthSleepInfo[monthNum] == nil {
		  monthRawInfo := make(map[int]*DaySleep)
			g.MonthSleepInfo[monthNum], err = NewMonthSleep(monthNum, daysInMonth, monthRawInfo)
			if err != nil {
			  return err
			}
		}

		dSleep, err := NewDaySleep(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
    if err != nil {
      return err
    }

		g.MonthSleepInfo[monthNum].RawInfo[dayOfMonth] = dSleep	   
  }  
  return nil
}




func(r *Sleep) GenYearSleepInfo() error {
	klog.InfoS("GenYearSleepInfo")
  for date, rawInfo := range r.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		if r.YearSleepInfo[yearNum] == nil {
			yearRawInfo := make(map[string]*DaySleep)
			r.YearSleepInfo[yearNum], err = NewYearSleep(yearNum, daysInYear, yearRawInfo)
			if err != nil {
			  return err
			}
		}

		dSleep, err := NewDaySleep(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
		if err != nil {
			return err
		}
		r.YearSleepInfo[yearNum].RawInfo[date] = dSleep
	   
  }  
  return nil
}



func (g *Sleep) CheckFinish() error {
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


func (g *Sleep) CheckDayFinish() error {
	klog.InfoS("CheckDayFinish")
	for _, dSleep := range g.DaySleepInfo {
		err := dSleep.CheckFinish()
		if err != nil {
		  return err
		}
	}
  return nil
}


func (g *Sleep) CheckWeekFinish() error {
  klog.InfoS("CheckWeekFinish")
	for _, wSleep := range g.WeekSleepInfo {
	  err := wSleep.CheckFinish()
		if err != nil {
		  return err
		}
	}
	return nil
}


func (g *Sleep) CheckMonthFinish() error {
  klog.InfoS("CheckMonthFinish")
	for _, mSleep := range g.MonthSleepInfo {
	  err := mSleep.CheckFinish()
		if err != nil {
		  return err
		}
	}
	return nil
}


func (g *Sleep) CheckYearFinish() error {
  klog.InfoS("CheckYearFinish")
	for _, gSleep := range g.YearSleepInfo {
	  err := gSleep.CheckFinish()
		if err != nil {
		  return err
		}
	}
	return nil
}


// ==================================================
// ==================================================

func (g *Sleep) ConvertSleepInfoToOrderSleepInfo() error {
	klog.InfoS("ConverSleepInfoToOrderSleepInfo")

	err := g.ConvertDaySleepInfoToDayOrderSleepInfo() 
	if err != nil {
		return err
	}

	err = g.ConvertWeekSleepInfoToWeekOrderSleepInfo()
	if err != nil {
		return err
	}

	err = g.ConvertMonthSleepInfoToMonthOrderSleepInfo()
	if err != nil {
		return err
	}

	err = g.ConvertYearSleepInfoToYearOrderSleepInfo()
	if err != nil {
		return err
	}

	return nil
}


func (g *Sleep) ConvertDaySleepInfoToDayOrderSleepInfo() error {
	klog.InfoS("ConvertDaySleepInfoToDayOrderSleepInfo")
	// 提取key并排序
	keys := make([]string, 0, len(g.DaySleepInfo))
	for k := range g.DaySleepInfo {
	  keys = append(keys, k)
	}
	sort.Sort(hDate.ByDate(keys))

	// 按照排序后的键顺序提取值到切片 
	g.DayOrderSleepInfo = make([]*DaySleep, len(keys))
	for i, k := range keys {
	  g.DayOrderSleepInfo[i] = g.DaySleepInfo[k]
	}
	return nil
}

func (r *Sleep) ConvertWeekSleepInfoToWeekOrderSleepInfo() error {
	klog.InfoS("ConvertWeekSleepInfoToWeekOrderSleepInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.WeekSleepInfo))
	for k := range r.WeekSleepInfo {
	  keys = append(keys, k)
	}
	sort.Sort(hDate.ByYearWeek(keys))

	// 按照排序后的键顺序提取值到切片 
	r.WeekOrderSleepInfo = make([]*WeekSleep, len(keys))
	for i, k := range keys {
	  r.WeekOrderSleepInfo[i] = r.WeekSleepInfo[k]
	}

	// for index, weekSleepInfo := range r.WeekOrderSleepInfo {
	//   klog.InfoS("ConvertWeekSleepInfoToWeekOrderSleepInfo", "index", index, "weekSleepInfo", weekSleepInfo)
	// }

  return nil
}

func (r *Sleep) ConvertMonthSleepInfoToMonthOrderSleepInfo() error {
  klog.InfoS("ConvertMonthSleepInfoToMonthOrderSleepInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.MonthSleepInfo))
	for k := range r.MonthSleepInfo {
	  keys = append(keys, k)
	}
	sort.Sort(hDate.ByYearMonth(keys))

	// 按照排序后的键顺序提取值到切片
	r.MonthOrderSleepInfo = make([]*MonthSleep, len(keys))
	for i, k := range keys {
	  r.MonthOrderSleepInfo[i] = r.MonthSleepInfo[k]
	}

	return nil
}

func (r *Sleep) ConvertYearSleepInfoToYearOrderSleepInfo() error {
  klog.InfoS("ConvertYearSleepInfoToYearOrderSleepInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.YearSleepInfo))
	for k := range r.YearSleepInfo {	  
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 按照排序后的键顺序提取值到切片
	r.YearOrderSleepInfo = make([]*YearSleep, len(keys))
	for i, k := range keys {
	  r.YearOrderSleepInfo[i] = r.YearSleepInfo[k]
	}

	return nil
}