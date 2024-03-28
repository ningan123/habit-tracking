package data

import (
	"fmt"

	hPiano "ningan.com/habit-tracking/pkg/piano"
)

var GlobalPiano *hPiano.Piano



func DealPianoData(input, output, target string) (error) {  
	dataMap, err :=readExcelAndCreateMapDataAndWriteFile(input, output, target) 
	if err != nil {
		fmt.Println("Error reading file:", err)  
		return err
	}
  
	// 打印生成的map  
	for date, rawData := range dataMap {  
		fmt.Printf("Date: %s, RawData: %s\n", date, rawData)  
	}  

	GlobalPiano = hPiano.NewPiano(dataMap)

	err = GlobalPiano.GenPianoInfo()
	if err != nil {
		return err
	}

	err = GlobalPiano.ComputePianoTime()
	if err != nil {
		return err
	}

	err = GlobalPiano.ConvertPianoInfoToOrderPianoInfo()
	if err != nil {
		return err
	}

	err = GlobalPiano.CheckFinish()
	if err != nil {
		return err
	}

	err = GlobalPiano.ComputeExtraPianoTime()
	if err != nil {	  
		return err
	}

	return nil
	
}