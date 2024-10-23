package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // Создаем новый экземпляр Upgrader

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // Upgrade HTTP соединение до WebSocket
	if err != nil {
		fmt.Println("Ошибка при установке соединения:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Клиент подключен")

	for {
		messageType, msg, err := conn.ReadMessage() // Читаем сообщение от клиента
		if err != nil {
			fmt.Println("Ошибка при чтении сообщения:", err)
			break
		}

		// Проверяем, если сообщение "hello"
		if string(msg) == "123" {
			err = conn.WriteMessage(messageType, []byte("hi")) // Отправляем "hi"
		} else {
			err = conn.WriteMessage(messageType, msg) // Отправляем обратно то же сообщение
		}
		if err != nil {
			fmt.Println("Ошибка при отправке сообщения:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnection) // Устанавливаем обработчик для /ws
	fmt.Println("Сервер запущен на порту 12345")
	err := http.ListenAndServe(":12345", nil) // Запускаем HTTP-сервер
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
