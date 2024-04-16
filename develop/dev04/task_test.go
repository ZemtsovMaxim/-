package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeduplicate(t *testing.T) {
	testTable := []struct {
		data     []string
		expected []string
	}{
		{
			data:     []string{"пятак", "пятак", "пятак", "листок", "листок", "листок", "кот", "кот", "кот", "кот"},
			expected: []string{"пятак", "листок", "кот"},
		},
		{
			data:     []string{"a", "b", "c", "a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			data:     []string{"1", "2", "3", "4", "1", "2", "3", "4"},
			expected: []string{"1", "2", "3", "4"},
		},
	}

	for _, testCase := range testTable {
		result := Deduplicate(testCase.data)

		t.Log(testCase.data, result)

		assert.Equal(t, testCase.expected, result, fmt.Sprintf("Incorrect result. Expected %s, got %s", testCase.expected, result))
	}
}

func TestSearchAnagrams(t *testing.T) {
	testTable := []struct {
		data     []string
		expected map[string][]string
	}{
		{
			data: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "кот", "ток", "окт", "кто"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
				"кот":    {"кот", "ток", "окт", "кто"},
			},
		},
		{
			data: []string{"a", "ab", "ba", "abc", "bca", "cab"},
			expected: map[string][]string{
				"ab":  {"ab", "ba"},
				"abc": {"abc", "bca", "cab"},
			},
		},
		{
			data:     []string{"1", "2", "3", "4", "5"},
			expected: map[string][]string{},
		},
	}

	for _, testCase := range testTable {
		newData := Deduplicate(testCase.data)
		result := SearchAnagrams(&newData)

		t.Log(testCase.data, result)

		for _, anagrams := range *result {
			sort.Strings(*anagrams)
		}

		assert.Equal(t, testCase.expected, mapFromPtr(result), fmt.Sprintf("Incorrect result. Expected %v, got %v", testCase.expected, mapFromPtr(result)))
	}
}

func mapFromPtr(m *map[string]*[]string) map[string][]string {
	res := make(map[string][]string)
	for k, v := range *m {
		res[k] = *v
	}
	return res
}
