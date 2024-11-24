package data


import (
	"fmt"

	hEyeMask "ningan.com/habit-tracking/pkg/habit/eyemask"
)

var GlobalEyeMask *hEyeMask.EyeMask

func DealEyeMaskData(input, output, target string) (error) {
	dataMap, err :=readExcelAndCreateMapDataAndWriteFile(input, output, target) 
	if err != nil {
		fmt.Println("Error reading file:", err)  
		return err
	}

	// 打印生成的map  
	for date, rawData := range dataMap {  
		fmt.Printf("Date: %s, RawData: %s\n", date, rawData)  
	}  

	GlobalEyeMask = hEyeMask.NewEyeMask(dataMap)

	err = GlobalEyeMask.GenInfo()
	if err != nil {
		return err
	}

	err = GlobalEyeMask.ConvertInfoToOrderInfo()
	if err != nil {
		return err
	}
	
	err = GlobalEyeMask.CheckFinish()
	if err != nil {
		return err
	}

	return nil
}