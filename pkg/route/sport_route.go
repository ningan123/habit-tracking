package route

import (
	"fmt"
	"net/http"
)

// 定义处理函数，用于处理/sport路径的请求
func SportHandler(w http.ResponseWriter, r *http.Request) {
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
		    <li><a href="/sport/day">Link to day sport</a></li>  
        <li><a href="/sport/week">Link to week sport</a></li> 
				<li><a href="/sport/month">Link to month sport</a></li>  
				<li><a href="/sport/year">Link to year sport</a></li>  
				<li><a href="/sport/all">Link to all sport info</a></li>
    </ul>  
</body>  
</html>`)

}

func AllSportHandler(w http.ResponseWriter, r *http.Request) {

	// 设置响应头，指定内容类型为HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	DaySportHtmlTable(w)
	WeekSportHtmlTable(w)
	MonthSportHtmlTable(w)
	YearSportHtmlTable(w)
}

// 定义处理函数，用于处理/sport/day路径的请求
func DaySportHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	DaySportHtmlTable(w)
}

// 定义处理函数，用于处理/sport/week路径的请求
func WeekSportHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	WeekSportHtmlTable(w)
}

// 定义处理函数，用于处理/sport/month路径的请求
func MonthSportHandler(w http.ResponseWriter, r *http.Request) {

	// 设置响应头，指定内容类型为HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	MonthSportHtmlTable(w)
}

// 定义处理函数，用于处理/sport/year路径的请求
func YearSportHandler(w http.ResponseWriter, r *http.Request) {

	// 设置响应头，指定内容类型为HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	YearSportHtmlTable(w)
}
