package data

import (
	"fmt"

	hGetup "ningan.com/habit-tracking/pkg/getup"
)

var GlobalGetup *hGetup.Getup

func DealGetupData(fileName string) (error) {
	// 调用readFileToMap函数读取文件并生成map  
	_, dataMap, err := readFileToMap(fileName)  
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

	err = GlobalGetup.CheckFinish()
	if err != nil {
		return err
	}

	return nil
}