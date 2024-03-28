package data

import (
	"fmt"

	hAudiobook "ningan.com/habit-tracking/pkg/audiobook"
)

var GlobalAudiobook *hAudiobook.Audiobook

func DealAudiobookData(input, output, target string) (error) {
	dataMap, err :=readExcelAndCreateMapDataAndWriteFile(input, output, target) 
	if err != nil {
		fmt.Println("Error reading file:", err)  
		return err
	}

	// 打印生成的map  
	for date, rawData := range dataMap {  
		fmt.Printf("Date: %s, RawData: %s\n", date, rawData)  
	}  

	GlobalAudiobook = hAudiobook.NewAudiobook(dataMap)

	err = GlobalAudiobook.GenInfo()
	if err != nil {
		return err
	}

	err = GlobalAudiobook.ConvertInfoToOrderInfo()
	if err != nil {
		return err
	}

	err = GlobalAudiobook.ComputeFinishBooks()
	if err != nil {
		return err
	}
	
	err = GlobalAudiobook.CheckFinish()
	if err != nil {
		return err
	}

	return nil
}