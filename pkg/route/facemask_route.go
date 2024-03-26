package route

import (
	"fmt"
	"net/http"
)

// 定义处理函数，用于处理/facemask路径的请求
func FaceMaskHandler(w http.ResponseWriter, r *http.Request) {  
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
		    <li><a href="/facemask/day">Link to day facemask</a></li>  
        <li><a href="/facemask/week">Link to week facemask</a></li> 
				<li><a href="/facemask/month">Link to month facemask</a></li>  
				<li><a href="/facemask/year">Link to year facemask</a></li>  
				<li><a href="/facemask/all">Link to all facemask info</a></li>
    </ul>  
</body>  
</html>`)  
	
} 

func AllFaceMaskHandler(w http.ResponseWriter, r *http.Request) {
  
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	DayFaceMaskHtmlTable(w)
	WeekFaceMaskHtmlTable(w)
	MonthFaceMaskHtmlTable(w)
	YearFaceMaskHtmlTable(w)
}


// 定义处理函数，用于处理/facemask/day路径的请求  
func DayFaceMaskHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	DayFaceMaskHtmlTable(w)
}


// 定义处理函数，用于处理/facemask/week路径的请求  
func WeekFaceMaskHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	WeekFaceMaskHtmlTable(w)
}


// 定义处理函数，用于处理/facemask/month路径的请求
func MonthFaceMaskHandler (w http.ResponseWriter, r *http.Request) {
  
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	MonthFaceMaskHtmlTable(w)
}


// 定义处理函数，用于处理/facemask/year路径的请求
func YearFaceMaskHandler (w http.ResponseWriter, r *http.Request) {
  
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	YearFaceMaskHtmlTable(w)
}