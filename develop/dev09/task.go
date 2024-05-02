package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/gocolly/colly"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {

	var urlFlag string
	var outputFlag string
	flag.StringVar(&urlFlag, "url", "", "URL для скачивания")
	flag.StringVar(&outputFlag, "output", "", "Каталог для сохранения файлов")
	flag.Parse()

	if urlFlag == "" {
		fmt.Println("Необходимо указать URL для скачивания")
		os.Exit(1)
	}
	if outputFlag == "" {
		fmt.Println("Необходимо указать каталог для сохранения файлов")
		os.Exit(1)
	}

	c := colly.NewCollector()

	c.SetRequestTimeout(10 * time.Second)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		absURL := e.Request.AbsoluteURL(link)
		go c.Visit(absURL)
	})

	c.OnResponse(func(r *colly.Response) {
		// Исключаем ошибки
		if r.StatusCode != http.StatusOK {
			log.Printf("Ошибка загрузки страницы: %s\n", r.Request.URL)
			return
		}

		u, err := url.Parse(r.Request.URL.String())
		if err != nil {
			log.Println("Ошибка при получении URL:", err)
			return
		}
		path := filepath.Join(outputFlag, u.Hostname(), u.Path)
		dir := filepath.Dir(path)

		// Создаем каталог, если его нет
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Println("Ошибка при создании каталога:", err)
			return
		}

		file, err := os.Create(path)
		if err != nil {
			log.Println("Ошибка при создании файла:", err)
			return
		}
		defer file.Close()

		_, err = file.Write(r.Body)
		if err != nil {
			log.Println("Ошибка при записи в файл:", err)
			return
		}

		log.Printf("Файл успешно сохранен: %s\n", path)
	})

	// Запускаем скачивание
	err := c.Visit(urlFlag)
	if err != nil {
		log.Println("Ошибка при загрузке URL:", err)
	}
}
