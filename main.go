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
	port         = flag.Int("p", 8888, "port")
	dataPath     = flag.String("data-path", "./data/real", "data-path")
	dataFileName = flag.String("data-filename", "data.xlsx", "data-filename")
	Version      string
)

func main() {
	flag.Parse()
	port := fmt.Sprintf(":%d", *port)
	klog.InfoS("Show version", "version", Version)

	// // reading
	// err := hData.DealReadingData(*dataPath+"/"+*dataFileName, *dataPath+"/output_reading.txt", "阅读")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// piano
	err := hData.DealPianoData(*dataPath+"/"+*dataFileName, *dataPath+"/output_piano.txt", "练琴")
	if err != nil {
		log.Fatal(err)
	}

	// // getup
	// err = hData.DealGetupData(*dataPath+"/"+*dataFileName, *dataPath+"/output_getup.txt", "起床")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // sleep
	// err = hData.DealSleepData(*dataPath+"/"+*dataFileName, *dataPath+"/output_sleep.txt", "睡觉")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // facemask
	// err = hData.DealFaceMaskData(*dataPath+"/"+*dataFileName, *dataPath+"/output_facemask.txt", "敷面膜")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // eyemask
	// err = hData.DealEyeMaskData(*dataPath+"/"+*dataFileName, *dataPath+"/output_eyemask.txt", "敷眼膜")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // skincare
	// err = hData.DealSkinCareData(*dataPath+"/"+*dataFileName, *dataPath+"/output_skincare.txt", "护肤")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // audiobook
	// err = hData.DealAudiobookData(*dataPath+"/"+*dataFileName, *dataPath+"/output_audiobook.txt", "听书")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // sport
	// err = hData.DealSportData(*dataPath+"/"+*dataFileName, *dataPath+"/output_sport.txt", "运动")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 设置路由
	http.HandleFunc("/", hRoute.RootHandler)
	// http.HandleFunc("/reading", hRoute.ReadingHandler)
	// http.HandleFunc("/reading/day", hRoute.DayReadingHandler)
	// http.HandleFunc("/reading/week", hRoute.WeekReadingHandler)
	// http.HandleFunc("/reading/month", hRoute.MonthReadingHandler)
	// http.HandleFunc("/reading/year", hRoute.YearReadingHandler)
	// http.HandleFunc("/reading/all", hRoute.AllReadingHandler)

	http.HandleFunc("/piano", hRoute.PianoHandler)
	http.HandleFunc("/piano/day", hRoute.DayPianoHandler)
	http.HandleFunc("/piano/week", hRoute.WeekPianoHandler)
	http.HandleFunc("/piano/month", hRoute.MonthPianoHandler)
	http.HandleFunc("/piano/year", hRoute.YearPianoHandler)
	http.HandleFunc("/piano/all", hRoute.AllPianoHandler)

	// http.HandleFunc("/getup", hRoute.GetupHandler)
	// http.HandleFunc("/getup/day", hRoute.DayGetupHandler)
	// http.HandleFunc("/getup/week", hRoute.WeekGetupHandler)
	// http.HandleFunc("/getup/month", hRoute.MonthGetupHandler)
	// http.HandleFunc("/getup/year", hRoute.YearGetupHandler)
	// http.HandleFunc("/getup/all", hRoute.AllGetupHandler)

	// http.HandleFunc("/sleep", hRoute.SleepHandler)
	// http.HandleFunc("/sleep/day", hRoute.DaySleepHandler)
	// http.HandleFunc("/sleep/week", hRoute.WeekSleepHandler)
	// http.HandleFunc("/sleep/month", hRoute.MonthSleepHandler)
	// http.HandleFunc("/sleep/year", hRoute.YearSleepHandler)
	// http.HandleFunc("/sleep/all", hRoute.AllSleepHandler)

	// http.HandleFunc("/facemask", hRoute.FaceMaskHandler)
	// http.HandleFunc("/facemask/day", hRoute.DayFaceMaskHandler)
	// http.HandleFunc("/facemask/week", hRoute.WeekFaceMaskHandler)
	// http.HandleFunc("/facemask/month", hRoute.MonthFaceMaskHandler)
	// http.HandleFunc("/facemask/year", hRoute.YearFaceMaskHandler)
	// http.HandleFunc("/facemask/all", hRoute.AllFaceMaskHandler)

	// http.HandleFunc("/skincare", hRoute.SkinCareHandler)
	// http.HandleFunc("/skincare/day", hRoute.DaySkinCareHandler)
	// http.HandleFunc("/skincare/week", hRoute.WeekSkinCareHandler)
	// http.HandleFunc("/skincare/month", hRoute.MonthSkinCareHandler)
	// http.HandleFunc("/skincare/year", hRoute.YearSkinCareHandler)
	// http.HandleFunc("/skincare/all", hRoute.AllSkinCareHandler)

	// http.HandleFunc("/audiobook", hRoute.AudiobookHandler)
	// http.HandleFunc("/audiobook/day", hRoute.DayAudiobookHandler)
	// http.HandleFunc("/audiobook/week", hRoute.WeekAudiobookHandler)
	// http.HandleFunc("/audiobook/month", hRoute.MonthAudiobookHandler)
	// http.HandleFunc("/audiobook/year", hRoute.YearAudiobookHandler)
	// http.HandleFunc("/audiobook/all", hRoute.AllAudiobookHandler)

	// http.HandleFunc("/sport", hRoute.SportHandler)
	// http.HandleFunc("/sport/day", hRoute.DaySportHandler)
	// http.HandleFunc("/sport/week", hRoute.WeekSportHandler)
	// http.HandleFunc("/sport/month", hRoute.MonthSportHandler)
	// http.HandleFunc("/sport/year", hRoute.YearSportHandler)
	// http.HandleFunc("/sport/all", hRoute.AllSportHandler)

	http.HandleFunc("/day", hRoute.DayHandler)
	http.HandleFunc("/week", hRoute.WeekHandler)
	http.HandleFunc("/month", hRoute.MonthHandler)
	http.HandleFunc("/year", hRoute.YearHandler)
	http.HandleFunc("/all", hRoute.AllHandler)

	// 启动服务器
	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
