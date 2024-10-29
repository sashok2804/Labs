package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // Создаем новый экземпляр Upgrader

var clients = make(map[*websocket.Conn]bool) // Храним подключенные клиенты
var mu sync.Mutex                            // Мьютекс для безопасного доступа к clients

// Обработчик для WebSocket соединения
func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // Upgrade HTTP соединение до WebSocket
	if err != nil {
		fmt.Println("Ошибка при установке соединения:", err)
		return
	}
	defer conn.Close()

	mu.Lock()
	clients[conn] = true // Добавляем нового клиента
	mu.Unlock()

	fmt.Println("Клиент подключен")

	for {
		messageType, msg, err := conn.ReadMessage() // Читаем сообщение от клиента
		if err != nil {
			fmt.Println("Ошибка при чтении сообщения:", err)
			break
		}

		// Рассылаем сообщение всем подключенным клиентам
		mu.Lock()
		for client := range clients {
			if err := client.WriteMessage(messageType, msg); err != nil {
				fmt.Println("Ошибка при отправке сообщения:", err)
				client.Close()          // Закрываем соединение, если возникла ошибка
				delete(clients, client) // Удаляем клиента из списка
			}
		}
		mu.Unlock()
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
