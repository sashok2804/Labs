package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // Создаем новый экземпляр Upgrader

// Middleware для логирования
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		fmt.Printf("Метод: %s, URL: %s, Время выполнения: %v\n", r.Method, r.URL, duration)
	})
}

// Обработчик для WebSocket соединения
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
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

// Обработчик для другого маршрута (пример)
func handleAnotherRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Это другой маршрут!")
}

func main() {
	http.Handle("/ws", loggingMiddleware(http.HandlerFunc(handleWebSocket)))         // Устанавливаем обработчик для /ws
	http.Handle("/another", loggingMiddleware(http.HandlerFunc(handleAnotherRoute))) // Другой маршрут

	fmt.Println("Сервер запущен на порту 12345")
	err := http.ListenAndServe(":12345", nil) // Запускаем HTTP-сервер
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}

//curl http://localhost:12345/another
