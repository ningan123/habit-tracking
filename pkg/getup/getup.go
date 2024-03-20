package getup

import (
	"fmt"
	"sort"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

var (
	TargetDayGetupTime string = "7:30"
)

type Getup struct {
	RawInfo map[string]string // 原始数据
	DayGetupInfo map[string]*DayGetup
	DayOrderGetupInfo []*DayGetup
}


func NewGetup(rawInfo map[string]string) *Getup {
	return &Getup{
		RawInfo: rawInfo,
		DayGetupInfo: make(map[string]*DayGetup),
	}
}

func(g *Getup) GenGetupInfo() error {
	klog.Info("GenGetupInfo")
	err := g.GenDayGetupInfo()
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


func (g *Getup) CheckFinish() error {
	klog.InfoS("CheckFinish")
	err := g.CheckDayFinish()
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





// ==================================================
// ==================================================

func (g *Getup) ConvertGetupInfoToOrderGetupInfo() error {
	klog.InfoS("ConverGetupInfoToOrderGetupInfo")

	err := g.ConvertDayGetupInfoToDayOrderGetupInfo() 
	if err != nil {
		return err
	}

	// err = r.ConvertWeekGetupInfoToWeekOrderGetupInfo()
	// if err != nil {
	// 	return err
	// }

	// err = r.ConvertMonthGetupInfoToMonthOrderGetupInfo()
	// if err != nil {
	// 	return err
	// }

	// err = r.ConvertYearGetupInfoToYearOrderGetupInfo()
	// if err != nil {
	// 	return err
	// }

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