package date

import (
	"fmt"
	"testing"
	"time"
)


func TestGetDateDetails(t *testing.T) {
	// 测试用例表
	testCases := []struct {
		inputDate        string
		expectedYear     int
		expectedDayOfYear int
		expectedMonth    time.Month
		expectedDayOfMonth int
		expectedWeekNum  int
		expectedWeekday  time.Weekday
		expectedErrorMsg string
	}{
		// // 2022年年初8天
		// {"2022-01-01", 2021, 52, time.Saturday, ""},   // 周六
		// {"2022-01-02", 2022, 1, time.Sunday, ""},     // 周日
		// {"2022-01-03", 2022, 1, time.Monday, ""},     // 周一
		// {"2022-01-04", 2022, 1, time.Tuesday, ""},    // 周二
		// {"2022-01-05", 2022, 1, time.Wednesday, ""},  // 周三
		// {"2022-01-06", 2022, 1, time.Thursday, ""},   // 周四
		// {"2022-01-07", 2022, 1, time.Friday, ""},     // 周五
		// {"2022-01-08", 2022, 1, time.Saturday, ""},   // 周六
		// // 2022年年末8天
		// {"2022-12-25", 2022, 51, time.Sunday, ""},    // 周日
		// {"2022-12-26", 2022, 52, time.Monday, ""},    // 周一
		// {"2022-12-27", 2022, 52, time.Tuesday, ""},   // 周二
		// {"2022-12-28", 2022, 52, time.Wednesday, ""}, // 周三
		// {"2022-12-29", 2022, 52, time.Thursday, ""},  // 周四
		// {"2022-12-30", 2022, 52, time.Friday, ""},    // 周五
		// {"2022-12-31", 2022, 52, time.Saturday, ""},  // 周六
		// {"2023-01-01", 2023, 1, time.Sunday, ""},     // 周日
		// // 2023年年初8天
		// {"2023-01-01", 2023, 1, time.Monday, ""},     // 周一
		// {"2023-01-02", 2023, 1, time.Tuesday, ""},    // 周二
		// {"2023-01-03", 2023, 1, time.Wednesday, ""},  // 周三
		// {"2023-01-04", 2023, 1, time.Thursday, ""},   // 周四
		// {"2023-01-05", 2023, 1, time.Friday, ""},     // 周五
		// {"2023-01-06", 2023, 1, time.Saturday, ""},   // 周六
		// {"2023-01-07", 2023, 1, time.Sunday, ""},     // 周日
		// {"2023-01-08", 2023, 2, time.Monday, ""},     // 周一
		// // 2023年年末8天
		// {"2023-12-25", 2023, 52, time.Monday, ""},    // 周一
		// {"2023-12-26", 2023, 52, time.Tuesday, ""},   // 周二
		// {"2023-12-27", 2023, 52, time.Wednesday, ""}, // 周三
		// {"2023-12-28", 2023, 52, time.Thursday, ""},  // 周四
		// {"2023-12-29", 2023, 52, time.Friday, ""},    // 周五
		// {"2023-12-30", 2023, 52, time.Saturday, ""},  // 周六
		// {"2023-12-31", 2023, 52, time.Sunday, ""},    // 周日
		// 2024年年初8天
		{"2024-01-01", 2024, 1, 1, 1, 1, time.Monday, ""},     // 周一
		{"2024-01-02", 2024, 2, 1, 2, 1, time.Tuesday, ""},
		{"2024-01-03", 2024, 3, 1, 3, 1, time.Wednesday, ""},
		{"2024-01-04", 2024, 4, 1, 4, 1, time.Thursday, ""},
		{"2024-01-05", 2024, 5, 1, 5, 1, time.Friday, ""},
		{"2024-01-06", 2024, 6, 1, 6, 1, time.Saturday, ""},
		{"2024-01-07", 2024, 7, 1, 7, 1, time.Sunday, ""},
		{"2024-01-08", 2024, 8, 1, 8, 2, time.Monday, ""},	
		// 2024年年末8天
		// {"2024-12-25", 2024, 12, 25, 52, time.Wednesday, ""},
		// {"2024-12-26", 2024, 12, 26, 52, time.Thursday, ""},
		// {"2024-12-27", 2024, 12, 27, 52, time.Friday, ""},
		// {"2024-12-28", 2024, 12, 28, 52, time.Saturday, ""},
		// {"2024-12-29", 2024, 12, 29, 52, time.Sunday, ""},
		// {"2024-12-30", 2024, 53, time.Monday, ""},
		// {"2024-12-31", 2024, 53, time.Tuesday, ""},

	}

	// 遍历测试用例表
	for _, tc := range testCases {
		year, dayOfYear, month, dayOfMonth, weekNum, weekday, err := GetDateDetails(tc.inputDate)

		// 检查错误情况
		if tc.expectedErrorMsg != "" {
			if err == nil || err.Error() != tc.expectedErrorMsg {
				t.Errorf("对于输入日期 %s，期望错误: %s，实际错误：%v", tc.inputDate, tc.expectedErrorMsg, err)
			}
			continue
		} else {
			if err != nil {
				t.Errorf("对于输入日期 %s，不应有错误: %v", tc.inputDate, err)
				continue
			}
		}

		// 检查结果
		if year != tc.expectedYear || dayOfYear != tc.expectedDayOfYear || month != tc.expectedMonth || dayOfMonth != tc.expectedDayOfMonth || weekNum != tc.expectedWeekNum || weekday != tc.expectedWeekday {
			t.Errorf("对于输入日期 %s，期望的结果为 年份：%d，日数：%d，月份：%d，日：%d，周数：%d，星期几：%s，实际结果为 年份：%d，日数：%d，月份：%d，日：%d，周数：%d，星期几：%s",
				tc.inputDate, tc.expectedYear, tc.expectedDayOfYear, tc.expectedMonth, tc.expectedDayOfMonth, tc.expectedWeekNum, tc.expectedWeekday, year, dayOfYear, month, dayOfMonth, weekNum, weekday)
			// t.Errorf("对于输入日期 %s，期望的结果为 年份：%d，周数：%d，星期几：%s，实际结果为 年份：%d，周数：%d，星期几：%s",
				// tc.inputDate, tc.expectedYear, tc.expectedWeekNum, tc.expectedWeekday, year, weekNum, weekday)
		}
	}
}






func TestDayOfYear(t *testing.T) {
  // 假设我们的日期字符串是"2023-07-18"且布局是"2006-01-02"  
	// dateStr := "2023-07-18"  
	// dateStr := "2024-01-01"  
	dateStr := "2024-01-02"  
	layout := "2006-01-02"  
  
	// 计算这一天是一年中的第几天  
	dayNum, err := dayOfYearFromString(dateStr, layout)  
	if err != nil {  
		fmt.Println("Error parsing date:", err)  
		return  
	}  
	fmt.Printf("The day %s is the %dth day of the year.\n", dateStr, dayNum)  

}


func TestParseDuration(t *testing.T) {
	// 测试用例表
	testCases := []struct {
		durationStr        string
		expectResult time.Duration
	} {
		{"15min", time.Duration(15)*time.Minute},
		{"1h", time.Duration(1)*time.Hour},
		{"2h30min", time.Duration(2)*time.Hour + time.Duration(30)*time.Minute},
	}

	for _, tc := range testCases {
		result, err := parseDuration(tc.durationStr)
		if err != nil {
			t.Errorf("对于输入时长 %s，期望的结果为 %s，实际结果为 %s", tc.durationStr, tc.expectResult, result)
		}
		
		if result != tc.expectResult {
			t.Errorf("对于输入时长 %s，期望的结果为 %s，实际结果为 %s", tc.durationStr, tc.expectResult, result)
		}
		fmt.Println(result)
	}
}


func TestFormatDurationSum(t *testing.T) {
	// 测试用例表
	testCases := []struct {
		durationStr1        string
		durationStr2     string
		expectResult string
	} {
		{"15min", "1h55min", "2h10min"},
		{"1h", "1h55min", "2h55min"},
		{"1h55min", "1h55min", "3h50min"},
		{"1h", "1h", "2h"},
		{"15min", "45min", "1h"},
	}

	for _, tc := range testCases {
		result, err := FormatDurationSum(tc.durationStr1, tc.durationStr2)
		if err != nil {
			t.Errorf("对于输入时间 %s 和 %s，期望的结果为 %s，但出现错误：%s", tc.durationStr1, tc.durationStr2, tc.expectResult, err.Error())
		}
		if result != tc.expectResult {
			t.Errorf("对于输入时间 %s 和 %s，期望的结果为 %s，但实际结果为 %s", tc.durationStr1, tc.durationStr2, tc.expectResult, result)
		}
	}
}

func TestFormatDurationMultiply(t *testing.T) {
	// 测试用例表
	testCases := []struct {	  
		durationStr        string
		multiplier       int
		expectResult string
	} {
		{"15min", 2, "30min"},
		{"1h", 2, "2h"},
		{"1h55min", 2, "3h50min"},
	}

	for _, tc := range testCases {
		result, err := FormatDurationMultiply(tc.durationStr, tc.multiplier)
		if err != nil {
			t.Errorf("对于输入时间 %s 和 %d，期望的结果为 %s，但出现错误：%s", tc.durationStr, tc.multiplier, tc.expectResult, err.Error())
		}
		if result != tc.expectResult {
			t.Errorf("对于输入时间 %s 和 %d，期望的结果为 %s，但实际结果为 %s", tc.durationStr, tc.multiplier, tc.expectResult, result)
		}
	}
}


func TestIsActualDurationLonger(t *testing.T) {
	// 测试用例表
	testCases := []struct {
		actualDurationStr        string
		targetDurationStr     string
		expectResult bool
	} {
		{"15min", "1h55min", false},
		{"1h", "1h55min", false},
		{"1h55min", "1h55min", true},
		{"2h", "1h55min", true},
		{"1h56min", "1h55min", true},
		{"1h55min", "55min", true},
	}

	for _, tc := range testCases {
	  
		result, err := IsActualDurationLongerOrEqualToTargetDuration(tc.actualDurationStr, tc.targetDurationStr)
		if err != nil {
			t.Errorf("Error parsing durations: %v", err)
		}
		if result != tc.expectResult {
			t.Errorf("Expected %v for %s vs %s, but got %v", tc.expectResult, tc.actualDurationStr, tc.targetDurationStr, result)
		}
	}
  
}