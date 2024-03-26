package route

import (
	"fmt"
	"net/http"
)

// 定义处理函数，用于处理/skincare路径的请求
func SkinCareHandler(w http.ResponseWriter, r *http.Request) {  
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
		    <li><a href="/skincare/day">Link to day skincare</a></li>  
        <li><a href="/skincare/week">Link to week skincare</a></li> 
				<li><a href="/skincare/month">Link to month skincare</a></li>  
				<li><a href="/skincare/year">Link to year skincare</a></li>  
				<li><a href="/skincare/all">Link to all skincare info</a></li>
    </ul>  
</body>  
</html>`)  
	
} 

func AllSkinCareHandler(w http.ResponseWriter, r *http.Request) {
  
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	DaySkinCareHtmlTable(w)
	WeekSkinCareHtmlTable(w)
	MonthSkinCareHtmlTable(w)
	YearSkinCareHtmlTable(w)
}


// 定义处理函数，用于处理/skincare/day路径的请求  
func DaySkinCareHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	DaySkinCareHtmlTable(w)
}


// 定义处理函数，用于处理/skincare/week路径的请求  
func WeekSkinCareHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	WeekSkinCareHtmlTable(w)
}


// 定义处理函数，用于处理/skincare/month路径的请求
func MonthSkinCareHandler (w http.ResponseWriter, r *http.Request) {
  
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	MonthSkinCareHtmlTable(w)
}


// 定义处理函数，用于处理/skincare/year路径的请求
func YearSkinCareHandler (w http.ResponseWriter, r *http.Request) {
  
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	YearSkinCareHtmlTable(w)
}