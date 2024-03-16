package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)


var (
	port = flag.Int("p", 8888, "port")
)


// 定义一个处理函数，用于返回包含两个链接的HTML页面  
func rootHandler(w http.ResponseWriter, r *http.Request) {  
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
        <li><a href="/a">Link to A</a></li>  
        <li><a href="/b">Link to B</a></li>  
    </ul>  
</body>  
</html>`)  
}  
  
// 定义处理函数，用于处理/a路径的请求  
func aHandler(w http.ResponseWriter, r *http.Request) {  
	fmt.Fprint(w, "You are at path /a")  
}  
  
// 定义处理函数，用于处理/b路径的请求  
func bHandler(w http.ResponseWriter, r *http.Request) {  
	fmt.Fprint(w, "You are at path /b")  
}  
  
  
func main() {  
	flag.Parse()
	port := fmt.Sprintf(":%d", *port)

	

	// 设置路由  
	http.HandleFunc("/", rootHandler)  
	http.HandleFunc("/a", aHandler)  
	http.HandleFunc("/b", bHandler)  
  
	// 启动服务器  
	log.Printf("Starting server on port %s", port)  
	if err := http.ListenAndServe(port, nil); err != nil {  
		log.Fatal(err)  
	}  
}