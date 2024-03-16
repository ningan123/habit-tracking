package reading

import (
	// "fmt"
	"testing"
)





func TestSplitRawInfo(t *testing.T) {
	rawInfo := "《三毛流浪记》,30min;《长安的荔枝》,20min"
	contentInfoList, err :=  SplitRawInfo(rawInfo)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	for _, contentInfo := range contentInfoList {
		t.Log(contentInfo)
	}
}