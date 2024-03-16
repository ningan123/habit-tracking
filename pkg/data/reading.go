package data

import (
	"bufio"
	"fmt"

	"os"
	"strings"

	hReading "ningan.com/habit-tracking/pkg/reading"
)

var GlobalReading *hReading.Reading

// readFileToMap 读取文件内容并返回一个map，其中键是日期，值是rawData
func readFileToMap(fileName string) (map[string]string, error) {  
	file, err := os.Open(fileName)  
	if err != nil {  
		return nil, err  
	}  
	defer file.Close()  
  
	dataMap := make(map[string]string)  
	scanner := bufio.NewScanner(file)  
	for scanner.Scan() {  
		line := scanner.Text()  
		fields := strings.Split(line, " ") // 假设每列之间使用逗号分隔  
		if len(fields) >= 2 {  
			date := strings.TrimSpace(fields[0]) // 去除可能的空白字符  
			rawData := strings.TrimSpace(fields[1]) // 去除可能的空白字符  
			dataMap[date] = rawData  
		}  
	}  
  
	if err := scanner.Err(); err != nil {  
		return nil, err  
	}  
  
	return dataMap, nil  
}  

func DealReadingData(fileName string) (error) {  
	// 调用readFileToMap函数读取文件并生成map  
	dataMap, err := readFileToMap(fileName)  
	if err != nil {  
		fmt.Println("Error reading file:", err)  
		return err
	}  
  
	// 打印生成的map  
	for date, rawData := range dataMap {  
		fmt.Printf("Date: %s, RawData: %s\n", date, rawData)  
	}  

	GlobalReading = hReading.NewReading(dataMap)

	err = GlobalReading.GenMonthAndWeekAndDayReadingInfo()
	if err != nil {
		return err
	}
	GlobalReading.Print()

	err = GlobalReading.ComputDayReadingTime()
	if err != nil {
		return err
	}
	
	err = GlobalReading.ComputWeekReadingTime()
	if err != nil {
		return err
	}

	err = GlobalReading.ComputMonthReadingTime()
	if err != nil {
		return err
	}
	GlobalReading.Print()

	return nil
	
}