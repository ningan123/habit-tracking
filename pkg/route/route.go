package route

import (
	"fmt"
	"net/http"

	hData "ningan.com/habit-tracking/pkg/data"
)

// 定义一个处理函数，用于返回包含两个链接的HTML页面
func RootHandler(w http.ResponseWriter, r *http.Request) {  
	fmt.Fprintf(w, `<!DOCTYPE html>  
<html lang="en">  
<head>  
    <meta charset="UTF-8">  
    <title>My Links</title>  
</head>  
<body>  
    <h1>Welcome to My Page</h1>  
    <p>Here are some links:</p>  
    <ul>  
		    <li><a href="/reading">Link to reading</a></li>  
        <li><a href="/a">Link to A</a></li>  
        <li><a href="/b">Link to B</a></li>  
    </ul>  
</body>  
</html>`)  
}  



// 定义处理函数，用于处理/a路径的请求  
func AHandler(w http.ResponseWriter, r *http.Request) {  
	fmt.Fprint(w, "You are at path /a")  
}  
  
// 定义处理函数，用于处理/b路径的请求  
func BHandler(w http.ResponseWriter, r *http.Request) {  
	fmt.Fprint(w, "You are at path /b")  
}  


// 定义处理函数，用于处理/reading路径的请求  
func ReadingHandler(w http.ResponseWriter, r *http.Request) {  
	fmt.Fprintf(w, `<!DOCTYPE html>  
<html lang="en">  
<head>  
    <meta charset="UTF-8">  
    <title>My Links</title>  
</head>  
<body>  
    <h1>Welcome to My Page</h1>  
    <p>Here are some links:</p>  
    <ul>  
		    <li><a href="/reading/day">Link to day reading</a></li>  
        <li><a href="/reading/week">Link to week reading</a></li> 
				<li><a href="/reading/month">Link to month reading</a></li>  
    </ul>  
</body>  
</html>`)  
	
} 

// 定义处理函数，用于处理/reading/day路径的请求  
func DayReadingHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>Date</th><th>DayReadingTime</th><th>content</th><th>contentReadingTime</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalReading.DayReadingInfo {  
		for content, conReadingTime := range item.DayReadingTimeOfDifferentContent {
			fmt.Fprintf(w, "<tr>")  
			fmt.Fprintf(w, "<td>%s</td>", item.DayDate)
			fmt.Fprintf(w, "<td>%s</td>", item.DayReadingTime)  
			fmt.Fprintf(w, "<td>%s</td>", content)  
			fmt.Fprintf(w, "<td>%s</td>", conReadingTime)  
			fmt.Fprintf(w, "</tr>\n")  
		}		
	}  
  
	// 构造HTML表格的结尾  
	fmt.Fprintf(w, "</table>\n")  
	fmt.Fprintf(w, "</body>\n")  
	fmt.Fprintf(w, "</html>\n")  
}


// 定义处理函数，用于处理/reading/week路径的请求  
func WeekReadingHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>WeekNum</th><th>WeekReadingTime</th><th>content</th><th>contentReadingTime</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalReading.WeekReadingInfo {  
		for content, conReadingTime := range item.WeekReadingTimeOfDifferentContent {
			fmt.Fprintf(w, "<tr>")  
			fmt.Fprintf(w, "<td>%d</td>", item.WeekNum)
			fmt.Fprintf(w, "<td>%s</td>", item.WeekReadingTime)  
			fmt.Fprintf(w, "<td>%s</td>", content)  
			fmt.Fprintf(w, "<td>%s</td>", conReadingTime)  
			fmt.Fprintf(w, "</tr>\n")  
		}		
	}  
  
	// 构造HTML表格的结尾  
	fmt.Fprintf(w, "</table>\n")  
	fmt.Fprintf(w, "</body>\n")  
	fmt.Fprintf(w, "</html>\n")  
}



// 定义处理函数，用于处理/reading/month路径的请求  
func MonthReadingHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	// 构造HTML表格的开头  
	fmt.Fprintf(w, "<html>\n")  
	fmt.Fprintf(w, "<head>\n")  
	fmt.Fprintf(w, "<title>MyStruct Table</title>\n")  
	fmt.Fprintf(w, "</head>\n")  
	fmt.Fprintf(w, "<body>\n")  
	fmt.Fprintf(w, "<table border='1'>\n")  
	fmt.Fprintf(w, "<tr><th>MonthNum</th><th>MonthReadingTime</th><th>content</th><th>contentReadingTime</th></tr>\n")  
  
	// 遍历数据并构造表格的行  
	for _, item := range hData.GlobalReading.MonthReadingInfo {  
		for content, conReadingTime := range item.MonthReadingTimeOfDifferentContent {
			fmt.Fprintf(w, "<tr>")  
			fmt.Fprintf(w, "<td>%d</td>", item.MonthNum)
			fmt.Fprintf(w, "<td>%s</td>", item.MonthReadingTime)  
			fmt.Fprintf(w, "<td>%s</td>", content)  
			fmt.Fprintf(w, "<td>%s</td>", conReadingTime)  
			fmt.Fprintf(w, "</tr>\n")  
		}		
	}  
  
	// 构造HTML表格的结尾  
	fmt.Fprintf(w, "</table>\n")  
	fmt.Fprintf(w, "</body>\n")  
	fmt.Fprintf(w, "</html>\n")  
}

