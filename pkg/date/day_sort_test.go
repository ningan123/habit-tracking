package date

import (
	"fmt"
	"sort"
	"testing"
)

func TestByDate (t *testing.T) {
	// 示例map  
	m := map[string]int{  
		"2023-01-02": 10,  
		"2024-11-01": 20,  
		"2023-10-31": 15,  
		"2024-11-02": 25,  
	}  
  
	// 提取键并排序  
	keys := make([]string, 0, len(m))  
	for k := range m {  
		keys = append(keys, k)  
	}  
	sort.Sort(ByDate(keys)) 
  
	// 按照排序后的键顺序提取值到切片  
	values := make([]int, len(keys))  
	for i, k := range keys {  
		values[i] = m[k]  
	}  
  
	// 打印结果  
	fmt.Println("Sorted keys:", keys)  
	fmt.Println("Sorted values:", values)  	
}