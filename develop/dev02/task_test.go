package main

import (
	"fmt"
	"testing"
)

func TestString_unpacking(t *testing.T) {
	testTable := []struct {
		str      string
		expected string
		err      error
	}{
		{
			str:      "a4bc2d5e",
			expected: "aaaabccddddde",
			err:      nil,
		},
		{
			str:      "abcd",
			expected: "abcd",
			err:      nil,
		},
		{
			str:      "45",
			expected: "",
			err:      fmt.Errorf("некорректная строка"),
		},
		{
			str:      "",
			expected: "",
			err:      nil,
		},
		{
			str:      `qwe\4\5`,
			expected: "qwe45",
			err:      nil,
		},
		{
			str:      `qwe\45`,
			expected: "qwe44444",
			err:      nil,
		},
		{
			str:      `qwe\\5`,
			expected: `qwe\\\\\`,
			err:      nil,
		},
	}

	for _, testCase := range testTable {
		result, _ := String_unpacking(testCase.str)
		if result != testCase.expected {
			t.Errorf("Incorrect result. Expected %s, got %s", testCase.expected, result)
		}
	}
}
