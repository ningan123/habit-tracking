package string

import "testing"

func TestStrTrimPrefix(t *testing.T) {
	// 测试用例
	tests := []struct {
		str     string
		prefix  string
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