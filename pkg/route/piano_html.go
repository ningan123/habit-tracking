package route

import (
	"fmt"
	"net/http"

	hData "ningan.com/habit-tracking/pkg/data"
)


func DayPianoHtmlTable(w http.ResponseWriter) {
	fmt.Fprintln(w, GlobalTable)  

	// // 构造HTML表格的开头  
	// fmt.Fprintf(w, "<html>\n")  
	// fmt.Fprintf(w, "<head>\n")  
	// fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	// fmt.Fprintf(w, "</head>\n")  
	// fmt.Fprintf(w, "<body>\n")  
	// fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>date</th><th>weekNum</th><th>weekday</th><th>dayPianoTime</th><th>targetPianoTime</th><th>content</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalPiano.DayOrderPianoInfo { 
		if item == nil {
		  continue
		}

		cellClass := ""
		if item.Day.Weekday == "一" {
			cellClass = "piano-cell"
		} 


		fmt.Fprintf(w, "<tr>")
		fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, item.Day.Date)
		fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, item.Day.WeekNum)
		fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, item.Day.Weekday)
		fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, item.PianoTime) 
		fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, item.TargetPianoTime) 
		fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, item.PianoTimeOfDifferentContentStr)
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

func WeekPianoHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>weekNum</th><th>weekPianoTime</th><th>targetPianoTime</th><th>extraPianoTime</th><th>content</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalPiano.WeekOrderPianoInfo { 
		if item == nil {
			continue 
		}  
		
		fmt.Fprintf(w, "<tr>")  
		fmt.Fprintf(w, "<td>%s</td>", item.Week.WeekNum)
		fmt.Fprintf(w, "<td>%s</td>", item.PianoTime)  
		fmt.Fprintf(w, "<td>%s</td>", item.TargetPianoTime) 
		fmt.Fprintf(w, "<td>%s</td>", item.ExtraPianoTime) 
		fmt.Fprintf(w, "<td>%s</td>", item.PianoTimeOfDifferentContentStr) 		
		if item.IsFinish {
			fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
			fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		}  
		fmt.Fprintf(w, "</tr>\n")  
		

		// for content, conPianoTime := range item.PianoTimeOfDifferentContent {
		// 	fmt.Fprintf(w, "<tr>")  
		// 	fmt.Fprintf(w, "<td>%d</td>", item.WeekNum)
		// 	fmt.Fprintf(w, "<td>%s</td>", item.PianoTime)  
		// 	fmt.Fprintf(w, "<td>%s</td>", content)  
		// 	fmt.Fprintf(w, "<td>%s</td>", conPianoTime)  
		// 	fmt.Fprintf(w, "</tr>\n")  
		// }		
	}  
  
	// 构造HTML表格的结尾  
	fmt.Fprintf(w, "</table>\n")  
	fmt.Fprintf(w, "</body>\n")  
	fmt.Fprintf(w, "</html>\n")  
}


func MonthPianoHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>monthNum</th><th>monthPianoTime</th><th>targetPianoTime</th><th>content</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalPiano.MonthOrderPianoInfo { 
		if item == nil {
			continue 
		} 

		fmt.Fprintf(w, "<tr>")  
		fmt.Fprintf(w, "<td>%s</td>", item.Month.MonthNum)
		fmt.Fprintf(w, "<td>%s</td>", item.PianoTime)  
		fmt.Fprintf(w, "<td>%s</td>", item.TargetPianoTime) 
		fmt.Fprintf(w, "<td>%s</td>", item.PianoTimeOfDifferentContentStr)
		if item.IsFinish {
			fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
			fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		}  
		fmt.Fprintf(w, "</tr>\n")  

		// for content, conPianoTime := range item.PianoTimeOfDifferentContent {
		// 	fmt.Fprintf(w, "<tr>")  
		// 	fmt.Fprintf(w, "<td>%d</td>", item.MonthNum)
		// 	fmt.Fprintf(w, "<td>%s</td>", item.PianoTime)  
		// 	fmt.Fprintf(w, "<td>%s</td>", content)  
		// 	fmt.Fprintf(w, "<td>%s</td>", conPianoTime)  
		// 	fmt.Fprintf(w, "</tr>\n")  
		// }		
	}  
  
	// 构造HTML表格的结尾  
	fmt.Fprintf(w, "</table>\n")  
	fmt.Fprintf(w, "</body>\n")  
	fmt.Fprintf(w, "</html>\n")  
}


func YearPianoHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>yearNum</th><th>yearPianoTime</th><th>targetPianoTime</th><th>content</th><th>finish</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalPiano.YearOrderPianoInfo { 
		if item == nil {
			continue
		} 
		fmt.Fprintf(w, "<tr>")  
		fmt.Fprintf(w, "<td>%s</td>", item.Year.YearNum)
		fmt.Fprintf(w, "<td>%s</td>", item.PianoTime)  
		fmt.Fprintf(w, "<td>%s</td>", item.TargetPianoTime)  
		fmt.Fprintf(w, "<td>%s</td>", item.PianoTimeOfDifferentContentStr) 
		if item.IsFinish {
			fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
			fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		}   
		fmt.Fprintf(w, "</tr>\n")  

		// for content, conPianoTime := range item.PianoTimeOfDifferentContent {
		// 	fmt.Fprintf(w, "<tr>")  
		// 	fmt.Fprintf(w, "<td>%d</td>", item.YearNum)
		// 	fmt.Fprintf(w, "<td>%s</td>", item.YearPianoTime)  
		// 	fmt.Fprintf(w, "<td>%s</td>", content)  
		// 	fmt.Fprintf(w, "<td>%s</td>", conPianoTime)  
		// 	fmt.Fprintf(w, "</tr>\n")  
		// }		
	}  
  
	// 构造HTML表格的结尾  
	fmt.Fprintf(w, "</table>\n")  
	fmt.Fprintf(w, "</body>\n")  
	fmt.Fprintf(w, "</html>\n")  
}