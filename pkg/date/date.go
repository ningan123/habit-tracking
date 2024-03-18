package date

import (
	"fmt"
	"time"
)

func convertDateStrToDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

func convertDateToDateStr(date time.Time) string {
	return date.Format("2006-01-02")
}



func dateDetail(date time.Time) (int, time.Month, int) {
	year := date.Year() // 年份，月份，日
	month := date.Month()
	day := date.Day()

	return year, month, day
}

func dateDetail2(date time.Time) (int, int, int) {
	year, month, day := date.Date()
	return year, int(month), day
}



// daysInMonth 返回给定时间的月份中的天数  
func daysInMonth(t time.Time) int {  
	// 获取下个月的第一天，然后减去一天，就得到了当前月份的最后一天  
	nextMonth := t.AddDate(0, 1, 0)  
	lastDayOfMonth := nextMonth.AddDate(0, 0, -1)  
	return lastDayOfMonth.Day()  
}  


// isLeapYear 判断给定年份是否是闰年  
func isLeapYear(year int) bool {  
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)  
}  
  
// daysInYear 返回给定年份的天数  
func daysInYear(t time.Time) int {  
	if isLeapYear(t.Year()) {  
		return 366 // 闰年有366天  
	}  
	return 365 // 平年有365天  
}  

// dayOfYear 计算给定日期是一年中的第几天
func dayOfYear(t time.Time) int {  
	return t.YearDay()
} 


func dateWeek(date time.Time) (int, int, time.Weekday) {
	year, week := date.ISOWeek()
	weekDay := date.Weekday()

	return year, week, weekDay
}

// TODO：2024-12-30的输出结果不太对
func GetDateDetails(inputDate string) (int, int, time.Month, int, int, time.Weekday, error) {
	// 将输入的日期解析为time.Time类型
	date, err := time.Parse("2006-01-02", inputDate)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, fmt.Errorf("无效的日期格式: %s", err)
	}

	// 获取年份、月份、周数和星期几
	year := date.Year()
	dayOfYear := dayOfYear(date)
	month := date.Month()
	dayOfMonth := date.Day()
	_, weekNumber := date.ISOWeek()
	weekday := date.Weekday()
	

	// 如果日期在第一周，但是年份不同，需要修正周数
	if weekNumber == 1 && date.Month() == 12 {
		year++
	}

	return year, dayOfYear, month, dayOfMonth, weekNumber, weekday, nil
}

