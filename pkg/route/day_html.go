package route

import (
	"fmt"
	"net/http"

	hData "ningan.com/habit-tracking/pkg/data"
)

func DayHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>date</th><th>weekNum</th><th>weekday</th><th>getup</th><th>targetTime</th><th>finish</th><th>sleep</th><th>targetTime</th><th>finish</th><th>ReadingTime</th><th>targetReadingTime</th><th>content</th><th>finish</th><th>PianoTime</th><th>targetPianoTime</th><th>content</th><th>finish</th><th>skincare</th><th>facemask</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalGetup.DayOrderGetupInfo { 
		if item == nil {
		  continue
		}
		fmt.Fprintf(w, "<tr>")
		fmt.Fprintf(w, "<td>%s</td>", item.Date)
		fmt.Fprintf(w, "<td>%s</td>", item.WeekNum)
		fmt.Fprintf(w, "<td>%s</td>", item.Weekday)
		fmt.Fprintf(w, "<td>%s</td>", item.RawInfo) 
		fmt.Fprintf(w, "<td>%s</td>", item.TargetTime) 
		if item.IsFinish {
			fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
			fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		}

		sItem := hData.GlobalSleep.DaySleepInfo[item.Date]
		if sItem == nil {
		  continue
		}
		fmt.Fprintf(w, "<td>%s</td>", sItem.RawInfo)
		fmt.Fprintf(w, "<td>%s</td>", sItem.TargetTime)
		if sItem.IsFinish {
		  fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
		  fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		}

		rItem := hData.GlobalReading.DayReadingInfo[item.Date]
		if rItem == nil {
		  continue
		}
		fmt.Fprintf(w, "<td>%s</td>", rItem.ReadingTime) 
		fmt.Fprintf(w, "<td>%s</td>", rItem.TargetReadingTime) 
		fmt.Fprintf(w, "<td>%s</td>", rItem.ReadingTimeOfDifferentContentStr)
		if rItem.IsFinish {
			fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
			fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		}

		pItem := hData.GlobalPiano.DayPianoInfo[item.Date]
		if pItem == nil {
		  continue
		}
		fmt.Fprintf(w, "<td>%s</td>", pItem.PianoTime)
		fmt.Fprintf(w, "<td>%s</td>", pItem.TargetPianoTime)
		fmt.Fprintf(w, "<td>%s</td>", pItem.PianoTimeOfDifferentContentStr)
		if pItem.IsFinish {
		  fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
		  fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		}

		scItem := hData.GlobalSkinCare.DayInfo[item.Date]
		if scItem == nil {
		  continue
		}
		if scItem.IsFinish {
		  fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
		  fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		}
		
		fItem := hData.GlobalFaceMask.DayInfo[item.Date]
		if fItem == nil {
		  continue
		}
		if fItem.IsFinish {
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
