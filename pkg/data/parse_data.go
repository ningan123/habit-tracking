package data

import (
	"bufio"
	"os"
	"strings"
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