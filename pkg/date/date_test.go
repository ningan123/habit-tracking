package date

import (
	"testing"
	"time"
)



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
		{"2023-02-05", 28},
		{"2023-02-28", 28},
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
		{"2024-01-02", 31},
		{"2024-01-03", 31},
		{"2024-01-31", 31},
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



func TestDateWeek(t *testing.T) {
	testCases := []struct {
		dateStr string
		expectedYear int
		expectedWeek int
		expectedWeekDay time.Weekday
	} {
		{"2022-01-01", 2021, 52, time.Saturday},
		{"2022-01-02", 2021, 52, time.Sunday},
		{"2022-01-03", 2022, 1, time.Monday},
		{"2022-01-04", 2022, 1, time.Tuesday},
		{"2022-01-05", 2022, 1, time.Wednesday},
		{"2022-01-06", 2022, 1, time.Thursday},
		{"2022-01-07", 2022, 1, time.Friday},

		{"2022-12-24", 2022, 51, time.Saturday},
		{"2022-12-25", 2022, 51, time.Sunday},
		{"2022-12-26", 2022, 52, time.Monday},
		{"2022-12-27", 2022, 52, time.Tuesday},
		{"2022-12-28", 2022, 52, time.Wednesday},
		{"2022-12-29", 2022, 52, time.Thursday},
		{"2022-12-30", 2022, 52, time.Friday},
		{"2022-12-31", 2022, 52, time.Saturday},

		{"2023-01-01", 2022, 52, time.Sunday},
		{"2023-01-02", 2023, 1, time.Monday},
		{"2023-01-03", 2023, 1, time.Tuesday},
		{"2023-01-04", 2023, 1, time.Wednesday},
		{"2023-01-05", 2023, 1, time.Thursday},
		{"2023-01-06", 2023, 1, time.Friday},
		{"2023-01-07", 2023, 1, time.Saturday},

		{"2023-12-23", 2023, 51, time.Saturday},
		{"2023-12-24", 2023, 51, time.Sunday},
		{"2023-12-25", 2023, 52, time.Monday},
		{"2023-12-26", 2023, 52, time.Tuesday},
		{"2023-12-27", 2023, 52, time.Wednesday},
		{"2023-12-28", 2023, 52, time.Thursday},
		{"2023-12-29", 2023, 52, time.Friday},
		{"2023-12-30", 2023, 52, time.Saturday},
		{"2023-12-31", 2023, 52, time.Sunday},

		{"2024-01-01", 2024, 1, time.Monday},
		{"2024-01-02", 2024, 1, time.Tuesday},
		{"2024-01-03", 2024, 1, time.Wednesday},
		{"2024-01-04", 2024, 1, time.Thursday},
		{"2024-01-05", 2024, 1, time.Friday},
		{"2024-01-06", 2024, 1, time.Saturday},
		{"2024-01-07", 2024, 1, time.Sunday},
		{"2024-01-08", 2024, 2, time.Monday},
		{"2024-01-09", 2024, 2, time.Tuesday},

		{"2024-12-22", 2024, 51, time.Sunday},
		{"2024-12-23", 2024, 52, time.Monday},
		{"2024-12-24", 2024, 52, time.Tuesday},
		{"2024-12-25", 2024, 52, time.Wednesday},
		{"2024-12-26", 2024, 52, time.Thursday},
		{"2024-12-27", 2024, 52, time.Friday},
		{"2024-12-28", 2024, 52, time.Saturday},
		{"2024-12-29", 2024, 52, time.Sunday},
		{"2024-12-30", 2025, 1, time.Monday},
		{"2024-12-31", 2025, 1, time.Tuesday},

		{"2025-12-27", 2025, 52, time.Saturday},
	}

	for _, tc := range testCases {	  
		date, err := convertDateStrToDate(tc.dateStr)
		if err != nil {
			t.Errorf("convertDateStrToDate(%s) returned error: %v", tc.dateStr, err)
		}
		year, week, weekDay := dateWeek(date)
		if year != tc.expectedYear || week != tc.expectedWeek || weekDay != tc.expectedWeekDay {
			t.Errorf("dateWeek(%s) returned (%d, %d, %s), expected (%d, %d, %s)", tc.dateStr, year, week, weekDay, tc.expectedYear, tc.expectedWeek, tc.expectedWeekDay)
			
		}
	}
}



func TestGetDateDetails(t *testing.T) {
	// 测试用例表
	testCases := []struct {
		inputDate        string
		expectedYear     int
		expectMonth      time.Month
		expectedWeekYear int
		expectedWeek      int
		expectedWeekday   string
		expectedDayOfMonth int
		expectedDayOfYear int
		expectedDaysInMonth int
		expectedDaysInYear int
	}{
		{"2021-01-01", 2021, 1, 2020, 53, "五", 1, 1, 31, 365},
		{"2021-01-02", 2021, 1, 2020, 53, "六", 2, 2, 31, 365},
		{"2021-01-03", 2021, 1, 2020, 53, "日", 3, 3, 31, 365},
		{"2021-01-04", 2021, 1, 2021, 1, "一", 4, 4, 31, 365},
		{"2021-01-05", 2021, 1, 2021, 1, "二", 5, 5, 31, 365},
		{"2021-12-27", 2021, 12, 2021, 52, "一", 27, 361, 31, 365},
		{"2021-12-28", 2021, 12, 2021, 52, "二", 28, 362, 31, 365},
		{"2024-01-01", 2024, 1, 2024, 1, "一", 1, 1, 31, 366},
		{"2024-01-02", 2024, 1, 2024, 1, "二", 2, 2, 31, 366},
		{"2024-01-03", 2024, 1, 2024, 1, "三", 3, 3, 31, 366},
		{"2024-01-04", 2024, 1, 2024, 1, "四", 4, 4, 31, 366},
		{"2024-01-05", 2024, 1, 2024, 1, "五", 5, 5, 31, 366},
		{"2024-01-06", 2024, 1, 2024, 1, "六", 6, 6, 31, 366},
		{"2024-01-07", 2024, 1, 2024, 1, "日", 7, 7, 31, 366},
		{"2024-01-08", 2024, 1, 2024, 2, "一", 8, 8, 31, 366},
		{"2024-02-01", 2024, 2, 2024, 5, "四", 1, 32, 29, 366},
		
	}

	// 遍历测试用例表
	for _, tc := range testCases {
		year, month, weekYear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, err := GetDateDetails(tc.inputDate)
		if err != nil {
			t.Errorf("getDateDetails(%s) returned error: %v", tc.inputDate, err)
		}
		if year != tc.expectedYear || month != tc.expectMonth || weekYear != tc.expectedWeekYear || week != tc.expectedWeek || weekday != tc.expectedWeekday || dayOfMonth != tc.expectedDayOfMonth || dayOfYear != tc.expectedDayOfYear || daysInMonth != tc.expectedDaysInMonth || daysInYear != tc.expectedDaysInYear {
				t.Errorf("getDateDetails(%s) returned (%d, %d, %d, %d, %s, %d, %d, %d, %d), expected (%d, %d, %d, %d, %s, %d, %d, %d, %d)", tc.inputDate, year, month, weekYear, week, weekday, dayOfMonth, dayOfYear, daysInMonth, daysInYear, tc.expectedYear, tc.expectMonth, tc.expectedWeekYear, tc.expectedWeek, tc.expectedWeekday, tc.expectedDayOfMonth, tc.expectedDayOfYear, tc.expectedDaysInMonth, tc.expectedDaysInYear)
			}
	}
}


func TestWeekdayToChinese(t *testing.T) {
	testCases := []struct {
		weekday time.Weekday
		expectedChinese string
	}{
		{time.Sunday, "日"},
		{time.Monday, "一"},
		{time.Tuesday, "二"},
		{time.Wednesday, "三"},
		{time.Thursday, "四"},
		{time.Friday, "五"},
		{time.Saturday, "六"},
	}

	for _, tc := range testCases {
		chinese := WeekdayToChinese(tc.weekday)
		if chinese != tc.expectedChinese {
			t.Errorf("weekdayToChinese(%s) returned %s, expected %s", tc.weekday, chinese, tc.expectedChinese)
		}
	}
}