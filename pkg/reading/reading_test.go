package reading

import (
	"fmt"
	"testing"
)

func TestGenMonthAndWeekAndDayReadingInfo(t *testing.T) {
	// rawInfo := []string{
	// 	"2024-01-01 《三毛流浪记》,20min",
	// 	"2024-01-02 《三毛流浪记》,30min",
	// 	"2024-01-02 《三三三》,30min",
	// 	"2024-01-03 《三毛流浪记》,20min",
	// 	"2024-01-04 《三体》,20min",
	// 	"2024-01-05 《三体》,20min",
	// 	"2024-01-06 《太白金星有点烦》,30min",
	// 	"2024-01-07 《三体》,20min",
	// 	"2024-01-08 《三体》,20min",
	// 	"2024-01-09 《三体》,30min",
	// }

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

	reading := NewReading(rawInfo)

	// test1
	err := reading.GenMonthAndWeekAndDayReadingInfo()
	if err != nil {
		t.Error(err)
	}
	fmt.Println()
	reading.Print()

	// test2 
	err = reading.ComputReadingTime()
	if err != nil {
		t.Error(err)
	}
	fmt.Println()
	reading.Print()
}