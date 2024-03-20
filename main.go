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
	dataPath = flag.String("data-path", "./data/test", "data-path")
)



  

  
func main() {  
	flag.Parse()
	port := fmt.Sprintf(":%d", *port)

	// reading
	err := hData.DealReadingData(*dataPath+"/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// getup
	err = hData.DealGetupData(*dataPath+"/data.txt")
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
	http.HandleFunc("/reading/year", hRoute.YearReadingHandler) 
	http.HandleFunc("/reading/all", hRoute.AllReadingHandler) 
  
	// 启动服务器  
	log.Printf("Starting server on port %s", port)  
	if err := http.ListenAndServe(port, nil); err != nil {  
		log.Fatal(err)  
	}  
}

