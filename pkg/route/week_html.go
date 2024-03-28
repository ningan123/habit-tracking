package route

import (
	"fmt"
	"net/http"

	hData "ningan.com/habit-tracking/pkg/data"
)


func WeekHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>weekNum</th><th>getupDays</th><th>getupTargetDays</th><th>finish</th><th>sleepDays</th><th>sleepTargetDays</th><th>finish</th><th>weekReadingTime</th><th>targetReadingTime</th><th>extraReadingTime</th><th>content</th><th>finish</th><th>weekPianoTime</th><th>targetPianoTime</th><th>extraPianoTime</th><th>content</th><th>finish</th><th>skincare</th><th>facemask</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalGetup.WeekOrderGetupInfo { 
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

		sItem := hData.GlobalSleep.WeekSleepInfo[item.WeekNum]
		if sItem == nil {
			continue
		}
		fmt.Fprintf(w, "<td>%d</td>", sItem.ActualFinishDays)
		fmt.Fprintf(w, "<td>%d</td>", sItem.TargetFinishDays)
		if sItem.IsFinish {
			fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
			fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		}

		rItem := hData.GlobalReading.WeekReadingInfo[item.WeekNum]
		if rItem == nil {
			continue
		}
		fmt.Fprintf(w, "<td>%s</td>", rItem.ReadingTime)  
		fmt.Fprintf(w, "<td>%s</td>", rItem.TargetReadingTime) 
		fmt.Fprintf(w, "<td>%s</td>", rItem.ExtraReadingTime) 
		fmt.Fprintf(w, "<td>%s</td>", rItem.ReadingTimeOfDifferentContentStr) 		
		if rItem.IsFinish {
			fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
			fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		} 
		
		pItem := hData.GlobalPiano.WeekPianoInfo[item.WeekNum]
		if pItem == nil {
			continue
		}
		fmt.Fprintf(w, "<td>%s</td>", pItem.PianoTime)
		fmt.Fprintf(w, "<td>%s</td>", pItem.TargetPianoTime)
		fmt.Fprintf(w, "<td>%s</td>", pItem.ExtraPianoTime)
		fmt.Fprintf(w, "<td>%s</td>", pItem.PianoTimeOfDifferentContentStr)
		if pItem.IsFinish {
		  fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
		  fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		}

		scItem :=  hData.GlobalSkinCare.WeekInfo[item.WeekNum]
		if scItem == nil {
		  continue
		}
		if scItem.IsFinish {
		  fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
		  fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		}

		fItem :=  hData.GlobalFaceMask.WeekInfo[item.WeekNum]
		if scItem == nil {
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
