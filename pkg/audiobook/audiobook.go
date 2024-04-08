package audiobook

import (
	"sort"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

var (
	TargetWeekFinishBooks int = 2
)

type Audiobook struct {
	RawInfo map[string]string // 原始数据
	YearInfo map[string]*YearAudiobook
	YearOrderInfo []*YearAudiobook
	MonthInfo map[string]*MonthAudiobook
	MonthOrderInfo []*MonthAudiobook
	WeekInfo map[string]*WeekAudiobook
	WeekOrderInfo []*WeekAudiobook
  DayInfo map[string]*DayAudiobook 
	DayOrderInfo []*DayAudiobook
}


func NewAudiobook(rawInfo map[string]string) *Audiobook {
	return &Audiobook{
	  RawInfo: rawInfo,
		YearInfo: make(map[string]*YearAudiobook),
		MonthInfo: make(map[string]*MonthAudiobook),
	  WeekInfo: make(map[string]*WeekAudiobook),
    DayInfo: make(map[string]*DayAudiobook),
	}
}


// ==============================================
// ==============================================

func(r *Audiobook) GenInfo() error {
	klog.Info("Generate Audiobook Info")
	err := r.GenDayInfo()
	if err != nil {
		return err
	}

	err = r.GenWeekInfo()
	if err != nil {
		return err
	}

	err = r.GenMonthInfo()
	if err != nil {
		return err
	}

	err = r.GenYearInfo()
	if err != nil {
		return err
	}

	return nil
}


func(r *Audiobook) GenDayInfo() error {
	klog.InfoS("GenDayInfo")
  for date, rawInfo := range r.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)
		
		dAudiobook, err := NewDayAudiobook(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
    if err != nil {
      return err
    }
		
		r.DayInfo[date] = dAudiobook
  }  
  return nil
}


func(r *Audiobook) GenWeekInfo() error {
	klog.InfoS("GenWeekInfo")
  for date, rawInfo := range r.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		if r.WeekInfo[weekNum] == nil {
			weekRawInfo := make(map[string]*DayAudiobook)
			r.WeekInfo[weekNum], err = NewWeekAudiobook(weekNum, weekRawInfo)
			if err != nil {
			  return err
			}
		}
		
		dAudiobook, err := NewDayAudiobook(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
    if err != nil {
      return err
		}

		r.WeekInfo[weekNum].RawInfo[weekday] = dAudiobook
	   
  }  
  return nil
}


func(r *Audiobook) GenMonthInfo() error {
	klog.InfoS("GenMonthInfo")
  for date, rawInfo := range r.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)
		
		if r.MonthInfo[monthNum] == nil {
		  monthRawInfo := make(map[int]*DayAudiobook)
			r.MonthInfo[monthNum], err = NewMonthAudiobook(monthNum, daysInMonth, monthRawInfo)
			if err != nil {
			  return err
			}
		}
		
		dAudiobook, err := NewDayAudiobook(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
    if err != nil {
      return err
    }

		r.MonthInfo[monthNum].RawInfo[dayOfMonth] = dAudiobook	   
  }  
  return nil
}


func(r *Audiobook) GenYearInfo() error {
	klog.InfoS("GenYearInfo")
  for date, rawInfo := range r.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		if r.YearInfo[yearNum] == nil {
			yearRawInfo := make(map[string]*DayAudiobook)
			r.YearInfo[yearNum], err = NewYearAudiobook(yearNum, yearRawInfo, daysInYear)
			if err != nil {
			  return err
			}
		}
		
		dAudiobook, err := NewDayAudiobook(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
		if err != nil {
			return err
		}

		r.YearInfo[yearNum].RawInfo[date] = dAudiobook
	   
  }  
  return nil
}

// ==============================================
// ==============================================

func(r *Audiobook) ComputeFinishBooks() error {
	klog.InfoS("ComputeFinishBooks")

	err := r.ComputeDayFinishBooks()
	if err != nil {
		return err
	}

	err = r.ComputeWeekFinishBooks()
	if err != nil {
		return err
	}

	err = r.ComputeMonthFinishBooks()
	if err != nil {
		return err
	}

	err = r.ComputeYearFinishBooks()
	if err != nil {
		return err
	}

	return nil
}



func(r *Audiobook) ComputeDayFinishBooks() error {
	klog.InfoS("ComputDayAudiobookTime")
	for _, dAudiobook := range r.DayInfo {
		err := dAudiobook.ComputeFinishBooks()
		if err != nil {
		  return err
		}
	}
  return nil
}


func(r *Audiobook) ComputeWeekFinishBooks() error {
	klog.InfoS("ComputWeekAudiobookTime")
	for _, wAudiobook := range r.WeekInfo {
		err := wAudiobook.ComputeFinishBooks()
		if err != nil {
		  return err
		}
	}
  return nil
}

func(r *Audiobook) ComputeMonthFinishBooks() error {
	klog.InfoS("ComputMonthAudiobookTime")
	for _, mAudiobook := range r.MonthInfo {
		err := mAudiobook.ComputeFinishBooks()
		if err != nil {
		  return err
		}
	}
  return nil
}


func(r *Audiobook) ComputeYearFinishBooks() error {
	klog.InfoS("ComputYearAudiobookTime")
	for _, mAudiobook := range r.YearInfo {
		err := mAudiobook.ComputeFinishBooks()
		if err != nil {
		  return err
		}
	}
  return nil
}


// ==================================================
// ==================================================

func (r *Audiobook) ConvertInfoToOrderInfo() error {
	klog.InfoS("ConverInfoToOrderInfo")

	err := r.ConvertDayInfoToDayOrderInfo() 
	if err != nil {
		return err
	}

	err = r.ConvertWeekInfoToWeekOrderInfo()
	if err != nil {
		return err
	}

	err = r.ConvertMonthInfoToMonthOrderInfo()
	if err != nil {
		return err
	}

	err = r.ConvertYearInfoToYearOrderInfo()
	if err != nil {
		return err
	}

	return nil
}


func (r *Audiobook) ConvertDayInfoToDayOrderInfo() error {
	klog.InfoS("ConvertDaInfoToDayOrderInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.DayInfo))
	for k := range r.DayInfo {
	  keys = append(keys, k)
	}
	sort.Sort(hDate.ByDate(keys))

	// 按照排序后的键顺序提取值到切片 
	r.DayOrderInfo = make([]*DayAudiobook, len(keys))
	for i, k := range keys {
	  r.DayOrderInfo[i] = r.DayInfo[k]
	}
	return nil
}


func (r *Audiobook) ConvertWeekInfoToWeekOrderInfo() error {
	klog.InfoS("ConvertWeekInfoToWeekOrderInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.WeekInfo))
	for k := range r.WeekInfo {
	  keys = append(keys, k)
	}
	sort.Sort(hDate.ByYearWeek(keys))

	// 按照排序后的键顺序提取值到切片 
	r.WeekOrderInfo = make([]*WeekAudiobook, len(keys))
	for i, k := range keys {
	  r.WeekOrderInfo[i] = r.WeekInfo[k]
	}

  return nil
}


func (r *Audiobook) ConvertMonthInfoToMonthOrderInfo() error {
	klog.InfoS("ConvertMonthInfoToMonthOrderInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.MonthInfo))
	for k := range r.MonthInfo {
		keys = append(keys, string(k))
	}
	sort.Sort(hDate.ByYearMonth(keys))

	// 按照排序后的键顺序提取值到切片 
	r.MonthOrderInfo = make([]*MonthAudiobook, len(keys))
	for i, k := range keys {
		r.MonthOrderInfo[i] = r.MonthInfo[k]
	}

  return nil
}


func (r *Audiobook) ConvertYearInfoToYearOrderInfo() error {
	klog.InfoS("ConvertYearInfoToYearOrderInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.YearInfo))
	for k := range r.YearInfo {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 按照排序后的键顺序提取值到切片 
	r.YearOrderInfo = make([]*YearAudiobook, len(keys))
	for i, k := range keys {
		r.YearOrderInfo[i] = r.YearInfo[k]
	}

  return nil
}

// ==============================================
// ==============================================


func (r *Audiobook) CheckFinish() error {
	klog.InfoS("CheckFinish")
	err := r.CheckDayFinish()
	if err != nil {
	  return err
	}

	err = r.CheckWeekFinish()
	if err != nil {
	  return err
	}

	err = r.CheckMonthFinish()
	if err != nil {
	  return err
	}

  return nil
}

func (r *Audiobook) CheckDayFinish() error {
	klog.InfoS("CheckDayFinish")
	for _, dAudiobook := range r.DayInfo {
		err := dAudiobook.CheckFinish()
		if err != nil {
		  return err
		}
	}
  return nil
}

func (r *Audiobook) CheckWeekFinish() error {
  klog.InfoS("CheckWeekFinish")
	for _, wAudiobook := range r.WeekInfo {
	  err := wAudiobook.CheckFinish()
		if err != nil {
		  return err
		}
	}
	return nil
}


func (r *Audiobook) CheckMonthFinish() error {
  klog.InfoS("CheckMonthFinish")
	for _, mAudiobook := range r.MonthInfo {
	  err := mAudiobook.CheckFinish()
		if err != nil {
		  return err
		}
	}
	return nil
}


func (r *Audiobook) CheckYearFinish() error {
  klog.InfoS("CheckYearFinish")
	for _, yAudiobook := range r.YearInfo {
	  err := yAudiobook.CheckFinish()
		if err != nil {
		  return err
		}
	}
	return nil
}



// ==============================================
// ==============================================



func(r *Audiobook) PrintAudiobookInfo() {
	klog.InfoS("PrintAudiobookInfo")

	r.PrintDayInfo()
	r.PrintWeekInfo()
	r.PrintMonthInfo()
	r.PrintYearInfo()
}



func(r *Audiobook) PrintDayInfo() {
	klog.InfoS("PrintDayInfo")
	for _, dAudiobook := range r.DayInfo {
		dAudiobook.Print()
	}
}

func(r *Audiobook) PrintWeekInfo() {
	klog.InfoS("PrintWeekInfo")
	for _, wAudiobook := range r.WeekInfo {
		wAudiobook.Print()
	}
}

func(r *Audiobook) PrintMonthInfo() {
	klog.InfoS("PrintMonthInfo")
	for _, mAudiobook := range r.MonthInfo {
		mAudiobook.Print()
	}
}

func(r *Audiobook) PrintYearInfo() {
	klog.InfoS("PrintYearInfo")
	for _, yAudiobook := range r.YearInfo {
		yAudiobook.Print()
	}
}


