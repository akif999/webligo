package main

import "testing"

func TestIsIncludeMultibyte(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			"apple",
			false,
		},
		{
			"umbrella",
			false,
		},
		{
			"----",
			false,
		},
		{
			"林檎",
			true,
		},
		{
			"かさ",
			true,
		},
		{
			"○",
			true,
		},
		{
			"Goラング",
			true,
		},
	}

	for i, test := range tests {
		got := isIncludeMultibyte(test.input)
		if got != test.expected {
			t.Errorf("tests[%d] got %v, want %v", i, got, test.expected)
		}
	}
}
