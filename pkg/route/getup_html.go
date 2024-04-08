package route

import (
	"fmt"
	"net/http"

	hData "ningan.com/habit-tracking/pkg/data"
)


func DayGetupHtmlTable(w http.ResponseWriter) {
	// 开始写入HTML  
	fmt.Fprintln(w, GlobalTable)  

	// // 构造HTML表格的开头  
	// fmt.Fprintf(w, "<html>\n")  
	// fmt.Fprintf(w, "<head>\n")  
	// fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	// fmt.Fprintf(w, "</head>\n")  
	// fmt.Fprintf(w, "<body>\n")  
	// fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>date</th><th>weekNum</th><th>weekday</th><th>time</th><th>targetTime</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalGetup.DayOrderGetupInfo { 
		if item == nil {
		  continue
		}

		cellClass := ""
		if item.Day.Weekday == "一" {
			cellClass = "color-cell"
		} 

		fmt.Fprintf(w, "<tr>")
		fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, item.Day.Date)
		fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, item.Day.WeekNum)
		fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, item.Day.Weekday)
		fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, item.RawInfo) 
		fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, item.TargetTime) 
		if item.IsFinish {
			fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, "&#x2705;")
		} else {
			fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, "&#x274C;")
		}
		
		fmt.Fprintf(w, "</tr>\n")
	}  
  
	// 构造HTML表格的结尾  
	fmt.Fprintf(w, "</table>\n")  
	fmt.Fprintf(w, "</body>\n")  
	fmt.Fprintf(w, "</html>\n")  	
}




func WeekGetupHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>weekNum</th><th>actualFinishDays</th><th>targetFinishDays</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalGetup.WeekOrderGetupInfo { 
		if item == nil {
			continue 
		}  
		
		fmt.Fprintf(w, "<tr>")  
		fmt.Fprintf(w, "<td>%s</td>", item.Week.WeekNum)
		fmt.Fprintf(w, "<td>%d</td>", item.ActualFinishDays)  
		fmt.Fprintf(w, "<td>%d</td>", item.TargetFinishDays) 
		if item.IsFinish {
			fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
			fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		}  
		fmt.Fprintf(w, "</tr>\n")  	
	}  
  
	// 构造HTML表格的结尾  
	fmt.Fprintf(w, "</table>\n")  
	fmt.Fprintf(w, "</body>\n")  
	fmt.Fprintf(w, "</html>\n")  
}




func MonthGetupHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>monthNum</th><th>actualFinishDays</th><th>targetFinishDays</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalGetup.MonthOrderGetupInfo { 
		if item == nil {
			continue 
		} 

		fmt.Fprintf(w, "<tr>")  
		fmt.Fprintf(w, "<td>%s</td>", item.Month.MonthNum)
		fmt.Fprintf(w, "<td>%d</td>", item.ActualFinishDays)  
		fmt.Fprintf(w, "<td>%d</td>", item.TargetFinishDays) 
		if item.IsFinish {
			fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
			fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		} 
		fmt.Fprintf(w, "</tr>\n")  

		// for content, conGetupTime := range item.GetupTimeOfDifferentContent {
		// 	fmt.Fprintf(w, "<tr>")  
		// 	fmt.Fprintf(w, "<td>%d</td>", item.MonthNum)
		// 	fmt.Fprintf(w, "<td>%s</td>", item.GetupTime)  
		// 	fmt.Fprintf(w, "<td>%s</td>", content)  
		// 	fmt.Fprintf(w, "<td>%s</td>", conGetupTime)  
		// 	fmt.Fprintf(w, "</tr>\n")  
		// }		
	}  
  
	// 构造HTML表格的结尾  
	fmt.Fprintf(w, "</table>\n")  
	fmt.Fprintf(w, "</body>\n")  
	fmt.Fprintf(w, "</html>\n")  
}


func YearGetupHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>yearNum</th><th>actualFinishDays</th><th>targetFinishDays</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalGetup.YearOrderGetupInfo { 
		if item == nil {
			continue
		} 
		fmt.Fprintf(w, "<tr>")  
		fmt.Fprintf(w, "<td>%s</td>", item.Year.YearNum)
		fmt.Fprintf(w, "<td>%d</td>", item.ActualFinishDays)  
		fmt.Fprintf(w, "<td>%d</td>", item.TargetFinishDays) 
		if item.IsFinish {
			fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
			fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		} 
		fmt.Fprintf(w, "</tr>\n")  

		// for content, conGetupTime := range item.GetupTimeOfDifferentContent {
		// 	fmt.Fprintf(w, "<tr>")  
		// 	fmt.Fprintf(w, "<td>%d</td>", item.YearNum)
		// 	fmt.Fprintf(w, "<td>%s</td>", item.YearGetupTime)  
		// 	fmt.Fprintf(w, "<td>%s</td>", content)  
		// 	fmt.Fprintf(w, "<td>%s</td>", conGetupTime)  
		// 	fmt.Fprintf(w, "</tr>\n")  
		// }		
	}  
  
	// 构造HTML表格的结尾  
	fmt.Fprintf(w, "</table>\n")  
	fmt.Fprintf(w, "</body>\n")  
	fmt.Fprintf(w, "</html>\n")  
}