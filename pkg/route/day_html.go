package route

import (
	"fmt"
	"net/http"

	hData "ningan.com/habit-tracking/pkg/data"
)

func DayHtmlTable(w http.ResponseWriter) {
	fmt.Fprintln(w, GlobalTable)

	// 构造HTML表格的开头
	fmt.Fprintf(w, "<html>\n")
	fmt.Fprintf(w, "<head>\n")
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")
	fmt.Fprintf(w, "</head>\n")
	fmt.Fprintf(w, "<body>\n")
	fmt.Fprintf(w, "<table border='1'>\n")
	fmt.Fprintf(w, "<tr><th class='%s'>dateDateDate</th><th class='%s'>weekNum</th><th class='%s'>weekday</th>", "fixed-header", "fixed-header", "fixed-header")
	// fmt.Fprintf(w, "<th class='%s'>getup</th><th class='%s'>target</th><th class='%s'>finish</th>", "fixed-header2", "fixed-header2", "fixed-header2")
	// fmt.Fprintf(w, "<th class='%s'>sleep</th><th class='%s'>target</th><th class='%s'>finish</th>", "fixed-header3", "fixed-header3", "fixed-header3")
	// fmt.Fprintf(w, "<th class='%s'>reading</th><th class='%s'>target</th><th class='%s'>readingcContentContent</th><th class='%s'>finish</th>", "fixed-header2", "fixed-header2", "fixed-header2", "fixed-header2")
	fmt.Fprintf(w, "<th class='%s'>piano</th><th class='%s'>target</th><th class='%s'>pianoContentContent</th><th class='%s'>finish</th>", "fixed-header3", "fixed-header3", "fixed-header3", "fixed-header3")
	// fmt.Fprintf(w, "<th class='%s'>skincare</th><th class='%s'>facemask</th>", "fixed-header2", "fixed-header2")
	// fmt.Fprintf(w, "<th class='%s'>audiobookContentContent</th><th class='%s'>finish</th>", "fixed-header3", "fixed-header3")
	// fmt.Fprintf(w, "<th class='%s'>sport</th><th class='%s'>sportContentContent</th><th class='%s'>finish</th></tr>\n", "fixed-header2", "fixed-header2", "fixed-header2")

	// 遍历数据并构造表格的行
	for _, item := range hData.GlobalPiano.DayOrderPianoInfo {
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
		// fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, item.RawInfo)
		// fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, item.TargetTime)
		// if item.IsFinish {
		// 	fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, "&#x2705;")
		// } else {
		// 	fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, "&#x274C;")
		// }

		// sItem := hData.GlobalSleep.DaySleepInfo[item.Day.Date]
		// if sItem == nil {
		// 	continue
		// }

		// fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, sItem.RawInfo)
		// fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, sItem.TargetTime)
		// if sItem.IsFinish {
		// 	fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, "&#x2705;")
		// } else {
		// 	fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, "&#x274C;")
		// }

		// rItem := hData.GlobalReading.DayReadingInfo[item.Day.Date]
		// if rItem == nil {
		// 	continue
		// }
		// fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, rItem.ReadingTime)
		// fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, rItem.TargetReadingTime)
		// fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, rItem.ReadingTimeOfDifferentContentStr)
		// if rItem.IsFinish {
		// 	fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, "&#x2705;")
		// } else {
		// 	fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, "&#x274C;")
		// }

		pItem := hData.GlobalPiano.DayPianoInfo[item.Day.Date]
		if pItem == nil {
			continue
		}
		fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, pItem.PianoTime)
		fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, pItem.TargetPianoTime)
		fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, pItem.PianoTimeOfDifferentContentStr)
		if pItem.IsFinish {
			fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, "&#x2705;")
		} else {
			fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, "&#x274C;")
		}

		// scItem := hData.GlobalSkinCare.DayInfo[item.Day.Date]
		// if scItem == nil {
		// 	continue
		// }
		// if scItem.IsFinish {
		// 	fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, "&#x2705;")
		// } else {
		// 	fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, "&#x274C;")
		// }

		// fItem := hData.GlobalFaceMask.DayInfo[item.Day.Date]
		// if fItem == nil {
		// 	continue
		// }
		// if fItem.IsFinish {
		// 	fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, "&#x2705;")
		// } else {
		// 	fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, "&#x274C;")
		// }

		// aItem := hData.GlobalAudiobook.DayInfo[item.Day.Date]
		// if aItem == nil {
		// 	continue
		// }
		// fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, aItem.RawInfo)
		// if aItem.IsFinish {
		// 	fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, "&#x2705;")
		// } else {
		// 	fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, "&#x274C;")
		// }

		// spItem := hData.GlobalSport.DayInfo[item.Day.Date]
		// if spItem == nil {
		// 	continue
		// }
		// fmt.Fprintf(w, "<td class='%s'>%d</td>", cellClass, spItem.SportTimes)
		// fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, spItem.SportTimesOfDifferentContentStr)
		// if spItem.IsFinish {
		// 	fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, "&#x2705;")
		// } else {
		// 	fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, "&#x274C;")
		// }

		fmt.Fprintf(w, "</tr>\n")
	}

	// 构造HTML表格的结尾
	fmt.Fprintf(w, "</table>\n")
	fmt.Fprintf(w, "</body>\n")
	fmt.Fprintf(w, "</html>\n")
}
