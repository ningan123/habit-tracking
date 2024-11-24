package route

import (
	"fmt"
	"net/http"
)

// 定义处理函数，用于处理/eyemask路径的请求
func EyeMaskHandler(w http.ResponseWriter, r *http.Request) {  
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
		    <li><a href="/eyemask/day">Link to day eyemask</a></li>  
        <li><a href="/eyemask/week">Link to week eyemask</a></li> 
				<li><a href="/eyemask/month">Link to month eyemask</a></li>  
				<li><a href="/eyemask/year">Link to year eyemask</a></li>  
				<li><a href="/eyemask/all">Link to all eyemask info</a></li>
    </ul>  
</body>  
</html>`)  
	
} 

func AllEyeMaskHandler(w http.ResponseWriter, r *http.Request) {
  
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	DayEyeMaskHtmlTable(w)
	WeekEyeMaskHtmlTable(w)
	MonthEyeMaskHtmlTable(w)
	YearEyeMaskHtmlTable(w)
}


// 定义处理函数，用于处理/eyemask/day路径的请求  
func DayEyeMaskHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	DayEyeMaskHtmlTable(w)
}


// 定义处理函数，用于处理/eyemask/week路径的请求  
func WeekEyeMaskHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	WeekEyeMaskHtmlTable(w)
}


// 定义处理函数，用于处理/eyemask/month路径的请求
func MonthEyeMaskHandler (w http.ResponseWriter, r *http.Request) {
  
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	MonthEyeMaskHtmlTable(w)
}


// 定义处理函数，用于处理/eyemask/year路径的请求
func YearEyeMaskHandler (w http.ResponseWriter, r *http.Request) {
  
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	YearEyeMaskHtmlTable(w)
}