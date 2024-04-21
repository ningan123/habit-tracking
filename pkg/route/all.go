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
				<li><a href="/piano">Link to piano</a></li> 
				<li><a href="/getup">Link to getup</a></li>
				<li><a href="/sleep">Link to sleep</a></li>
				<li><a href="/facemask">Link to facemask</a></li>
				<li><a href="/skincare">Link to skincare</a></li>
				<li><a href="/audiobook">Link to audiobook</a></li>
				<li><a href="/audiobook">Link to sport</a></li>

				<li><a href="/day">Link to day info</a></li>
				<li><a href="/week">Link to week info</a></li>

    </ul>  
</body>  
</html>`)
}

var GlobalTable = `<!DOCTYPE html>  
<html lang="en">  
<head>  
<meta charset="UTF-8">  
<title>Go Table Example</title>  
<style>  
    .color-cell {  
        background-color: pink;  
    }

	.fixed-header{  
		background-color: #f2f2f2;  
		position: sticky;  
		top: 0; /* 这将使表头固定在视口的顶部 */  
		z-index: 1; /* 确保表头在其他内容之上 */  
		border-collapse: collapse; 
		padding: 8px;
	} 
	.fixed-header2{  
		background-color: #63B8FF;  
		position: sticky;  
		top: 0; /* 这将使表头固定在视口的顶部 */  
		z-index: 1; /* 确保表头在其他内容之上 */  
		border-collapse: collapse; 
		padding: 8px;
	}   

	.fixed-header3{  
		background-color: 	#90EE90;  
		position: sticky;  
		top: 0; /* 这将使表头固定在视口的顶部 */  
		z-index: 1; /* 确保表头在其他内容之上 */  
		border-collapse: collapse; 
		padding: 8px;
	}   
			
</style>  
</head>  
<body>  
<table border="1">`
