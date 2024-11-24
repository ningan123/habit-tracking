package route

import (
	"fmt"
	"net/http"

	hData "ningan.com/habit-tracking/pkg/data"
)

func MonthHtmlTable(w http.ResponseWriter) {
	fmt.Fprintln(w, GlobalTable)

	// // 构造HTML表格的开头
	// fmt.Fprintf(w, "<html>\n")
	// fmt.Fprintf(w, "<head>\n")
	// fmt.Fprintf(w, "<title>MyStruct Table</title>\n")
	// fmt.Fprintf(w, "</head>\n")
	// fmt.Fprintf(w, "<body>\n")
	// fmt.Fprintf(w, "<table border='1'>\n")
	fmt.Fprintf(w, "<tr><th class='%s'>monthNum</th>", "fixed-header")
	// fmt.Fprintf(w, "<th class='%s'>getup</th><th class='%s'>target</th><th class='%s'>finish</th>", "fixed-header2", "fixed-header2", "fixed-header2")
	// fmt.Fprintf(w, "<th class='%s'>sleep</th><th class='%s'>target</th><th class='%s'>finish</th>", "fixed-header3", "fixed-header3", "fixed-header3")
	// fmt.Fprintf(w, "<th class='%s'>reading</th><th class='%s'>target</th><th class='%s'>extra</th><th class='%s'>readingContentContent</th><th class='%s'>finish</th>", "fixed-header2", "fixed-header2", "fixed-header2", "fixed-header2", "fixed-header2")
	fmt.Fprintf(w, "<th class='%s'>piano</th><th class='%s'>target</th><th class='%s'>extra</th><th class='%s'>pianoContentContent</th><th class='%s'>finish</th>", "fixed-header3", "fixed-header3", "fixed-header3", "fixed-header3", "fixed-header3")
	// fmt.Fprintf(w, "<th class='%s'>skincare</th><th class='%s'>target</th><th class='%s'>finish</th>", "fixed-header2", "fixed-header2", "fixed-header2")
	// fmt.Fprintf(w, "<th class='%s'>facemask</th><th class='%s'>target</th><th class='%s'>finish</th>", "fixed-header3", "fixed-header3", "fixed-header3")
	// fmt.Fprintf(w, "<th class='%s'>audio</th><th class='%s'>target</th><th class='%s'>finish</th>", "fixed-header2", "fixed-header2", "fixed-header2")
	// fmt.Fprintf(w, "<th class='%s'>sport</th><th class='%s'>actual</th><th class='%s'>target</th><th class='%s'>sportContentContent</th><th class='%s'>finish</th></tr>\n", "fixed-header3", "fixed-header3", "fixed-header3", "fixed-header3", "fixed-header3")

	// 遍历数据并构造表格的行
	for _, item := range hData.GlobalGetup.MonthOrderGetupInfo {
		if item == nil {
			continue
		}

		fmt.Fprintf(w, "<tr>")
		fmt.Fprintf(w, "<td>%s</td>", item.Month.MonthNum)
		// fmt.Fprintf(w, "<td>%d</td>", item.ActualFinishDays)
		// fmt.Fprintf(w, "<td>%d</td>", item.TargetFinishDays)
		// if item.IsFinish {
		// 	fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		// } else {
		// 	fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		// }

		// sItem := hData.GlobalSleep.MonthSleepInfo[item.Month.MonthNum]
		// if sItem == nil {
		// 	continue
		// }
		// fmt.Fprintf(w, "<td>%d</td>", sItem.ActualFinishDays)
		// fmt.Fprintf(w, "<td>%d</td>", sItem.TargetFinishDays)
		// if sItem.IsFinish {
		// 	fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		// } else {
		// 	fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		// }

		// rItem := hData.GlobalReading.MonthReadingInfo[item.Month.MonthNum]
		// if rItem == nil {
		// 	continue
		// }
		// fmt.Fprintf(w, "<td>%s</td>", rItem.ReadingTime)
		// fmt.Fprintf(w, "<td>%s</td>", rItem.TargetReadingTime)
		// fmt.Fprintf(w, "<td>%s</td>", rItem.ExtraReadingTime)
		// fmt.Fprintf(w, "<td>%s</td>", rItem.ReadingTimeOfDifferentContentStr)
		// if rItem.IsFinish {
		// 	fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		// } else {
		// 	fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		// }

		pItem := hData.GlobalPiano.MonthPianoInfo[item.Month.MonthNum]
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

		// scItem := hData.GlobalSkinCare.MonthInfo[item.Month.MonthNum]
		// if scItem == nil {
		// 	continue
		// }
		// fmt.Fprintf(w, "<td>%d</td>", scItem.ActualFinishDays)
		// fmt.Fprintf(w, "<td>%d</td>", scItem.TargetFinishDays)
		// if scItem.IsFinish {
		// 	fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		// } else {
		// 	fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		// }

		// fItem := hData.GlobalFaceMask.MonthInfo[item.Month.MonthNum]
		// if fItem == nil {
		// 	continue
		// }
		// fmt.Fprintf(w, "<td>%d</td>", fItem.ActualFinishDays)
		// fmt.Fprintf(w, "<td>%d</td>", fItem.TargetFinishDays)
		// if fItem.IsFinish {
		// 	fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		// } else {
		// 	fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		// }

		// aItem := hData.GlobalAudiobook.MonthInfo[item.Month.MonthNum]
		// if aItem == nil {
		// 	continue
		// }
		// fmt.Fprintf(w, "<td>%d</td>", aItem.FinishBooks)
		// fmt.Fprintf(w, "<td>%d</td>", aItem.TargetFinishBooks)
		// if aItem.IsFinish {
		// 	fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		// } else {
		// 	fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		// }

		// spItem := hData.GlobalSport.MonthInfo[item.Month.MonthNum]
		// if spItem == nil {
		// 	continue
		// }
		// fmt.Fprintf(w, "<td>%d</td>", spItem.SportTimes)
		// fmt.Fprintf(w, "<td>%d</td>", spItem.ActualFinishDays)
		// fmt.Fprintf(w, "<td>%d</td>", spItem.TargetFinishDays)
		// fmt.Fprintf(w, "<td>%s</td>", spItem.SportTimesOfDifferentContentStr)
		// if spItem.IsFinish {
		// 	fmt.Fprintf(w, "<td>%s</td>", "&#x2705;")
		// } else {
		// 	fmt.Fprintf(w, "<td>%s</td>", "&#x274C;")
		// }

		fmt.Fprintf(w, "</tr>\n")
	}

	// 构造HTML表格的结尾
	fmt.Fprintf(w, "</table>\n")
	fmt.Fprintf(w, "</body>\n")
	fmt.Fprintf(w, "</html>\n")
}
