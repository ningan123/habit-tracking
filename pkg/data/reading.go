package data

import (
	"fmt"

	hReading "ningan.com/habit-tracking/pkg/reading"
)

var GlobalReading *hReading.Reading



func DealReadingData(fileName string) (error) {  
	// 调用readFileToMap函数读取文件并生成map  
	dataMap, _, err := readFileToMap(fileName)  
	if err != nil {  
		fmt.Println("Error reading file:", err)  
		return err
	}  
  
	// 打印生成的map  
	for date, rawData := range dataMap {  
		fmt.Printf("Date: %s, RawData: %s\n", date, rawData)  
	}  

	GlobalReading = hReading.NewReading(dataMap)

	err = GlobalReading.GenReadingInfo()
	if err != nil {
		return err
	}

	err = GlobalReading.ComputeReadingTime()
	if err != nil {
		return err
	}

	err = GlobalReading.ConvertReadingInfoToOrderReadingInfo()
	if err != nil {
		return err
	}

	err = GlobalReading.CheckFinish()
	if err != nil {
		return err
	}

	GlobalReading.PrintReadingInfo()

	return nil
	
}