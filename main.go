package main

import (
	"os"

	reading "ningan.com/habit-tracking/pkg/reading"
)


func main() {
	rawInfo := map[string]string{
		"2024-01-01": "《三毛流浪记》,20min",
		"2024-01-02": " 《三毛流浪记》,30min",
		"2024-01-03": " 《三毛流浪记》,20min",
		"2024-01-04": "《三体》,20min",
		"2024-01-05": " 《三体》,20min",
		"2024-01-06": " 《太白金星有点烦》,30min",
		"2024-01-07": " 《三体》,20min",
		"2024-01-08": " 《三体》,20min",
		"2024-01-09": " 《三体》,30min",
	}

	reading := reading.NewReading(rawInfo)

	err := reading.GenMonthAndWeekAndDayReadingInfo()
	if err != nil {
		os.Exit(1)
	}
	reading.Print()

	err = reading.ComputReadingTime()
	if err != nil {
		os.Exit(1)
	}
	reading.Print()	
}