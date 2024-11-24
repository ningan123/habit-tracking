package sport

import (
	"sort"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

var (
	TargetWeekSportDays  int = 3
	TargetMonthSportDays int = 3 * 4
	TargetYearSportDays  int = 3 * 52
)

type Sport struct {
	RawInfo        map[string]string // 原始数据
	YearInfo       map[string]*YearSport
	YearOrderInfo  []*YearSport
	MonthInfo      map[string]*MonthSport
	MonthOrderInfo []*MonthSport
	WeekInfo       map[string]*WeekSport
	WeekOrderInfo  []*WeekSport
	DayInfo        map[string]*DaySport
	DayOrderInfo   []*DaySport
}

func NewSport(rawInfo map[string]string) *Sport {
	return &Sport{
		RawInfo:   rawInfo,
		YearInfo:  make(map[string]*YearSport),
		MonthInfo: make(map[string]*MonthSport),
		WeekInfo:  make(map[string]*WeekSport),
		DayInfo:   make(map[string]*DaySport),
	}
}

// ==============================================
// ==============================================

func (r *Sport) GenSportInfo() error {
	klog.Info("Generate sport Info")
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

func (r *Sport) GenDayInfo() error {
	klog.InfoS("GenDayInfo")
	for date, rawInfo := range r.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		dsport, err := NewDaySport(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
		if err != nil {
			return err
		}

		r.DayInfo[date] = dsport
	}
	return nil
}

func (r *Sport) GenWeekInfo() error {
	klog.InfoS("GenWeekInfo")
	for date, rawInfo := range r.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		if r.WeekInfo[weekNum] == nil {
			weekRawInfo := make(map[string]*DaySport)
			r.WeekInfo[weekNum], err = NewWeekSport(weekNum, weekRawInfo)
			if err != nil {
				return err
			}
		}

		dsport, err := NewDaySport(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
		if err != nil {
			return err
		}

		r.WeekInfo[weekNum].RawInfo[weekday] = dsport

	}
	return nil
}

func (r *Sport) GenMonthInfo() error {
	klog.InfoS("GenMonthInfo")
	for date, rawInfo := range r.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		if r.MonthInfo[monthNum] == nil {
			monthRawInfo := make(map[int]*DaySport)
			r.MonthInfo[monthNum], err = NewMonthSport(monthNum, daysInMonth, monthRawInfo)
			if err != nil {
				return err
			}
		}

		dsport, err := NewDaySport(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
		if err != nil {
			return err
		}

		r.MonthInfo[monthNum].RawInfo[dayOfMonth] = dsport
	}
	return nil
}

func (r *Sport) GenYearInfo() error {
	klog.InfoS("GenYearInfo")
	for date, rawInfo := range r.RawInfo {
		year, yearNum, month, monthNum, weekyear, week, weekday, weekNum, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := hDate.GetDateDetails(date)
		if err != nil {
			return err
		}
		klog.V(2).InfoS("date detail", "date", date, "year", year, "yearNum", yearNum, "month", month, "monthNum", monthNum, "weekyear", weekyear, "week", week, "weekday", weekday, "weekNum", weekNum, "dayOfMonth", dayOfMonth, "dayOfYear", dayOfYear, "daysInMonth", daysInMonth, "daysInYear", daysInYear)

		if r.YearInfo[yearNum] == nil {
			yearRawInfo := make(map[string]*DaySport)
			r.YearInfo[yearNum], err = NewYearSport(yearNum, daysInYear, yearRawInfo)
			if err != nil {
				return err
			}
		}

		dsport, err := NewDaySport(date, weekday, weekNum, monthNum, yearNum, dayOfMonth, dayOfYear, rawInfo)
		if err != nil {
			return err
		}

		r.YearInfo[yearNum].RawInfo[date] = dsport

	}
	return nil
}

// ==============================================
// ==============================================

func (r *Sport) ComputeSportTimes() error {
	klog.InfoS("ComputeSportTimes")

	err := r.ComputeDaySportTimes()
	if err != nil {
		return err
	}

	err = r.ComputeWeekSportTime()
	if err != nil {
		return err
	}

	err = r.ComputeMonthSportTime()
	if err != nil {
		return err
	}

	err = r.ComputeYearSportTime()
	if err != nil {
		return err
	}

	return nil
}

func (r *Sport) ComputeDaySportTimes() error {
	klog.InfoS("ComputDaySportTime")
	for _, dsport := range r.DayInfo {
		err := dsport.ComputeSportTimes()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Sport) ComputeWeekSportTime() error {
	klog.InfoS("ComputWeekSportTime")
	for _, wsport := range r.WeekInfo {
		err := wsport.ComputeSportTimes()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Sport) ComputeMonthSportTime() error {
	klog.InfoS("ComputMonthSportTime")
	for _, msport := range r.MonthInfo {
		err := msport.ComputeSportTimes()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Sport) ComputeYearSportTime() error {
	klog.InfoS("ComputYearSportTime")
	for _, msport := range r.YearInfo {
		err := msport.ComputeSportTimes()
		if err != nil {
			return err
		}
	}
	return nil
}

// ==================================================
// ==================================================

func (r *Sport) ConvertSportInfoToOrderSportInfo() error {
	klog.InfoS("ConversportInfoToOrdersportInfo")

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

func (r *Sport) ConvertDayInfoToDayOrderInfo() error {
	klog.InfoS("ConvertDayInfoToDayOrderInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.DayInfo))
	for k := range r.DayInfo {
		keys = append(keys, k)
	}
	sort.Sort(hDate.ByDate(keys))

	// 按照排序后的键顺序提取值到切片
	r.DayOrderInfo = make([]*DaySport, len(keys))
	for i, k := range keys {
		r.DayOrderInfo[i] = r.DayInfo[k]
	}
	return nil
}

func (r *Sport) ConvertWeekInfoToWeekOrderInfo() error {
	klog.InfoS("ConvertWeekInfoToWeekOrderInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.WeekInfo))
	for k := range r.WeekInfo {
		keys = append(keys, k)
	}
	sort.Sort(hDate.ByYearWeek(keys))

	// 按照排序后的键顺序提取值到切片
	r.WeekOrderInfo = make([]*WeekSport, len(keys))
	for i, k := range keys {
		r.WeekOrderInfo[i] = r.WeekInfo[k]
	}

	return nil
}

func (r *Sport) ConvertMonthInfoToMonthOrderInfo() error {
	klog.InfoS("ConvertMonthInfoToMonthOrderInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.MonthInfo))
	for k := range r.MonthInfo {
		keys = append(keys, string(k))
	}
	sort.Sort(hDate.ByYearMonth(keys))

	// 按照排序后的键顺序提取值到切片
	r.MonthOrderInfo = make([]*MonthSport, len(keys))
	for i, k := range keys {
		r.MonthOrderInfo[i] = r.MonthInfo[k]
	}

	return nil
}

func (r *Sport) ConvertYearInfoToYearOrderInfo() error {
	klog.InfoS("ConvertYearInfoToYearOrderInfo")
	// 提取key并排序
	keys := make([]string, 0, len(r.YearInfo))
	for k := range r.YearInfo {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 按照排序后的键顺序提取值到切片
	r.YearOrderInfo = make([]*YearSport, len(keys))
	for i, k := range keys {
		r.YearOrderInfo[i] = r.YearInfo[k]
	}

	return nil
}

// ==============================================
// ==============================================

func (r *Sport) CheckFinish() error {
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

func (r *Sport) CheckDayFinish() error {
	klog.InfoS("CheckDayFinish")
	for _, dsport := range r.DayInfo {
		err := dsport.CheckFinish()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Sport) CheckWeekFinish() error {
	klog.InfoS("CheckWeekFinish")
	for _, wsport := range r.WeekInfo {
		err := wsport.CheckFinish()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Sport) CheckMonthFinish() error {
	klog.InfoS("CheckMonthFinish")
	for _, msport := range r.MonthInfo {
		err := msport.CheckFinish()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Sport) CheckYearFinish() error {
	klog.InfoS("CheckYearFinish")
	for _, ysport := range r.YearInfo {
		err := ysport.CheckFinish()
		if err != nil {
			return err
		}
	}
	return nil
}

// ==============================================
// ==============================================

func (r *Sport) PrintSportInfo() {
	klog.InfoS("PrintSportInfo")

	r.PrintDayInfo()
	r.PrintWeekInfo()
	r.PrintMonthInfo()
	r.PrintYearInfo()
}

func (r *Sport) PrintDayInfo() {
	klog.InfoS("PrintDayInfo")
	for _, dsport := range r.DayInfo {
		dsport.Print()
	}
}

func (r *Sport) PrintWeekInfo() {
	klog.InfoS("PrintWeekInfo")
	for _, wsport := range r.WeekInfo {
		wsport.Print()
	}
}

func (r *Sport) PrintMonthInfo() {
	klog.InfoS("PrintMonthInfo")
	for _, msport := range r.MonthInfo {
		msport.Print()
	}
}

func (r *Sport) PrintYearInfo() {
	klog.InfoS("PrintYearInfo")
	for _, ysport := range r.YearInfo {
		ysport.Print()
	}
}
