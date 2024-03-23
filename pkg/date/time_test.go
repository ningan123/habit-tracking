package date

import "testing"

func TestIsStr1BeforeOrEqualStr2(t *testing.T) {
	// Test cases
	testCases := []struct {
		str1      string
		str2      string
		expected bool
  } {
		{"6:00", "7:30", true},
		{"7:30", "6:00", false},
		{"6:00", "6:00", true},
		{"7:30", "7:30", true},
		{"8:00", "7:30", false},
		{"23:30", "23:30", true},
		{"23:40", "23:30", false},
		{"00:00+1", "23:30", false},
		{"00:10+1", "23:30", false},
	}

	// Iterate through test cases
	for _, tc := range testCases {
		result, err := IsStr1BeforeOrEqualStr2(tc.str1, tc.str2)
		if err != nil {
			t.Errorf("IsStr1BeforeOrEqualStr2(%s, %s) returned error: %s", tc.str1, tc.str2, err)
		}
		if result != tc.expected {
			t.Errorf("IsStr1BeforeOrEqualStr2(%s, %s) = %t; expected %t", tc.str1, tc.str2, result, tc.expected)
		}
	}
}