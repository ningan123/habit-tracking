package data

import (
	"bufio"
	"os"
	"strings"

	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// readFileToMap 读取文件内容并返回一个map，其中键是日期，值是rawData
func readFileToMap(fileName string) (map[string]string,  error) {  
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
			readingRawData := strings.TrimSpace(fields[1]) // 去除可能的空白字符 
			dataMap[date] = readingRawData
		}  
	}  
  
	if err := scanner.Err(); err != nil {  
		return nil, err  
	}  
  
	return dataMap, nil  
}  




// findColumnIndex 找到包含指定标题的列索引  
func findColumnIndex(sheet *excelize.File, target string) (int, error) {  
	rows := sheet.GetRows("Sheet1") // 假设数据在Sheet1中  
	for colIndex, col := range rows[0] { // 第一行是标题行  
		if strings.EqualFold(col, target) {  
			return colIndex, nil  
		}  
	}  
	return -1, fmt.Errorf("column %q not found", target)  
}  
  
// readExcelAndCreateMapDataAndWriteFile 读取Excel文件，找到目标列，并创建map，然后将数据写到目标文件中  
func readExcelAndCreateMapDataAndWriteFile(input string, output string, target string) (map[string]string, error) {  
	f, err := excelize.OpenFile(input)  
	if err != nil {  
		return nil, err  
	}  
  
	columnIndex, err := findColumnIndex(f, target)  
	if err != nil {  
		return nil, err  
	}  


	file, err := os.Create(output)  
	if err != nil {  
		return nil, err 
	}  
	defer file.Close()  
  
	dataMap := make(map[string]string)  
	rows := f.GetRows("Sheet1") // 假设数据在Sheet1中  
	for _, row := range rows[1:] { // 跳过标题行  
		if len(row) > columnIndex {  
			date := row[0]  
			data := row[columnIndex]  
			dataMap[date] = data  

			_, err := fmt.Fprintf(file, "%s %s\n", date, data)  
			if err != nil {  
				return nil, err   
			}  
		}  
	}   
	return dataMap, nil  
}  
  
