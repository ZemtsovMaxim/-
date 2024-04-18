package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrep(t *testing.T) {
	lines := []string{
		"apple banana cherry",
		"apple",
		"banana apple",
		"cherry",
	}

	testTable := []struct {
		name     string
		lines    []string
		pattern  string
		flags    map[string]int
		expected []string
	}{
		{
			name:     "SimpleMatch",
			lines:    lines,
			pattern:  "apple",
			flags:    map[string]int{},
			expected: []string{"apple banana cherry", "apple", "banana apple"},
		},
		{
			name:     "IgnoreCase",
			lines:    lines,
			pattern:  "APPLE",
			flags:    map[string]int{"ignore-case": 1},
			expected: []string{"apple banana cherry", "apple", "banana apple"},
		},
		{
			name:     "InvertMatch",
			lines:    lines,
			pattern:  "apple",
			flags:    map[string]int{"invert": 1},
			expected: []string{"cherry"},
		},
		{
			name:     "BeforeAndAfter",
			lines:    lines,
			pattern:  "apple",
			flags:    map[string]int{"before": 1, "after": 1},
			expected: []string{"apple banana cherry", "apple", "apple banana cherry", "apple", "banana apple", "apple", "banana apple", "cherry"},
		},
		{
			name:    "Context",
			lines:   lines,
			pattern: "apple",
			flags:   map[string]int{"context": 1},
			expected: []string{
				"apple banana cherry", "apple", "apple banana cherry", "apple", "banana apple", "apple", "banana apple", "cherry",
			},
		},
		{
			name:    "LineNumbers",
			lines:   lines,
			pattern: "apple",
			flags:   map[string]int{"line num": 1},
			expected: []string{
				"1:apple banana cherry",
				"2:apple",
				"3:banana apple",
			},
		},
		{
			name:    "FixedStringMatch",
			lines:   lines,
			pattern: "apple",
			flags:   map[string]int{"fixed": 1},
			expected: []string{
				"apple banana cherry",
				"apple",
				"banana apple",
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := Grep(testCase.lines, testCase.pattern, testCase.flags)

			t.Log(testCase.lines, result)

			assert.Equal(t, testCase.expected, result, fmt.Sprintf("Incorrect result for test case %s. Expected %s, got %s", testCase.name, testCase.expected, result))
		})
	}
}
