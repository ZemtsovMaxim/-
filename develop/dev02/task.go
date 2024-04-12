package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func String_unpacking(s string) (str string, err error) {
	if s == "" {
		return "", nil
	}
	if s[0] < 97 || s[0] > 122 {
		return str, fmt.Errorf("некорректная строка")
	}

	for i := 0; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) {
			str += string(s[i])
		} else if s[i] == '\\' {
			str += string(s[i+1])
			i++
		} else {
			bytes := make([]byte, 0)
			for n := i; n < len(s); n++ {
				if unicode.IsNumber(rune(s[n])) {
					bytes = append(bytes, s[n])
				} else {
					break
				}
			}
			atoi, err := strconv.Atoi(string(bytes))
			if err != nil {
				panic(err)
			}
			y := strings.Repeat(string(s[i-1]), atoi-1)
			str += y
		}
	}

	return
}

func main() {
	tests := []string{"a4bc2d5e", "abcd", "45", "", `qwe\4\5`, `qwe\45`, `qwe\\5`}

	for _, test := range tests {
		unpacked, err := String_unpacking(test)
		if err != nil {
			fmt.Printf("Ошибка при распаковке строки '%s': %s\n", test, err)
		} else {
			fmt.Printf("Распакованная строка '%s': '%s'\n", test, unpacked)
		}
	}
}
