package data

import (
	"testing"
)

func TestReadExcelAndCreateMapDataAndWriteFile(t *testing.T) {
	input := "../../data/real/data.xlsx"
	output := "../../data/real/output_reading.txt"
	target := "阅读"

	dataMap, err :=readExcelAndCreateMapDataAndWriteFile(input, output, target) 
	if err != nil {
		t.Errorf("readExcelAndCreateMapDataAndWriteFile() error = %v", err)
		return
	}
	for k, v := range dataMap {
		t.Logf("key: %s, value: %s", k, v)
	}
}


func TestReadExcelAndCreateMapDataAndWriteFile_getup(t *testing.T) {
	input := "../../data/real/data.xlsx"
	output := "../../data/real/output_getup.txt"
	target := "起床"

	dataMap, err :=readExcelAndCreateMapDataAndWriteFile(input, output, target) 
	if err != nil {
		t.Errorf("readExcelAndCreateMapDataAndWriteFile() error = %v", err)
		return
	}
	for k, v := range dataMap {
		t.Logf("key: %s, value: %s", k, v)
	}
}
