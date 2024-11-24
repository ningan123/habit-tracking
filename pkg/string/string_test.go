package string

import "testing"

func TestStrTrimPrefix(t *testing.T) {
	// 测试用例
	tests := []struct {
		str      string
		prefix   string
		expected string
	}{
		{"~01:00+1", "~", "01:00+1"},
		{"~01:00", "~", "01:00"},
	}

	for _, test := range tests {
		result := StrTrimPrefix(test.str, test.prefix)
		if result != test.expected {
			t.Errorf("StrTrimPrefix(%q, %q) = %q; expected %q", test.str, test.prefix, result, test.expected)
		}
	}
}

func Test_SortString(t *testing.T) {
	testcases := []struct {
		str         string
		expectedRes string
	}{
		{
			str:         "a<br>c<br>b<br>",
			expectedRes: "a<br>b<br>c<br>",
		},
	}

	for _, test := range testcases {
		result := SortString(test.str)
		if result != test.expectedRes {
			t.Errorf("SortString(%q) = %q; expected %q", test.str, result, test.expectedRes)
		}
	}
}
