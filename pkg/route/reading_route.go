package route

import (
	"fmt"
	"net/http"
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
				<li><a href="/getup">Link to getup</a></li>
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
				<li><a href="/reading/year">Link to year reading</a></li>  
				<li><a href="/reading/all">Link to all reading info</a></li>
    </ul>  
</body>  
</html>`)  
	
} 

// 定义处理函数，用于处理/reading/all路径的请求  
func AllReadingHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	DayReadingHtmlTable(w)
	WeekReadingHtmlTable(w)
	MonthReadingHtmlTable(w)
	YearReadingHtmlTable(w)
}


// 定义处理函数，用于处理/reading/day路径的请求  
func DayReadingHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	DayReadingHtmlTable(w)
}


// 定义处理函数，用于处理/reading/week路径的请求  
func WeekReadingHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	WeekReadingHtmlTable(w)
}



// 定义处理函数，用于处理/reading/month路径的请求  
func MonthReadingHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")  
  
	MonthReadingHtmlTable(w)
}




// 定义处理函数，用于处理/reading/year路径的请求  
func YearReadingHandler (w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML  
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

  YearReadingHtmlTable(w)
}



