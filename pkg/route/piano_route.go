package route

import (
	"fmt"
	"net/http"
)
 


// 定义处理函数，用于处理/piano路径的请求  
func PianoHandler(w http.ResponseWriter, r *http.Request) {  
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
		    <li><a href="/piano/day">Link to day piano</a></li>  
        <li><a href="/piano/week">Link to week piano</a></li> 
				<li><a href="/piano/month">Link to month piano</a></li>  
				<li><a href="/piano/year">Link to year piano</a></li>  
				<li><a href="/piano/all">Link to all piano info</a></li>
    </ul>  
</body>  
</html>`)  
	
} 

// 定义处理函数，用于处理/piano/all路径的请求  
func AllPianoHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	DayPianoHtmlTable(w)
	WeekPianoHtmlTable(w)
	MonthPianoHtmlTable(w)
	YearPianoHtmlTable(w)
}


// 定义处理函数，用于处理/piano/day路径的请求  
func DayPianoHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	DayPianoHtmlTable(w)
}


// 定义处理函数，用于处理/piano/week路径的请求  
func WeekPianoHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	WeekPianoHtmlTable(w)
}



// 定义处理函数，用于处理/piano/month路径的请求  
func MonthPianoHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	MonthPianoHtmlTable(w)
}




// 定义处理函数，用于处理/piano/year路径的请求  
func YearPianoHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

  YearPianoHtmlTable(w)
}



