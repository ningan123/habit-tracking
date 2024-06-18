package date

import (
	"fmt"
	"testing"
	"time"
)

func Test_ParseDuration(t *testing.T) {
	// 测试用例表
	testCases := []struct {
		durationStr  string
		expectResult time.Duration
	}{
		{"15min", time.Duration(15) * time.Minute},
		{"1h", time.Duration(1) * time.Hour},
		{"2h30min", time.Duration(2)*time.Hour + time.Duration(30)*time.Minute},
		{"60min", time.Duration(1) * time.Hour},
		{"110min", time.Duration(1)*time.Hour + time.Duration(50)*time.Minute},
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

func Test_FormarDuration(t *testing.T) {
	// 测试用例表
	testCases := []struct {
		str          string
		expectResult string
	}{
		{"1h", "1h"},
		{"70min", "1h10min"},
	}

	for _, tc := range testCases {
		result, err := FormatDuration(tc.str)
		if err != nil {
			t.Errorf("对于输入时长 %s ，期望的结果为 %s，实际结果为 %s", tc.str, tc.expectResult, result)
		}
		if result != tc.expectResult {
			t.Errorf("对于输入时长 %s ，期望的结果为 %s，实际结果为 %s", tc.str, tc.expectResult, result)
		}
	}

}

func Test_FormatDurationSub(t *testing.T) {
	// 测试用例表
	testCases := []struct {
		durationStr1 string
		durationStr2 string
		expectResult string
	}{
		{"1h", "30min", "30min"},
		{"1h20min", "1h", "20min"},
		{"1h30min", "30min", "1h"},
		{"1h20min", "1h20min", "0"},

		{"1h20min", "1h30min", "-10min"},
		{"1h20min", "2h20min", "-1h"},
		{"20min", "2h10min", "-1h50min"},
	}

	for _, tc := range testCases {
		result, err := FormatDurationSub(tc.durationStr1, tc.durationStr2)
		if err != nil {
			t.Errorf("对于输入时长 %s 和 %s，期望的结果为 %s，实际结果为 %s", tc.durationStr1, tc.durationStr2, tc.expectResult, result)
		}
		if result != tc.expectResult {
			t.Errorf("对于输入时长 %s 和 %s，期望的结果为 %s，实际结果为 %s", tc.durationStr1, tc.durationStr2, tc.expectResult, result)
		}
	}
}

func Test_FormatDurationSum(t *testing.T) {
	// 测试用例表
	testCases := []struct {
		durationStr1 string
		durationStr2 string
		expectResult string
	}{
		{"15min", "1h55min", "2h10min"},
		{"1h", "1h55min", "2h55min"},
		{"1h55min", "1h55min", "3h50min"},
		{"1h", "1h", "2h"},
		{"15min", "45min", "1h"},
		{"0min", "60min", "1h"},
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

func Test_FormatDurationMultiply(t *testing.T) {
	// 测试用例表
	testCases := []struct {
		durationStr  string
		multiplier   int
		expectResult string
	}{
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

func Test_IsActualDurationLongerOrEqualToTargetDuration(t *testing.T) {
	// 测试用例表
	testCases := []struct {
		actualDurationStr string
		targetDurationStr string
		expectResult      bool
	}{
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
