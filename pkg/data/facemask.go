package data


import (
	"fmt"

	hFaceMask "ningan.com/habit-tracking/pkg/habit/facemask"
)

var GlobalFaceMask *hFaceMask.FaceMask

func DealFaceMaskData(input, output, target string) (error) {
	dataMap, err :=readExcelAndCreateMapDataAndWriteFile(input, output, target) 
	if err != nil {
		fmt.Println("Error reading file:", err)  
		return err
	}

	// 打印生成的map  
	for date, rawData := range dataMap {  
		fmt.Printf("Date: %s, RawData: %s\n", date, rawData)  
	}  

	GlobalFaceMask = hFaceMask.NewFaceMask(dataMap)

	err = GlobalFaceMask.GenInfo()
	if err != nil {
		return err
	}

	err = GlobalFaceMask.ConvertInfoToOrderInfo()
	if err != nil {
		return err
	}
	
	err = GlobalFaceMask.CheckFinish()
	if err != nil {
		return err
	}

	return nil
}