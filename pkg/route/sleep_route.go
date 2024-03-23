package route

import (
	"fmt"
	"net/http"
)

// 定义处理函数，用于处理/sleep路径的请求
func SleepHandler(w http.ResponseWriter, r *http.Request) {  
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
		    <li><a href="/sleep/day">Link to day sleep</a></li>  
        <li><a href="/sleep/week">Link to week sleep</a></li> 
				<li><a href="/sleep/month">Link to month sleep</a></li>  
				<li><a href="/sleep/year">Link to year sleep</a></li>  
				<li><a href="/sleep/all">Link to all sleep info</a></li>
    </ul>  
</body>  
</html>`)  
	
} 

func AllSleepHandler(w http.ResponseWriter, r *http.Request) {
  
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	DaySleepHtmlTable(w)
	WeekSleepHtmlTable(w)
	MonthSleepHtmlTable(w)
	YearSleepHtmlTable(w)
}


// 定义处理函数，用于处理/sleep/day路径的请求  
func DaySleepHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	DaySleepHtmlTable(w)
}


// 定义处理函数，用于处理/sleep/week路径的请求  
func WeekSleepHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	WeekSleepHtmlTable(w)
}


// 定义处理函数，用于处理/sleep/month路径的请求
func MonthSleepHandler (w http.ResponseWriter, r *http.Request) {
  
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	MonthSleepHtmlTable(w)
}


// 定义处理函数，用于处理/sleep/year路径的请求
func YearSleepHandler (w http.ResponseWriter, r *http.Request) {
  
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	YearSleepHtmlTable(w)
}