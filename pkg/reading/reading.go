package reading

import (
	"sort"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

var (
	TargetDayReadingTime string = "20min"
)

type Reading struct {
	RawInfo               map[string]string // 原始数据
	YearReadingInfo       map[string]*YearReading
	YearOrderReadingInfo  []*YearReading
	MonthReadingInfo      map[string]*MonthReading
	MonthOrderReadingInfo []*MonthReading
	WeekReadingInfo       map[string]*WeekReading
	WeekOrderReadingInfo  []*WeekReading
	DayReadingInfo        map[string]*DayReading
	DayOrderReadingInfo   []*DayReading
}

func NewReading(rawInfo map[string]string) *Reading {
	return &Reading{
		RawInfo:          rawInfo,
		YearReadingInfo:  make(map[string]*YearReading),
		MonthReadingInfo: make(map[string]*MonthReading),
		WeekReadingInfo:  make(map[string]*WeekReading),
		DayReadingInfo:   make(map[string]*DayReading),
	}
}

// ==============================================
// ==============================================

func (r *Reading) GenReadingInfo() error {
	klog.Info("Generate Reading Info")
	err := r.GenDayReadingInfo()
	if err != nil {
		return err
	}

	err = r.GenWeekReadingInfo()
	if err != nil {
		return err
	}

	err = r.GenMonthReadingInfo()
	if err != nil {
		return err
	}

	err = r.GenYearReadingInfo()
	if err != nil {
		return err
	}

	return nil
}

func (r *Reading) GenDayReadingInfo() error {
	klog.InfoS("GenDayReadingInfo")
	for date, rawInfo := range r.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		dReading, err := NewDayReading(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
		if err != nil {
			return err
		}

		r.DayReadingInfo[date] = dReading
	}
	return nil
}

func (r *Reading) GenWeekReadingInfo() error {
	klog.InfoS("GenWeekReadingInfo")
	for date, rawInfo := range r.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		if r.WeekReadingInfo[weekNum] == nil {
			weekRawInfo := make(map[string]*DayReading)
			r.WeekReadingInfo[weekNum], err = NewWeekReading(weekNum, weekRawInfo)
			if err != nil {
				return err
			}
		}

		dReading, err := NewDayReading(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
		if err != nil {
			return err
		}

		r.WeekReadingInfo[weekNum].RawInfo[weekday] = dReading

	}
	return nil
}

func (r *Reading) GenMonthReadingInfo() error {
	klog.InfoS("GenMonthReadingInfo")
	for date, rawInfo := range r.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		if r.MonthReadingInfo[monthNum] == nil {
			monthRawInfo := make(map[int]*DayReading)
			r.MonthReadingInfo[monthNum], err = NewMonthReading(monthNum, daysInMonth, monthRawInfo)
			if err != nil {
				return err
			}
		}

		dReading, err := NewDayReading(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
		if err != nil {
			return err
		}

		r.MonthReadingInfo[monthNum].RawInfo[dayOfMonth] = dReading
	}
	return nil
}

func (r *Reading) GenYearReadingInfo() error {
	klog.InfoS("GenYearReadingInfo")
	for date, rawInfo := range r.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		if r.YearReadingInfo[yearNum] == nil {
			yearRawInfo := make(map[string]*DayReading)
			r.YearReadingInfo[yearNum], err = NewYearReading(yearNum, daysInYear, yearRawInfo)
			if err != nil {
				return err
			}
		}

		dReading, err := NewDayReading(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
		if err != nil {
			return err
		}

		r.YearReadingInfo[yearNum].YearRawInfo[date] = dReading

	}
	return nil
}

// ==============================================
// ==============================================

func (r *Reading) ComputeReadingTime() error {
	klog.InfoS("ComputeReadingTime")

	err := r.ComputeDayReadingTime()
	if err != nil {
		return err
	}

	err = r.ComputeWeekReadingTime()
	if err != nil {
		return err
	}

	err = r.ComputeMonthReadingTime()
	if err != nil {
		return err
	}

	err = r.ComputeYearReadingTime()
	if err != nil {
		return err
	}

	return nil
}

func (r *Reading) ComputeDayReadingTime() error {
	klog.InfoS("ComputDayReadingTime")
	for _, dReading := range r.DayReadingInfo {
		err := dReading.ComputeReadingTime()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Reading) ComputeWeekReadingTime() error {
	klog.InfoS("ComputWeekReadingTime")
	for _, wReading := range r.WeekReadingInfo {
		err := wReading.ComputeReadingTime()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Reading) ComputeMonthReadingTime() error {
	klog.InfoS("ComputMonthReadingTime")
	for _, mReading := range r.MonthReadingInfo {
		err := mReading.ComputeReadingTime()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Reading) ComputeYearReadingTime() error {
	klog.InfoS("ComputYearReadingTime")
	for _, mReading := range r.YearReadingInfo {
		err := mReading.ComputeReadingTime()
		if err != nil {
			return err
		}
	}
	return nil
}

// ==================================================
// ==================================================

func (r *Reading) ConvertReadingInfoToOrderReadingInfo() error {
	klog.InfoS("ConverReadingInfoToOrderReadingInfo")

	err := r.ConvertDayReadingInfoToDayOrderReadingInfo()
	if err != nil {
		return err
	}

	err = r.ConvertWeekReadingInfoToWeekOrderReadingInfo()
	if err != nil {
		return err
	}

	err = r.ConvertMonthReadingInfoToMonthOrderReadingInfo()
	if err != nil {
		return err
	}

	err = r.ConvertYearReadingInfoToYearOrderReadingInfo()
	if err != nil {
		return err
	}

	return nil
}

func (r *Reading) ConvertDayReadingInfoToDayOrderReadingInfo() error {
	klog.InfoS("ConvertDayReadingInfoToDayOrderReadingInfo")
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
	klog.InfoS("ConvertWeekReadingInfoToWeekOrderReadingInfo")
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
	klog.InfoS("ConvertMonthReadingInfoToMonthOrderReadingInfo")
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

func (r *Reading) ConvertYearReadingInfoToYearOrderReadingInfo() error {
	klog.InfoS("ConvertYearReadingInfoToYearOrderReadingInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.YearReadingInfo))
	for k := range r.YearReadingInfo {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 按照排序后的键顺序提取值到切片
	r.YearOrderReadingInfo = make([]*YearReading, len(keys))
	for i, k := range keys {
		r.YearOrderReadingInfo[i] = r.YearReadingInfo[k]
	}

	return nil
}

// ==============================================
// ==============================================

func (r *Reading) CheckFinish() error {
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

func (r *Reading) CheckWeekFinish() error {
	klog.InfoS("CheckWeekFinish")
	for _, wReading := range r.WeekReadingInfo {
		err := wReading.CheckFinish()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Reading) CheckMonthFinish() error {
	klog.InfoS("CheckMonthFinish")
	for _, mReading := range r.MonthReadingInfo {
		err := mReading.CheckFinish()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Reading) CheckYearFinish() error {
	klog.InfoS("CheckYearFinish")
	for _, yReading := range r.YearReadingInfo {
		err := yReading.CheckFinish()
		if err != nil {
			return err
		}
	}
	return nil
}

// ==============================================
// ==============================================

func (r *Reading) ComputeExtraReadingTime() error {
	klog.InfoS("ComputeExtraReadingTime")
	err := r.ComputeWeekExtraReadingTime()
	if err != nil {
		return err
	}

	err = r.ComputeMonthExtraReadingTime()
	if err != nil {
		return err
	}

	err = r.ComputeYearExtraReadingTime()
	if err != nil {
		return err
	}
	return nil
}

func (r *Reading) ComputeWeekExtraReadingTime() error {
	klog.InfoS("ComputeWeekExtraReadingTime")
	for _, wReading := range r.WeekReadingInfo {
		err := wReading.ComputeExtraReadingTime()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Reading) ComputeMonthExtraReadingTime() error {
	klog.InfoS("ComputeMonthExtraReadingTime")
	for _, mReading := range r.MonthReadingInfo {
		err := mReading.ComputeExtraReadingTime()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Reading) ComputeYearExtraReadingTime() error {
	klog.InfoS("ComputeYearExtraReadingTime")
	for _, yReading := range r.YearReadingInfo {
		err := yReading.ComputeExtraReadingTime()
		if err != nil {
			return err
		}
	}
	return nil
}

// ==============================================
// ==============================================

func (r *Reading) PrintReadingInfo() {
	klog.InfoS("PrintReadingInfo")

	r.PrintDayReadingInfo()
	r.PrintWeekReadingInfo()
	r.PrintMonthReadingInfo()
	r.PrintYearReadingInfo()
}

func (r *Reading) PrintDayReadingInfo() {
	klog.InfoS("PrintDayReadingInfo")
	for _, dReading := range r.DayReadingInfo {
		dReading.Print()
	}
}

func (r *Reading) PrintWeekReadingInfo() {
	klog.InfoS("PrintWeekReadingInfo")
	for _, wReading := range r.WeekReadingInfo {
		wReading.Print()
	}
}

func (r *Reading) PrintMonthReadingInfo() {
	klog.InfoS("PrintMonthReadingInfo")
	for _, mReading := range r.MonthReadingInfo {
		mReading.Print()
	}
}

func (r *Reading) PrintYearReadingInfo() {
	klog.InfoS("PrintYearReadingInfo")
	for _, yReading := range r.YearReadingInfo {
		yReading.Print()
	}
}
