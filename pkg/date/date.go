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
	// 设置日期为那个月的最后一天（通常是下个月的第0天）  
	lastDay := time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, t.Location())  
	return lastDay.Day()   
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





func GetDateDetails(inputDate string) (int, time.Month, int, int, time.Weekday, int, int, int, int, error) {
	// 将输入的日期解析为time.Time类型
	date, err := time.Parse("2006-01-02", inputDate)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0,0, fmt.Errorf("无效的日期格式: %s", err)
	}

	year := date.Year()
	month := date.Month()
	
	weekyear, week := date.ISOWeek()
	weekday := date.Weekday()  // 星期几

	dayOfMonth := date.Day()
	dayOfYear := dayOfYear(date) 

	daysInMonth := daysInMonth(date)
	daysInYear := daysInYear(date)
	
	
	return year, month, weekyear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, nil
}

