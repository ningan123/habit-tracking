package data

import (
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize"
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


func Test_findRowIndex(t *testing.T) {
	input := "../../data/real/data.xlsx"
	testcases :=  []struct {
		dateStr string
		expectedRow int
	}	{
		{"2024-01-01", 1},
		{"2024-01-02", 2},
		{"2024-01-03", 3},
		{"2024-01-31", 31},
		{"2024-02-01", 32},
		{"2024-02-02", 33},
	}
		
	f, err := excelize.OpenFile(input)  
	if err != nil {  
		t.Errorf("open file error = %v", err)
	}  

	for _, tc := range testcases {
		rowNum, err :=  findRowIndex(f, "Sheet1", tc.dateStr)
		if err != nil {
			t.Errorf("findRowByDate() error = %v", err)
		} 

		if rowNum != tc.expectedRow {
			t.Errorf("findRowByDate() = %v, want %v", rowNum, tc.expectedRow)
		}
	}
}
