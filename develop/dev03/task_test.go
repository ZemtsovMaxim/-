package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeduplicate(t *testing.T) {
	testTable := []struct {
		str      []string
		expected []string
	}{
		{
			str:      []string{"1", "1", "1", "2"},
			expected: []string{"1", "2"},
		},
		{
			str:      []string{"1", "2", "3", "4"},
			expected: []string{"1", "2", "3", "4"},
		},
		{
			str:      []string{"a", "a", "a", "a"},
			expected: []string{"a"},
		},
	}

	for _, testCase := range testTable {
		result := Deduplicate(testCase.str)

		t.Log(testCase.str, result)

		assert.Equal(t, testCase.expected, result, fmt.Sprintf("Incorrect result. Expected %s, got %s", testCase.expected, result))
	}
}

func TestSort(t *testing.T) {
	testTable := []struct {
		str      []string
		expected []string
	}{
		{
			str:      []string{"1", "2", "4", "3"},
			expected: []string{"1", "2", "3", "4"},
		},
		{
			str:      []string{"1", "2", "3", "4"},
			expected: []string{"1", "2", "3", "4"},
		},
		{
			str:      []string{"c", "b", "a", "d"},
			expected: []string{"a", "b", "c", "d"},
		},
	}

	for _, testCase := range testTable {
		result := Sort(testCase.str)

		t.Log(testCase.str, result)

		assert.Equal(t, testCase.expected, result, fmt.Sprintf("Incorrect result. Expected %s, got %s", testCase.expected, result))
	}
}

func TestReversSort(t *testing.T) {
	testTable := []struct {
		str      []string
		expected []string
	}{
		{
			str:      []string{"1", "2", "4", "3"},
			expected: []string{"4", "3", "2", "1"},
		},
		{
			str:      []string{"1", "2", "3", "4"},
			expected: []string{"4", "3", "2", "1"},
		},
		{
			str:      []string{"c", "b", "a", "d"},
			expected: []string{"d", "c", "b", "a"},
		},
	}

	for _, testCase := range testTable {
		result := ReversSort(testCase.str)

		t.Log(testCase.str, result)

		assert.Equal(t, testCase.expected, result, fmt.Sprintf("Incorrect result. Expected %s, got %s", testCase.expected, result))
	}
}

func TestSortByColumn(t *testing.T) {
	testTable := []struct {
		str      []string
		expected []string
	}{
		{
			str:      []string{},
			expected: []string{},
		},
		{
			str: []string{
				"3 1 4",
				"1 2 3",
				"2 3 1",
				"4 2 3",
			},
			expected: []string{
				"1 2 3",
				"2 3 1",
				"3 1 4",
				"4 2 3",
			},
		},
	}

	for _, testCase := range testTable {
		result := SortByColumn(testCase.str, 0)

		t.Log(testCase.str, result)

		assert.Equal(t, testCase.expected, result, fmt.Sprintf("Incorrect result. Expected %s, got %s", testCase.expected, result))
	}
}
