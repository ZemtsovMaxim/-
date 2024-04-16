package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Deduplicate(data []string) []string {
	allKeys := make(map[string]bool)
	var uniqData []string
	for _, item := range data {
		if !allKeys[item] {
			allKeys[item] = true
			uniqData = append(uniqData, item)
		}
	}
	return uniqData
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func SearchAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)

	for _, word := range words {
		key := sortString(word) // ключом будет отсортированное слово

		if anagramList, ok := anagrams[key]; ok {
			anagrams[key] = append(anagramList, word) // используем исходное слово
		} else {
			anagrams[key] = []string{word}
		}
	}

	for key, words := range anagrams {
		if len(words) == 1 {
			delete(anagrams, key)
		}
	}

	return anagrams
}

func main() {
	data := []string{
		"пятак", "пятка", "тяпка",
		"листок", "слиток", "столик",
		"кот", "ток", "окт", "кто",
	}

	// Приводим все слова к нижнему регистру перед удалением дубликатов
	for i, word := range data {
		data[i] = strings.ToLower(word)
	}

	newData := Deduplicate(data)

	anagramMap := SearchAnagrams(newData)

	for _, anagrams := range anagramMap {
		// Проверяем, что в множестве больше одного слова
		if len(anagrams) > 1 {
			fmt.Printf("Множество анаграмм для слова \"%s\": %s\n", anagrams[0], strings.Join(anagrams, ", "))
		}
	}
}
