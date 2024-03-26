package route

import (
	"fmt"
	"net/http"

	hData "ningan.com/habit-tracking/pkg/data"
)


func DaySkinCareHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>date</th><th>weekNum</th><th>weekday</th><th>raw</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalSkinCare.DayOrderInfo { 
		if item == nil {
		  continue
		}
		fmt.Fprintf(w, "<tr>")
		fmt.Fprintf(w, "<td>%s</td>", item.Date)
		fmt.Fprintf(w, "<td>%s</td>", item.WeekNum)
		fmt.Fprintf(w, "<td>%s</td>", item.Weekday)
		fmt.Fprintf(w, "<td>%s</td>", item.RawInfo) 
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




func WeekSkinCareHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>weekNum</th><th>actualFinishDays</th><th>targetFinishDays</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalSkinCare.WeekOrderInfo { 
		if item == nil {
			continue 
		}  
		
		fmt.Fprintf(w, "<tr>")  
		fmt.Fprintf(w, "<td>%s</td>", item.WeekNum)
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




func MonthSkinCareHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>monthNum</th><th>actualFinishDays</th><th>targetFinishDays</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalSkinCare.MonthOrderInfo { 
		if item == nil {
			continue 
		} 

		fmt.Fprintf(w, "<tr>")  
		fmt.Fprintf(w, "<td>%s</td>", item.MonthNum)
		fmt.Fprintf(w, "<td>%d</td>", item.ActualFinishDays)  
		fmt.Fprintf(w, "<td>%d</td>", item.TargetFinishDays) 
		if item.IsFinish {
			fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
			fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		} 
		fmt.Fprintf(w, "</tr>\n")  

		// for content, conSkinCareTime := range item.SkinCareTimeOfDifferentContent {
		// 	fmt.Fprintf(w, "<tr>")  
		// 	fmt.Fprintf(w, "<td>%d</td>", item.MonthNum)
		// 	fmt.Fprintf(w, "<td>%s</td>", item.SkinCareTime)  
		// 	fmt.Fprintf(w, "<td>%s</td>", content)  
		// 	fmt.Fprintf(w, "<td>%s</td>", conSkinCareTime)  
		// 	fmt.Fprintf(w, "</tr>\n")  
		// }		
	}  
  
	// 构造HTML表格的结尾  
	fmt.Fprintf(w, "</table>\n")  
	fmt.Fprintf(w, "</body>\n")  
	fmt.Fprintf(w, "</html>\n")  
}


func YearSkinCareHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>yearNum</th><th>actualFinishDays</th><th>targetFinishDays</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalSkinCare.YearOrderInfo { 
		if item == nil {
			continue
		} 
		fmt.Fprintf(w, "<tr>")  
		fmt.Fprintf(w, "<td>%s</td>", item.YearNum)
		fmt.Fprintf(w, "<td>%d</td>", item.ActualFinishDays)  
		fmt.Fprintf(w, "<td>%d</td>", item.TargetFinishDays) 
		if item.IsFinish {
			fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
			fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		} 
		fmt.Fprintf(w, "</tr>\n")  

		// for content, conSkinCareTime := range item.SkinCareTimeOfDifferentContent {
		// 	fmt.Fprintf(w, "<tr>")  
		// 	fmt.Fprintf(w, "<td>%d</td>", item.YearNum)
		// 	fmt.Fprintf(w, "<td>%s</td>", item.YearSkinCareTime)  
		// 	fmt.Fprintf(w, "<td>%s</td>", content)  
		// 	fmt.Fprintf(w, "<td>%s</td>", conSkinCareTime)  
		// 	fmt.Fprintf(w, "</tr>\n")  
		// }		
	}  
  
	// 构造HTML表格的结尾  
	fmt.Fprintf(w, "</table>\n")  
	fmt.Fprintf(w, "</body>\n")  
	fmt.Fprintf(w, "</html>\n")  
}