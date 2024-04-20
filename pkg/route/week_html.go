package route

import (
	"fmt"
	"net/http"

	hData "ningan.com/habit-tracking/pkg/data"
)


func WeekHtmlTable(w http.ResponseWriter) {
	fmt.Fprintln(w, GlobalTable)  

	// // 构造HTML表格的开头  
	// fmt.Fprintf(w, "<html>\n")  
	// fmt.Fprintf(w, "<head>\n")  
	// fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	// fmt.Fprintf(w, "</head>\n")  
	// fmt.Fprintf(w, "<body>\n")  
	// fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th class='%s'>weekNum</th>","fixed-header")  
	fmt.Fprintf(w, "<th class='%s'>getup</th><th class='%s'>target</th><th class='%s'>finish</th>","fixed-header2","fixed-header2","fixed-header2")  
	fmt.Fprintf(w, "<th class='%s'>sleep</th><th class='%s'>target</th><th class='%s'>finish</th>","fixed-header3","fixed-header3","fixed-header3")  
	fmt.Fprintf(w, "<th class='%s'>reading</th><th class='%s'>target</th><th class='%s'>extra</th><th class='%s'>readingContentContent</th><th class='%s'>finish</th>","fixed-header2","fixed-header2","fixed-header2","fixed-header2","fixed-header2")  
	fmt.Fprintf(w, "<th class='%s'>piano</th><th class='%s'>target</th><th class='%s'>extra</th><th class='%s'>pianoContentContent</th><th class='%s'>finish</th>","fixed-header3","fixed-header3","fixed-header3","fixed-header3","fixed-header3")  
	fmt.Fprintf(w, "<th class='%s'>skincare</th><th class='%s'>facemask</th>","fixed-header2","fixed-header2")  
	fmt.Fprintf(w, "<th class='%s'>audio</th><th class='%s'>target</th><th class='%s'>finish</th></tr>\n", "fixed-header3","fixed-header3","fixed-header3")  
  
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

		sItem := hData.GlobalSleep.WeekSleepInfo[item.Week.WeekNum]
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

		rItem := hData.GlobalReading.WeekReadingInfo[item.Week.WeekNum]
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
		
		pItem := hData.GlobalPiano.WeekPianoInfo[item.Week.WeekNum]
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

		scItem :=  hData.GlobalSkinCare.WeekInfo[item.Week.WeekNum]
		if scItem == nil {
		  continue
		}
		if scItem.IsFinish {
		  		fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
		  		fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		}

		fItem :=  hData.GlobalFaceMask.WeekInfo[item.Week.WeekNum]
		if fItem == nil {
		  continue
		}
		if fItem.IsFinish {
		  		fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		} else {
		  		fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		}
		

		aItem :=  hData.GlobalAudiobook.WeekInfo[item.Week.WeekNum]
		if aItem == nil {
		  continue
		}
		fmt.Fprintf(w, "<td>%d</td>", aItem.FinishBooks)
		fmt.Fprintf(w, "<td>%d</td>", aItem.TargetFinishBooks)
		if aItem.IsFinish {
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
