package route

import (
	"fmt"
	"net/http"
)

// 定义处理函数，用于处理/getup路径的请求
func GetupHandler(w http.ResponseWriter, r *http.Request) {  
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
		    <li><a href="/getup/day">Link to day getup</a></li>  
        <li><a href="/getup/week">Link to week getup</a></li> 
				<li><a href="/getup/month">Link to month getup</a></li>  
				<li><a href="/getup/year">Link to year getup</a></li>  
				<li><a href="/getup/all">Link to all getup info</a></li>
    </ul>  
</body>  
</html>`)  
	
} 

func AllGetupHandler(w http.ResponseWriter, r *http.Request) {
  
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	DayGetupHtmlTable(w)
	WeekGetupHtmlTable(w)
	MonthGetupHtmlTable(w)
	YearGetupHtmlTable(w)
}


// 定义处理函数，用于处理/getup/day路径的请求  
func DayGetupHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	DayGetupHtmlTable(w)
}


// 定义处理函数，用于处理/getup/week路径的请求  
func WeekGetupHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	WeekGetupHtmlTable(w)
}


// 定义处理函数，用于处理/getup/month路径的请求
func MonthGetupHandler (w http.ResponseWriter, r *http.Request) {
  
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	MonthGetupHtmlTable(w)
}


// 定义处理函数，用于处理/getup/year路径的请求
func YearGetupHandler (w http.ResponseWriter, r *http.Request) {
  
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	YearGetupHtmlTable(w)
}