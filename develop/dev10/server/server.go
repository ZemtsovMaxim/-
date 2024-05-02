package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Новое соединение:", conn.RemoteAddr())

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Ошибка чтения данных:", err)
			return
		}
		if n == 0 {
			fmt.Println("Соединение закрыто клиентом:", conn.RemoteAddr())
			return
		}
		fmt.Printf("Получено от %s: %s", conn.RemoteAddr(), string(buf[:n]))
	}
}

func main() {
	// Запуск сервера на порту 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Сервер запущен на порту 8080")

	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при принятии соединения:", err)
			continue
		}

		go handleConnection(conn)
	}
}
