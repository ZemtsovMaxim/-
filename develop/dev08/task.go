package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	for {

		fmt.Print("> ")

		var cmdString string
		fmt.Scanln(&cmdString)

		parts := strings.Fields(cmdString)
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "exit":
			fmt.Println("Выход из оболочки.")
			return
		case "cd":
			if len(parts) < 2 {
				fmt.Println("Необходимо указать директорию")
				continue
			}
			err := os.Chdir(parts[1])
			if err != nil {
				fmt.Println("Ошибка при смене директории:", err)
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("Ошибка при получении текущей директории:", err)
				continue
			}
			fmt.Println(dir)
		case "echo":
			if len(parts) < 2 {
				fmt.Println("")
				continue
			}
			fmt.Println(strings.Join(parts[1:], " "))
		case "kill":
			if len(parts) < 2 {
				fmt.Println("Необходимо указать PID процесса")
				continue
			}
			pid := parts[1]
			cmd := exec.Command("kill", "-9", pid)
			err := cmd.Run()
			if err != nil {
				fmt.Println("Ошибка при завершении процесса:", err)
			}
		case "ps":
			cmd := exec.Command("ps")
			output, err := cmd.Output()
			if err != nil {
				fmt.Println("Ошибка при выполнении команды ps:", err)
				continue
			}
			fmt.Println(string(output))
		default:
			cmd := exec.Command(parts[0], parts[1:]...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if err != nil {
				fmt.Println("Ошибка при выполнении команды:", err)
			}
		}
	}
}
