package route

import "net/http"

// 定义处理函数，用于处理/year路径的请求
func YearHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头，指定内容类型为HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	YearHtmlTable(w)
}
