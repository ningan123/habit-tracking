package route

import (
	"fmt"
	"net/http"

	hData "ningan.com/habit-tracking/pkg/data"
)


func DayAudiobookHtmlTable(w http.ResponseWriter) {
	// 开始写入HTML  
	fmt.Fprintln(w, GlobalTable)  

	// // 构造HTML表格的开头  
	// fmt.Fprintf(w, "<html>\n")  
	// fmt.Fprintf(w, "<head>\n")  
	// fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	// fmt.Fprintf(w, "</head>\n")  
	// fmt.Fprintf(w, "<body>\n")  
	// fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>date</th><th>weekNum</th><th>weekday</th><th>raw</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalAudiobook.DayOrderInfo { 
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




func WeekAudiobookHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>weekNum</th><th>FinishBooks</th><th>targetFinishBooks</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalAudiobook.WeekOrderInfo { 
		if item == nil {
			continue 
		}  
		
		fmt.Fprintf(w, "<tr>")  
		fmt.Fprintf(w, "<td>%s</td>", item.Week.WeekNum)
		fmt.Fprintf(w, "<td>%d</td>", item.FinishBooks)  
		fmt.Fprintf(w, "<td>%d</td>", item.TargetFinishBooks) 
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




func MonthAudiobookHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>monthNum</th><th>FinishBooks</th><th>targetFinishBooks</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalAudiobook.MonthOrderInfo { 
		if item == nil {
			continue 
		} 

		fmt.Fprintf(w, "<tr>")  
		fmt.Fprintf(w, "<td>%s</td>", item.Month.MonthNum)
		fmt.Fprintf(w, "<td>%d</td>", item.FinishBooks)  
		fmt.Fprintf(w, "<td>%d</td>", item.TargetFinishBooks) 
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


func YearAudiobookHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>yearNum</th><th>FinishBooks</th><th>targetFinishBooks</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalAudiobook.YearOrderInfo { 
		if item == nil {
			continue
		} 
		fmt.Fprintf(w, "<tr>")  
		fmt.Fprintf(w, "<td>%s</td>", item.Year.YearNum)
		fmt.Fprintf(w, "<td>%d</td>", item.FinishBooks)  
		fmt.Fprintf(w, "<td>%d</td>", item.TargetFinishBooks) 
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