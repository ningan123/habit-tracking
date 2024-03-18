package date

import (
	"testing"
	"time"
)

func convertDateStrToDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}


func TestDateDetail(t *testing.T) {
  testCases := []struct {
		dateStr string
		expectedYear int
		expectedMonth time.Month
		expectedDay int
	}	{
		{"2022-01-01", 2022, time.January, 1},
		{"2022-02-28", 2022, time.February, 28},
		{"2022-03-01", 2022, time.March, 1},
		{"2023-01-01", 2023, time.January, 1},
		{"2023-01-02", 2023, time.January, 2},
		{"2023-01-03", 2023, time.January, 3},
	}

	for _, tc := range testCases {
		// 执行测试逻辑
		date, err := convertDateStrToDate(tc.dateStr)
		if err != nil {
			t.Errorf("convertDateStrToDate(%s) returned error: %v", tc.dateStr, err)
		}

		year, month, day := dateDetail(date)
		if year != tc.expectedYear || month != tc.expectedMonth || day != tc.expectedDay {
			t.Errorf("DateDetail(%s) = %d, %s, %d; want %d, %s, %d", tc.dateStr, year, month, day, tc.expectedYear, tc.expectedMonth, tc.expectedDay)
		}
	}
}


func TestDaysInMonth(t *testing.T) {
	testCases := []struct {
		dateStr       string
		expectedDay int
	} {
		{"2023-01-01", 31},
		{"2023-02-01", 28},
		{"2023-03-01", 31},
		{"2023-04-01", 30},
		{"2023-05-01", 31},
		{"2023-06-01", 30},
		{"2023-07-01", 31},
		{"2023-08-01", 31},
		{"2023-09-01", 30},
		{"2023-10-01", 31},
		{"2023-11-01", 30},
		{"2023-12-01", 31},
		{"2024-01-01", 31},
		{"2024-02-01", 29},
	}

	for _, tc := range testCases {
		date, err := convertDateStrToDate(tc.dateStr)
		if err != nil {
			t.Errorf("convertDateStrToDate(%s) returned error: %v", tc.dateStr, err)
		}
		day := daysInMonth(date)
		if day != tc.expectedDay {
			t.Errorf("对于输入日期 %s，期望的结果为 %d，实际结果为 %d", tc.dateStr,
				tc.expectedDay, day)
		}
	}
}




func TestDaysInYear(t *testing.T) {
	testCases := []struct {
		dateStr       string
		expectedDay int
	} {
		{"2023-01-01", 365},
		{"2024-01-01", 366},
		{"2025-01-01", 365},
		{"2026-01-01", 365},
	}

	for _, tc := range testCases {
		date, err := convertDateStrToDate(tc.dateStr)
		if err != nil {
			t.Errorf("convertDateStrToDate(%s) returned error: %v", tc.dateStr, err)
		}

		day := daysInYear(date)
		if day != tc.expectedDay {
			t.Errorf("对于输入日期 %s，期望的结果为 %d，实际结果为 %d", tc.dateStr,
				tc.expectedDay, day)
		}		
	}
}


func TestDayOfYear(t *testing.T) {
  testCases := []struct {
		dateStr string
		expectedDay int
	} {
		{"2023-01-01", 1},
		{"2023-01-02", 2},
		{"2023-01-03", 3},

		{"2024-02-28", 59},
		{"2024-02-29", 60},
		{"2024-03-01", 61},
	}

	for _, tc := range testCases {
		// 解析日期字符串为time.Time对象  
		date, err := convertDateStrToDate(tc.dateStr)
		if err != nil {
			t.Errorf("convertDateStrToDate(%s) returned error: %v", tc.dateStr, err)
		}
		// 计算日期是这一年的第几天
		day := dayOfYear(date)
		if day != tc.expectedDay {
			t.Errorf("For date %s, expected %d, but got %d", tc.dateStr, tc.expectedDay, day)
		}
	  
	} 

}



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


