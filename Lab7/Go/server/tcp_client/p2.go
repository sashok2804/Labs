package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	serverAddress := "localhost:12345"
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println("Ошибка подключения к серверу:", err)
		return
	}
	defer conn.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	// Фоновая горутина для проверки соединения
	go func() {
		defer wg.Done()
		buf := make([]byte, 1) // Используется для проверки соединения
		for {
			time.Sleep(2 * time.Second) // Пауза между проверками
			_, err := conn.Read(buf)
			if err != nil {
				fmt.Println("Соединение потеряно. Закрытие клиента...")
				os.Exit(1) // Завершение программы при потере соединения
			}
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите сообщение (или 'exit' для выхода): ")
		message, _ := reader.ReadString('\n')

		// Проверка на выход из программы
		if message == "exit\n" {
			fmt.Println("Выход из клиента...")
			break
		}

		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Ошибка отправки сообщения:", err)
			break
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Ошибка чтения ответа:", err)
			break
		}

		fmt.Println("Ответ сервера:", string(buf[:n]))
	}

	wg.Wait() // Ждем завершения фоновой горутины
}
