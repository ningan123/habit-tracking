package route

import (
	"fmt"
	"net/http"

	hData "ningan.com/habit-tracking/pkg/data"
)


func DayReadingHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>date</th><th>weekNum</th><th>weekday</th><th>dayReadingTime</th><th>targetReadingTime</th><th>content</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalReading.DayOrderReadingInfo { 
		if item == nil {
		  continue
		}
		fmt.Fprintf(w, "<tr>")
		fmt.Fprintf(w, "<td>%s</td>", item.Date)
		fmt.Fprintf(w, "<td>%s</td>", item.WeekNum)
		fmt.Fprintf(w, "<td>%s</td>", item.Weekday)
		fmt.Fprintf(w, "<td>%s</td>", item.ReadingTime) 
		fmt.Fprintf(w, "<td>%s</td>", item.TargetReadingTime) 
		fmt.Fprintf(w, "<td>%s</td>", item.ReadingTimeOfDifferentContentStr)
		if item.IsFinish {
			fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
			fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		}
		
		fmt.Fprintf(w, "</tr>\n")
		

		// for content, conReadingTime := range item.DayReadingTimeOfDifferentContent {
		// 	fmt.Fprintf(w, "<tr>")  
		// 	fmt.Fprintf(w, "<td>%s</td>", item.DayDate)
		// 	fmt.Fprintf(w, "<td>%s</td>", item.DayReadingTime)  
		// 	fmt.Fprintf(w, "<td>%s</td>", content)  
		// 	fmt.Fprintf(w, "<td>%s</td>", conReadingTime)  
		// 	fmt.Fprintf(w, "</tr>\n")  
		// }		
	}  
  
	// 构造HTML表格的结尾  
	fmt.Fprintf(w, "</table>\n")  
	fmt.Fprintf(w, "</body>\n")  
	fmt.Fprintf(w, "</html>\n")  	
}

func WeekReadingHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>weekNum</th><th>weekReadingTime</th><th>targetReadingTime</th><th>extraReadingTime</th><th>content</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalReading.WeekOrderReadingInfo { 
		if item == nil {
			continue 
		}  
		
		fmt.Fprintf(w, "<tr>")  
		fmt.Fprintf(w, "<td>%s</td>", item.WeekNum)
		fmt.Fprintf(w, "<td>%s</td>", item.WeekReadingTime)  
		fmt.Fprintf(w, "<td>%s</td>", item.TargetReadingTime) 
		fmt.Fprintf(w, "<td>%s</td>", item.ExtraReadingTime) 
		fmt.Fprintf(w, "<td>%s</td>", item.WeekReadingTimeOfDifferentContentStr) 		
		if item.IsFinish {
			fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
			fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		}  
		fmt.Fprintf(w, "</tr>\n")  
		

		// for content, conReadingTime := range item.WeekReadingTimeOfDifferentContent {
		// 	fmt.Fprintf(w, "<tr>")  
		// 	fmt.Fprintf(w, "<td>%d</td>", item.WeekNum)
		// 	fmt.Fprintf(w, "<td>%s</td>", item.WeekReadingTime)  
		// 	fmt.Fprintf(w, "<td>%s</td>", content)  
		// 	fmt.Fprintf(w, "<td>%s</td>", conReadingTime)  
		// 	fmt.Fprintf(w, "</tr>\n")  
		// }		
	}  
  
	// 构造HTML表格的结尾  
	fmt.Fprintf(w, "</table>\n")  
	fmt.Fprintf(w, "</body>\n")  
	fmt.Fprintf(w, "</html>\n")  
}


func MonthReadingHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>monthNum</th><th>monthReadingTime</th><th>targetReadingTime</th><th>content</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalReading.MonthOrderReadingInfo { 
		if item == nil {
			continue 
		} 

		fmt.Fprintf(w, "<tr>")  
		fmt.Fprintf(w, "<td>%s</td>", item.MonthNum)
		fmt.Fprintf(w, "<td>%s</td>", item.MonthReadingTime)  
		fmt.Fprintf(w, "<td>%s</td>", item.TargetReadingTime) 
		fmt.Fprintf(w, "<td>%s</td>", item.MonthReadingTimeOfDifferentContentStr)
		if item.IsFinish {
			fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
			fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		}  
		fmt.Fprintf(w, "</tr>\n")  

		// for content, conReadingTime := range item.MonthReadingTimeOfDifferentContent {
		// 	fmt.Fprintf(w, "<tr>")  
		// 	fmt.Fprintf(w, "<td>%d</td>", item.MonthNum)
		// 	fmt.Fprintf(w, "<td>%s</td>", item.MonthReadingTime)  
		// 	fmt.Fprintf(w, "<td>%s</td>", content)  
		// 	fmt.Fprintf(w, "<td>%s</td>", conReadingTime)  
		// 	fmt.Fprintf(w, "</tr>\n")  
		// }		
	}  
  
	// 构造HTML表格的结尾  
	fmt.Fprintf(w, "</table>\n")  
	fmt.Fprintf(w, "</body>\n")  
	fmt.Fprintf(w, "</html>\n")  
}


func YearReadingHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>yearNum</th><th>yearReadingTime</th><th>targetReadingTime</th><th>content</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalReading.YearOrderReadingInfo { 
		if item == nil {
			continue
		} 
		fmt.Fprintf(w, "<tr>")  
		fmt.Fprintf(w, "<td>%s</td>", item.YearNum)
		fmt.Fprintf(w, "<td>%s</td>", item.YearReadingTime)  
		fmt.Fprintf(w, "<td>%s</td>", item.TargetReadingTime)  
		fmt.Fprintf(w, "<td>%s</td>", item.YearReadingTimeOfDifferentContentStr) 
		if item.IsFinish {
			fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
			fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		}   
		fmt.Fprintf(w, "</tr>\n")  

		// for content, conReadingTime := range item.YearReadingTimeOfDifferentContent {
		// 	fmt.Fprintf(w, "<tr>")  
		// 	fmt.Fprintf(w, "<td>%d</td>", item.YearNum)
		// 	fmt.Fprintf(w, "<td>%s</td>", item.YearReadingTime)  
		// 	fmt.Fprintf(w, "<td>%s</td>", content)  
		// 	fmt.Fprintf(w, "<td>%s</td>", conReadingTime)  
		// 	fmt.Fprintf(w, "</tr>\n")  
		// }		
	}  
  
	// 构造HTML表格的结尾  
	fmt.Fprintf(w, "</table>\n")  
	fmt.Fprintf(w, "</body>\n")  
	fmt.Fprintf(w, "</html>\n")  
}