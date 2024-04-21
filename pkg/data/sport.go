package data

import (
	"fmt"

	hSport "ningan.com/habit-tracking/pkg/sport"
)

var GlobalSport *hSport.Sport

func DealSportData(input, output, target string) error {
	dataMap, err := readExcelAndCreateMapDataAndWriteFile(input, output, target)
	if err != nil {
		fmt.Println("Error Sport file:", err)
		return err
	}

	// 打印生成的map
	for date, rawData := range dataMap {
		fmt.Printf("Date: %s, RawData: %s\n", date, rawData)
	}

	GlobalSport = hSport.NewSport(dataMap)

	err = GlobalSport.GenSportInfo()
	if err != nil {
		return err
	}

	err = GlobalSport.ComputeSportTimes()
	if err != nil {
		return err
	}

	err = GlobalSport.ConvertSportInfoToOrderSportInfo()
	if err != nil {
		return err
	}

	err = GlobalSport.CheckFinish()
	if err != nil {
		return err
	}

	return nil

}
