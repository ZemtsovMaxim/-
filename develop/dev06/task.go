package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
	=== Утилита cut ===

	Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

	Поддержать флаги:
	-f - "fields" - выбрать поля (колонки)
	-d - "delimiter" - использовать другой разделитель
	-s - "separated" - только строки с разделителем

	Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {

	var fields int
	var delimiter string
	var separated bool

	flag.IntVar(&fields, "f", 0, "выбрать поля (колонки)")
	flag.StringVar(&delimiter, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&separated, "s", false, "только строки с разделителем")

	flag.Parse()

	if fields <= 0 {
		log.Fatal("выберете корректный номер поля, номер должен быть положительным целым числом")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), delimiter)

		if separated && len(words) == 1 {
			continue
		}

		if len(words) < fields {
			fmt.Println()
		} else {
			fmt.Println(words[fields-1])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
