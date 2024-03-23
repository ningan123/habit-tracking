package piano

import (
	"fmt"
	"sort"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

var (
	TargetDayPianoTime string = "30min"
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

func (p *Piano) GenPianoInfo() error {
	klog.Info("Generate Piano Info")
	err := p.GenDayPianoInfo()
	if err != nil {
		return err
	}

	err = p.GenWeekPianoInfo()
	if err != nil {
		return err
	}

	err = p.GenMonthPianoInfo()
	if err != nil {
		return err
	}

	err = p.GenYearPianoInfo()
	if err != nil {
		return err
	}

	return nil
}


func (p *Piano) GenDayPianoInfo() error {
	klog.InfoS("GenDayPianoInfo")
  for date, info := range p.RawInfo {
		year, month, weekyear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "month", month, "weekyear", weekyear, "week", week, "weekday", weekday, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)
		
		weekNum := fmt.Sprintf("%d-%02d", weekyear, week)
		dPiano, err := NewDayPiano(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
    if err != nil {
      return err
    }
		
		p.DayPianoInfo[date] = dPiano
  }  
  return nil
}


func (p *Piano) GenWeekPianoInfo() error {
	klog.InfoS("GenWeekPianoInfo")
  for date, info := range p.RawInfo {
		year, month, weekyear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "month", month, "weekyear", weekyear, "week", week, "weekday", weekday, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		weekNum := fmt.Sprintf("%d-%02d", weekyear, week) 
		if p.WeekPianoInfo[weekNum] == nil {
			weekRawInfo := make(map[string]*DayPiano)
			p.WeekPianoInfo[weekNum], err = NewWeekPiano(weekNum, weekRawInfo)
			if err != nil {
			  return err
			}
		}
		
		dPiano, err := NewDayPiano(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
    if err != nil {
      return err
		}

		p.WeekPianoInfo[weekNum].RawInfo[weekday] = dPiano
	   
  }  
  return nil
}


func (p *Piano) GenMonthPianoInfo() error {
	klog.InfoS("GenMonthPianoInfo")
  for date, info := range p.RawInfo {
		year, month, weekyear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "month", month, "weekyear", weekyear, "week", week, "weekday", weekday, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)
		
		monthNum := fmt.Sprintf("%d-%02d", year, month)
		if p.MonthPianoInfo[monthNum] == nil {
		  monthRawInfo := make(map[int]*DayPiano)
			p.MonthPianoInfo[monthNum], err = NewMonthPiano(monthNum, monthRawInfo, daysInMonth)
			if err != nil {
			  return err
			}
		}
		weekNum := fmt.Sprintf("%d-%02d", weekyear, week) 

		dPiano, err := NewDayPiano(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
    if err != nil {
      return err
    }

		p.MonthPianoInfo[monthNum].RawInfo[dayOfMonth] = dPiano	   
  }  
  return nil
}


func (p *Piano) GenYearPianoInfo() error {
	klog.InfoS("GenYearPianoInfo")
  for date, info := range p.RawInfo {
		year, month, weekyear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "month", month, "weekyear", weekyear, "week", week, "weekday", weekday, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		yearNum := fmt.Sprintf("%d", year)
		if p.YearPianoInfo[yearNum] == nil {
			yearRawInfo := make(map[string]*DayPiano)
			p.YearPianoInfo[yearNum], err = NewYearPiano(yearNum, yearRawInfo, daysInYear)
			if err != nil {
			  return err
			}
		}
		weekNum := fmt.Sprintf("%d-%02d", weekyear, week)

		dPiano, err := NewDayPiano(date, year,dayOfYear, month, dayOfMonth, weekNum, weekday, info)
		if err != nil {
			return err
		}

		p.YearPianoInfo[yearNum].YearRawInfo[date] = dPiano
	   
  }  
  return nil
}

// ==============================================
// ==============================================

func (p *Piano) ComputePianoTime() error {
	klog.InfoS("ComputePianoTime")

	err := p.ComputeDayPianoTime()
	if err != nil {
		return err
	}

	err = p.ComputeWeekPianoTime()
	if err != nil {
		return err
	}

	err = p.ComputeMonthPianoTime()
	if err != nil {
		return err
	}

	err = p.ComputeYearPianoTime()
	if err != nil {
		return err
	}

	return nil
}



func (p *Piano) ComputeDayPianoTime() error {
	klog.InfoS("ComputDayPianoTime")
	for _, dPiano := range p.DayPianoInfo {
		err := dPiano.ComputePianoTime()
		if err != nil {
		  return err
		}
	}
  return nil
}


func (p *Piano) ComputeWeekPianoTime() error {
	klog.InfoS("ComputWeekPianoTime")
	for _, wPiano := range p.WeekPianoInfo {
		err := wPiano.ComputePianoTime()
		if err != nil {
		  return err
		}
	}
  return nil
}

func (p *Piano) ComputeMonthPianoTime() error {
	klog.InfoS("ComputMonthPianoTime")
	for _, mPiano := range p.MonthPianoInfo {
		err := mPiano.ComputePianoTime()
		if err != nil {
		  return err
		}
	}
  return nil
}


func (p *Piano) ComputeYearPianoTime() error {
	klog.InfoS("ComputYearPianoTime")
	for _, mPiano := range p.YearPianoInfo {
		err := mPiano.ComputePianoTime()
		if err != nil {
		  return err
		}
	}
  return nil
}


// ==================================================
// ==================================================

func (p *Piano) ConvertPianoInfoToOrderPianoInfo() error {
	klog.InfoS("ConverPianoInfoToOrderPianoInfo")

	err := p.ConvertDayPianoInfoToDayOrderPianoInfo() 
	if err != nil {
		return err
	}

	err = p.ConvertWeekPianoInfoToWeekOrderPianoInfo()
	if err != nil {
		return err
	}

	err = p.ConvertMonthPianoInfoToMonthOrderPianoInfo()
	if err != nil {
		return err
	}

	err = p.ConvertYearPianoInfoToYearOrderPianoInfo()
	if err != nil {
		return err
	}

	return nil
}


func (p *Piano) ConvertDayPianoInfoToDayOrderPianoInfo() error {
	klog.InfoS("ConvertDayPianoInfoToDayOrderPianoInfo")
	// 提取key并排序
	keys := make([]string, 0, len(p.DayPianoInfo))
	for k := range p.DayPianoInfo {
	  keys = append(keys, k)
	}
	sort.Sort(hDate.ByDate(keys))

	// 按照排序后的键顺序提取值到切片 
	p.DayOrderPianoInfo = make([]*DayPiano, len(keys))
	for i, k := range keys {
	  p.DayOrderPianoInfo[i] = p.DayPianoInfo[k]
	}
	return nil
}


func (p *Piano) ConvertWeekPianoInfoToWeekOrderPianoInfo() error {
	klog.InfoS("ConvertWeekPianoInfoToWeekOrderPianoInfo")
	// 提取key并排序
	keys := make([]string, 0, len(p.WeekPianoInfo))
	for k := range p.WeekPianoInfo {
	  keys = append(keys, k)
	}
	sort.Sort(hDate.ByYearWeek(keys))

	// 按照排序后的键顺序提取值到切片 
	p.WeekOrderPianoInfo = make([]*WeekPiano, len(keys))
	for i, k := range keys {
	  p.WeekOrderPianoInfo[i] = p.WeekPianoInfo[k]
	}

  return nil
}


func (p *Piano) ConvertMonthPianoInfoToMonthOrderPianoInfo() error {
	klog.InfoS("ConvertMonthPianoInfoToMonthOrderPianoInfo")
	// 提取key并排序
	keys := make([]string, 0, len(p.MonthPianoInfo))
	for k := range p.MonthPianoInfo {
		keys = append(keys, string(k))
	}
	sort.Sort(hDate.ByYearMonth(keys))

	// 按照排序后的键顺序提取值到切片 
	p.MonthOrderPianoInfo = make([]*MonthPiano, len(keys))
	for i, k := range keys {
		p.MonthOrderPianoInfo[i] = p.MonthPianoInfo[k]
	}

  return nil
}


func (p *Piano) ConvertYearPianoInfoToYearOrderPianoInfo() error {
	klog.InfoS("ConvertYearPianoInfoToYearOrderPianoInfo")
	// 提取key并排序
	keys := make([]string, 0, len(p.YearPianoInfo))
	for k := range p.YearPianoInfo {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 按照排序后的键顺序提取值到切片 
	p.YearOrderPianoInfo = make([]*YearPiano, len(keys))
	for i, k := range keys {
		p.YearOrderPianoInfo[i] = p.YearPianoInfo[k]
	}

  return nil
}

// ==============================================
// ==============================================


func (p *Piano) CheckFinish() error {
	klog.InfoS("CheckFinish")
	err := p.CheckDayFinish()
	if err != nil {
	  return err
	}

	err = p.CheckWeekFinish()
	if err != nil {
	  return err
	}

	err = p.CheckMonthFinish()
	if err != nil {
	  return err
	}

	err = p.CheckYearFinish()
	if err != nil {
	  return err
	}

  return nil
}

func (p *Piano) CheckDayFinish() error {
	klog.InfoS("CheckDayFinish")
	for _, dPiano := range p.DayPianoInfo {
		err := dPiano.CheckFinish()
		if err != nil {
		  return err
		}
	}
  return nil
}

func (p *Piano) CheckWeekFinish() error {
  klog.InfoS("CheckWeekFinish")
	for _, wPiano := range p.WeekPianoInfo {
	  err := wPiano.CheckFinish()
		if err != nil {
		  return err
		}
	}
	return nil
}

func (p *Piano) CheckMonthFinish() error {
  klog.InfoS("CheckMonthFinish")
	for _, mPiano := range p.MonthPianoInfo {
	  err := mPiano.CheckFinish()
		if err != nil {
		  return err
		}
	}
	return nil
}

func (p *Piano) CheckYearFinish() error {
  klog.InfoS("CheckYearFinish")
	for _, yPiano := range p.YearPianoInfo {
	  err := yPiano.CheckFinish()
		if err != nil {
		  return err
		}
	}
	return nil
}

// ==============================================
// ==============================================


func (p *Piano) ComputeExtraPianoTime() error {
  klog.InfoS("ComputeExtraPianoTime")
	err := p.ComputeWeekExtraPianoTime()
	if err != nil {
	  return err
	}
	return nil
}


func (p *Piano) ComputeWeekExtraPianoTime() error {
  klog.InfoS("ComputeWeekExtraPianoTime")
	for _, wPiano := range p.WeekPianoInfo {
	  err := wPiano.ComputeExtraPianoTime()
		if err != nil {
		  return err
		}
	}
	return nil
}


// ==============================================
// ==============================================



func (p *Piano) PrintPianoInfo() {
	klog.InfoS("PrintPianoInfo")

	p.PrintDayPianoInfo()
	p.PrintWeekPianoInfo()
	p.PrintMonthPianoInfo()
	p.PrintYearPianoInfo()
}



func (p *Piano) PrintDayPianoInfo() {
	klog.InfoS("PrintDayPianoInfo")
	for _, dPiano := range p.DayPianoInfo {
		dPiano.Print()
	}
}

func (p *Piano) PrintWeekPianoInfo() {
	klog.InfoS("PrintWeekPianoInfo")
	for _, wPiano := range p.WeekPianoInfo {
		wPiano.Print()
	}
}

func (p *Piano) PrintMonthPianoInfo() {
	klog.InfoS("PrintMonthPianoInfo")
	for _, mPiano := range p.MonthPianoInfo {
		mPiano.Print()
	}
}

func (p *Piano) PrintYearPianoInfo() {
	klog.InfoS("PrintYearPianoInfo")
	for _, yPiano := range p.YearPianoInfo {
		yPiano.Print()
	}
}


