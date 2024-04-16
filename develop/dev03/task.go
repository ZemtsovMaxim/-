package main

import (
	"bufio"
	"flag"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===
Отсортировать строки в файле по аналогии
с консольной утилитой sort (man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.
# Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительное

# Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func readFile(fileName string) []string {
	var result []string //Придумать как оптимизировать путем инициализации длины и cap. Чтобы не перевыделялась память каждый раз при append-e.

	file, err := os.Open(fileName)
	if err != nil {
		panic("Cant open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		result = append(result, str)
	}
	return result
}

func recordFile(fileName string, array []string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic("Cant create file")
	}
	defer file.Close()

	for i := 0; i < len(array); i++ {
		file.WriteString(array[i] + "\n")
	}
}

func Deduplicate(rawString []string) []string {
	allKeys := make(map[string]bool)
	var uniqString []string
	for _, item := range rawString {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			uniqString = append(uniqString, item)
		}
	}
	return uniqString
}

func Sort(unsorted []string) []string {
	sort.Strings(unsorted)
	return unsorted
}

func ReversSort(unsorted []string) []string {
	sorted := Sort(unsorted)
	last := len(sorted) - 1
	for i := 0; i < len(sorted)/2; i++ {
		sorted[i], sorted[last-i] = sorted[last-i], sorted[i]
	}
	return sorted
}

func SortByColumn(data []string, colNum int) []string {
	sort.Slice(data, func(i, j int) bool {
		fields1 := strings.Fields(data[i])
		fields2 := strings.Fields(data[j])

		if len(fields1) > colNum && len(fields2) > colNum {
			val1, err1 := strconv.Atoi(fields1[colNum])
			val2, err2 := strconv.Atoi(fields2[colNum])

			if err1 == nil && err2 == nil {
				return val1 < val2
			}
		}

		return fields1[colNum] < fields2[colNum]
	})

	return data
}

func main() {

	k := flag.Int("k", 0, "enter column")
	n := flag.Bool("n", false, "sort by num")
	r := flag.Bool("r", false, "revers sort")
	u := flag.Bool("u", false, "do not dublicate lines")

	flag.Parse()

	dataFile := readFile("data.txt")

	switch {
	case *u:
		dataFile = Deduplicate(dataFile)
	case *k != -1:
		dataFile = SortByColumn(dataFile, *k)
	case *r:
		dataFile = ReversSort(dataFile)
	case *n:
		dataFile = Sort(dataFile)
	}

	recordFile("result.txt", dataFile)
}
