package data

import (
	"fmt"

	hSleep "ningan.com/habit-tracking/pkg/habit/sleep"
)

var GlobalSleep *hSleep.Sleep

func DealSleepData(input, output, target string) (error) {
	dataMap, err :=readExcelAndCreateMapDataAndWriteFile(input, output, target) 
	if err != nil {
		fmt.Println("Error reading file:", err)  
		return err
	}

	// 打印生成的map  
	for date, rawData := range dataMap {  
		fmt.Printf("Date: %s, RawData: %s\n", date, rawData)  
	}  

	GlobalSleep = hSleep.NewSleep(dataMap)

	err = GlobalSleep.GenSleepInfo()
	if err != nil {
		return err
	}

	err = GlobalSleep.ConvertSleepInfoToOrderSleepInfo()
	if err != nil {
		return err
	}
	
	err = GlobalSleep.CheckFinish()
	if err != nil {
		return err
	}

	return nil
}