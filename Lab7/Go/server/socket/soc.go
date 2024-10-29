package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	serverAddress := "ws://localhost:12345/ws" // адрес сервера

	// Подключение к серверу
	conn, _, err := websocket.DefaultDialer.Dial(serverAddress, nil)
	if err != nil {
		fmt.Println("Не удалось подключиться к серверу:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Соединение успешно установлено!")

	// Чтение ввода пользователя в отдельной горутине
	go func() {
		for {
			_, response, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Ошибка при получении ответа:", err)
				return
			}
			fmt.Println("Сообщение от сервера:", string(response))
		}
	}()

	// Чтение ввода пользователя
	inputReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите сообщение (или 'exit' для выхода): ")
		message, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка при чтении ввода:", err)
			return
		}

		// Удаление символа новой строки
		message = message[:len(message)-1]

		if message == "exit" {
			break // Выход из программы
		}

		// Отправка сообщения в формате веб-сокетов
		err = conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			fmt.Println("Ошибка при отправке сообщения:", err)
			return
		}
	}

	fmt.Println("Завершение работы клиента.")
}
