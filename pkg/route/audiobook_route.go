package route

import (
	"fmt"
	"net/http"
)

// 定义处理函数，用于处理/audiobook路径的请求
func AudiobookHandler(w http.ResponseWriter, r *http.Request) {  
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
		    <li><a href="/audiobook/day">Link to day audiobook</a></li>  
        <li><a href="/audiobook/week">Link to week audiobook</a></li> 
				<li><a href="/audiobook/month">Link to month audiobook</a></li>  
				<li><a href="/audiobook/year">Link to year audiobook</a></li>  
				<li><a href="/audiobook/all">Link to all audiobook info</a></li>
    </ul>  
</body>  
</html>`)  
	
} 

func AllAudiobookHandler(w http.ResponseWriter, r *http.Request) {
  
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	DayAudiobookHtmlTable(w)
	WeekAudiobookHtmlTable(w)
	MonthAudiobookHtmlTable(w)
	YearAudiobookHtmlTable(w)
}


// 定义处理函数，用于处理/audiobook/day路径的请求  
func DayAudiobookHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	DayAudiobookHtmlTable(w)
}


// 定义处理函数，用于处理/audiobook/week路径的请求  
func WeekAudiobookHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	WeekAudiobookHtmlTable(w)
}


// 定义处理函数，用于处理/audiobook/month路径的请求
func MonthAudiobookHandler (w http.ResponseWriter, r *http.Request) {
  
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	MonthAudiobookHtmlTable(w)
}


// 定义处理函数，用于处理/audiobook/year路径的请求
func YearAudiobookHandler (w http.ResponseWriter, r *http.Request) {
  
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	YearAudiobookHtmlTable(w)
}