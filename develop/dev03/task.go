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

func MakeStringUniq(rawString []string) []string {
	allKeys := make(map[string]bool)
	uniqString := make([]string, len(rawString))
	for _, item := range rawString {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			uniqString = append(uniqString, item)
		}
	}
	return uniqString
}

func SortNum(unsorted []string) []string {

	for i := range unsorted {
		var count1 int
		l := strings.Split(unsorted[i], " ")

		for j := range l {
			_, err := strconv.Atoi(l[j])
			if err == nil {
				count1++
			}
		}

		if count1 == len(l) {
			var result []int

			for j := range l {
				k, _ := strconv.Atoi(l[j])
				result = append(result, k)
			}
			sort.Ints(result)

			for m := range result {
				l = append(l[:m], strconv.Itoa(result[m]))
			}
		}
		unsorted[i] = strings.Join(l, " ")
	}
	return unsorted
}

func ReversSort(unsorted []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(unsorted)))
	return unsorted
}

func SortByColumn(data []string, colNum int) []string {
	sort.Slice(data, func(i, j int) bool {
		return data[i][colNum] < data[j][colNum]
	})
	return data
}

func main() {

	k := flag.Int("k", 0, "enter column")
	n := flag.Bool("n", false, "sort by num")
	r := flag.Bool("r", false, "revers sort")
	u := flag.Bool("d", false, "do not dublicate lines")

	dataFile := readFile("data.txt")

	switch {
	case *u:
		dataFile = MakeStringUniq(dataFile)
	case *k != 0:
		dataFile = SortByColumn(dataFile, *k)
	case *r == true:
		dataFile = ReversSort(dataFile)
	case *n == true:
		dataFile = SortNum(dataFile)
	}

	recordFile("outText.txt", dataFile)
}
