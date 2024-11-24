package data

import (
	"fmt"

	hGetup "ningan.com/habit-tracking/pkg/habit/getup"
)

var GlobalGetup *hGetup.Getup

func DealGetupData(input, output, target string) (error) {
	dataMap, err :=readExcelAndCreateMapDataAndWriteFile(input, output, target) 
	if err != nil {
		fmt.Println("Error reading file:", err)  
		return err
	}

	// 打印生成的map  
	for date, rawData := range dataMap {  
		fmt.Printf("Date: %s, RawData: %s\n", date, rawData)  
	}  

	GlobalGetup = hGetup.NewGetup(dataMap)

	err = GlobalGetup.GenGetupInfo()
	if err != nil {
		return err
	}

	err = GlobalGetup.ConvertGetupInfoToOrderGetupInfo()
	if err != nil {
		return err
	}
	
	err = GlobalGetup.CheckFinish()
	if err != nil {
		return err
	}

	return nil
}