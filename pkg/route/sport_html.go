package route

import (
	"fmt"
	"net/http"

	hData "ningan.com/habit-tracking/pkg/data"
)

func DaySportHtmlTable(w http.ResponseWriter) {
	// 开始写入HTML
	fmt.Fprintln(w, GlobalTable)

	// // 构造HTML表格的开头
	// fmt.Fprintf(w, "<html>\n")
	// fmt.Fprintf(w, "<head>\n")
	// fmt.Fprintf(w, "<title>MyStruct Table</title>\n")
	// fmt.Fprintf(w, "</head>\n")
	// fmt.Fprintf(w, "<body>\n")
	// fmt.Fprintf(w, "<table border='1'>\n")
	fmt.Fprintf(w, "<tr><th>date</th><th>weekNum</th><th>weekday</th><th>raw</th><th>daySportTimes</th><th>content</th><th>finish</th></tr>\n")

	// 遍历数据并构造表格的行
	for _, item := range hData.GlobalSport.DayOrderInfo {
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
		fmt.Fprintf(w, "<td class='%s'>%d</td>", cellClass, item.SportTimes)
		fmt.Fprintf(w, "<td class='%s'>%s</td>", cellClass, item.SportTimesOfDifferentContentStr)
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

func WeekSportHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头
	fmt.Fprintf(w, "<html>\n")
	fmt.Fprintf(w, "<head>\n")
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")
	fmt.Fprintf(w, "</head>\n")
	fmt.Fprintf(w, "<body>\n")
	fmt.Fprintf(w, "<table border='1'>\n")
	fmt.Fprintf(w, "<tr><th>weekNum</th><th>sportTimes</th><th>FinishDays</th><th>targetFinishDays</th><th>content</th><th>finish</th></tr>\n")

	// 遍历数据并构造表格的行
	for _, item := range hData.GlobalSport.WeekOrderInfo {
		if item == nil {
			continue
		}

		fmt.Fprintf(w, "<tr>")
		fmt.Fprintf(w, "<td>%s</td>", item.Week.WeekNum)
		fmt.Fprintf(w, "<td>%d</td>", item.SportTimes)
		fmt.Fprintf(w, "<td>%d</td>", item.ActualFinishDays)
		fmt.Fprintf(w, "<td>%d</td>", item.TargetFinishDays)
		fmt.Fprintf(w, "<td>%s</td>", item.SportTimesOfDifferentContentStr)
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

func MonthSportHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头
	fmt.Fprintf(w, "<html>\n")
	fmt.Fprintf(w, "<head>\n")
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")
	fmt.Fprintf(w, "</head>\n")
	fmt.Fprintf(w, "<body>\n")
	fmt.Fprintf(w, "<table border='1'>\n")
	fmt.Fprintf(w, "<tr><th>monthNum</th><th>sportTimes</th><th>FinishDays</th><th>targetFinishDays</th><th>content</th><th>finish</th></tr>\n")

	// 遍历数据并构造表格的行
	for _, item := range hData.GlobalSport.MonthOrderInfo {
		if item == nil {
			continue
		}

		fmt.Fprintf(w, "<tr>")
		fmt.Fprintf(w, "<td>%s</td>", item.Month.MonthNum)
		fmt.Fprintf(w, "<td>%d</td>", item.SportTimes)
		fmt.Fprintf(w, "<td>%d</td>", item.ActualFinishDays)
		fmt.Fprintf(w, "<td>%d</td>", item.TargetFinishDays)
		fmt.Fprintf(w, "<td>%s</td>", item.SportTimesOfDifferentContentStr)
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

func YearSportHtmlTable(w http.ResponseWriter) {
	// 构造HTML表格的开头
	fmt.Fprintf(w, "<html>\n")
	fmt.Fprintf(w, "<head>\n")
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")
	fmt.Fprintf(w, "</head>\n")
	fmt.Fprintf(w, "<body>\n")
	fmt.Fprintf(w, "<table border='1'>\n")
	fmt.Fprintf(w, "<tr><th>yearNum</th><th>sportTimes</th><th>FinishDays</th><th>targetFinishDays</th><th>content</th><th>finish</th></tr>\n")

	// 遍历数据并构造表格的行
	for _, item := range hData.GlobalSport.YearOrderInfo {
		if item == nil {
			continue
		}
		fmt.Fprintf(w, "<tr>")
		fmt.Fprintf(w, "<td>%s</td>", item.Year.YearNum)
		fmt.Fprintf(w, "<td>%d</td>", item.SportTimes)
		fmt.Fprintf(w, "<td>%d</td>", item.ActualFinishDays)
		fmt.Fprintf(w, "<td>%d</td>", item.TargetFinishDays)
		fmt.Fprintf(w, "<td>%s</td>", item.SportTimesOfDifferentContentStr)
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
