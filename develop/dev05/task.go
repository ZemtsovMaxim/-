package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Grep(lines []string, pattern string, flags map[string]int) []string {
	var result []string

	for i, line := range lines {
		match := strings.Contains(line, pattern)
		if flags["ignore-case"] == 1 {
			match = strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
		}

		if flags["invert"] == 1 {
			match = !match
		}

		if match {
			beforeLines := flags["before"]
			afterLines := flags["after"]
			contextLines := flags["context"]

			start := i - beforeLines
			if start < 0 {
				start = 0
			}

			end := i + afterLines + 1
			if end > len(lines) {
				end = len(lines)
			}

			if contextLines > 0 {
				start = i - contextLines
				if start < 0 {
					start = 0
				}
				end = i + contextLines + 1
				if end > len(lines) {
					end = len(lines)
				}
			}

			for j := start; j < end; j++ {
				lineNum := j + 1
				if flags["line num"] == 1 {
					result = append(result, fmt.Sprintf("%d:%s", lineNum, lines[j]))
				} else {
					result = append(result, lines[j])
				}
			}
		}
	}

	return result
}

func main() {
	after := flag.Int("A", 0, "Print +N lines after match")
	before := flag.Int("B", 0, "Print +N lines before match")
	context := flag.Int("C", 0, "Print ±N lines around match")
	count := flag.Bool("c", false, "Count matched lines")
	ignoreCase := flag.Bool("i", false, "Ignore case")
	invert := flag.Bool("v", false, "Invert match")
	fixed := flag.Bool("F", false, "Fixed string match")
	lineNum := flag.Bool("n", false, "Print line numbers")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: grep [flags] pattern file")
		os.Exit(1)
	}

	flags := map[string]int{
		"after":       *after,
		"before":      *before,
		"context":     *context,
		"count":       0,
		"ignore-case": 0,
		"invert":      0,
		"fixed":       0,
		"line num":    0,
	}

	if *count {
		flags["count"] = 1
	}

	if *ignoreCase {
		flags["ignore-case"] = 1
	}

	if *invert {
		flags["invert"] = 1
	}

	if *fixed {
		flags["fixed"] = 1
	}

	if *lineNum {
		flags["line num"] = 1
	}

	pattern := args[0]
	filename := args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	var result []string
	if *fixed {
		for _, line := range Grep(lines, pattern, flags) {
			result = append(result, line)
		}
	} else {
		for _, line := range Grep(lines, pattern, flags) {
			result = append(result, line)
		}
	}

	if *count {
		fmt.Println(len(result))
	} else {
		for _, line := range result {
			fmt.Println(line)
		}
	}
}
