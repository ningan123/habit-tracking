package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	hData "ningan.com/habit-tracking/pkg/data"
	hRoute "ningan.com/habit-tracking/pkg/route"
)


var (
	port = flag.Int("p", 8888, "port")
)



  

  
func main() {  
	flag.Parse()
	port := fmt.Sprintf(":%d", *port)

	// 
	err := hData.DealReadingData("./data/reading.txt")
	if err != nil {
		log.Fatal(err)
	}

	// 设置路由  
	http.HandleFunc("/", hRoute.RootHandler)  
	http.HandleFunc("/a", hRoute.AHandler)  
	http.HandleFunc("/b", hRoute.BHandler)  
	http.HandleFunc("/reading", hRoute.ReadingHandler) 
	http.HandleFunc("/reading/day", hRoute.DayReadingHandler)  
	http.HandleFunc("/reading/week", hRoute.WeekReadingHandler) 
	http.HandleFunc("/reading/month", hRoute.MonthReadingHandler) 
	
  
	// 启动服务器  
	log.Printf("Starting server on port %s", port)  
	if err := http.ListenAndServe(port, nil); err != nil {  
		log.Fatal(err)  
	}  
}

