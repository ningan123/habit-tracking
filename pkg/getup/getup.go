package getup

import (
	"fmt"

	"k8s.io/klog/v2"
	hDate "ningan.com/habit-tracking/pkg/date"
)

var (
	TargetDayGetupTime string = "7:30"
)

type Getup struct {
	RawInfo map[string]string // 原始数据
	DayGetupInfo map[string]*DayGetup
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
