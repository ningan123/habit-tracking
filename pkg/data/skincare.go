package data


import (
	"fmt"

	hSkinCare "ningan.com/habit-tracking/pkg/habit/skincare"
)

var GlobalSkinCare *hSkinCare.SkinCare

func DealSkinCareData(input, output, target string) (error) {
	dataMap, err :=readExcelAndCreateMapDataAndWriteFile(input, output, target) 
	if err != nil {
		fmt.Println("Error reading file:", err)  
		return err
	}

	// 打印生成的map  
	for date, rawData := range dataMap {  
		fmt.Printf("Date: %s, RawData: %s\n", date, rawData)  
	}  

	GlobalSkinCare = hSkinCare.NewSkinCare(dataMap)

	err = GlobalSkinCare.GenInfo()
	if err != nil {
		return err
	}

	err = GlobalSkinCare.ConvertInfoToOrderInfo()
	if err != nil {
		return err
	}
	
	err = GlobalSkinCare.CheckFinish()
	if err != nil {
		return err
	}

	return nil
}