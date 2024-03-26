package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"k8s.io/klog/v2"
	hData "ningan.com/habit-tracking/pkg/data"
	hRoute "ningan.com/habit-tracking/pkg/route"
)


var (
	port = flag.Int("p", 8888, "port")
	dataPath = flag.String("data-path", "./data/real/", "data-path")
	Version string
)

  
func main() {  
	flag.Parse()
	port := fmt.Sprintf(":%d", *port)
	klog.InfoS("Show version", "version", Version)

	// reading
	err := hData.DealReadingData(*dataPath+"/data.xlsx", *dataPath+"/output_reading.txt", "阅读")
	if err != nil {
		log.Fatal(err)
	}

	// piano
	err = hData.DealPianoData(*dataPath+"/data.xlsx", *dataPath+"/output_piano.txt", "练琴")
	if err != nil {
		log.Fatal(err)
	}	

	// getup
	err = hData.DealGetupData(*dataPath+"/data.xlsx", *dataPath+"/output_getup.txt", "起床")
	if err != nil {
		log.Fatal(err)
	}

	// sleep
	err = hData.DealSleepData(*dataPath+"/data.xlsx", *dataPath+"/output_sleep.txt", "睡觉")
	if err != nil {
		log.Fatal(err)
	}

	// facemask
	err = hData.DealFaceMaskData(*dataPath+"/data.xlsx", *dataPath+"/output_facemask.txt", "敷面膜")
	if err != nil {
		log.Fatal(err)
	}

	// 设置路由  
	http.HandleFunc("/", hRoute.RootHandler)  
	http.HandleFunc("/reading", hRoute. ReadingHandler) 
	http.HandleFunc("/reading/day", hRoute.DayReadingHandler)  
	http.HandleFunc("/reading/week", hRoute.WeekReadingHandler) 
	http.HandleFunc("/reading/month", hRoute.MonthReadingHandler) 
	http.HandleFunc("/reading/year", hRoute.YearReadingHandler) 
	http.HandleFunc("/reading/all", hRoute.AllReadingHandler) 

	http.HandleFunc("/piano", hRoute.PianoHandler) 
	http.HandleFunc("/piano/day", hRoute.DayPianoHandler)  
	http.HandleFunc("/piano/week", hRoute.WeekPianoHandler) 
	http.HandleFunc("/piano/month", hRoute.MonthPianoHandler) 
	http.HandleFunc("/piano/year", hRoute.YearPianoHandler) 
	http.HandleFunc("/piano/all", hRoute.AllPianoHandler) 

	http.HandleFunc("/getup", hRoute.GetupHandler)
	http.HandleFunc("/getup/day", hRoute.DayGetupHandler)  
	http.HandleFunc("/getup/week", hRoute.WeekGetupHandler)  
	http.HandleFunc("/getup/month", hRoute.MonthGetupHandler)  
	http.HandleFunc("/getup/year", hRoute.YearGetupHandler)
	http.HandleFunc("/getup/all", hRoute.AllGetupHandler) 

	http.HandleFunc("/sleep", hRoute.SleepHandler)
	http.HandleFunc("/sleep/day", hRoute.DaySleepHandler)  
	http.HandleFunc("/sleep/week", hRoute.WeekSleepHandler)  
	http.HandleFunc("/sleep/month", hRoute.MonthSleepHandler)  
	http.HandleFunc("/sleep/year", hRoute.YearSleepHandler)
	http.HandleFunc("/sleep/all", hRoute.AllSleepHandler) 

	http.HandleFunc("/facemask", hRoute.FaceMaskHandler)
	http.HandleFunc("/facemask/day", hRoute.DayFaceMaskHandler)  
	http.HandleFunc("/facemask/week", hRoute.WeekFaceMaskHandler)  
	http.HandleFunc("/facemask/month", hRoute.MonthFaceMaskHandler)  
	http.HandleFunc("/facemask/year", hRoute.YearFaceMaskHandler)
	http.HandleFunc("/facemask/all", hRoute.AllFaceMaskHandler) 

	http.HandleFunc("/day", hRoute.DayHandler)
	http.HandleFunc("/week", hRoute.WeekHandler)  

  
	// 启动服务器  
	log.Printf("Starting server on port %s", port)  
	if err := http.ListenAndServe(port, nil); err != nil {  
		log.Fatal(err)  
	}  
}

