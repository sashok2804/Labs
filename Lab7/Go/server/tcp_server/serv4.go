package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wg sync.WaitGroup // WaitGroup для отслеживания горутин

// Структура для JSON-данных
type Data struct {
	Message string `json:"message"`
}

// обработчик для GET-запроса на /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Привет! Это ваше приветственное сообщение."))
}

// обработчик для POST-запроса на /data
func dataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Ошибка при обработке JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("Полученные данные:", data.Message)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Данные успешно получены!"))
}

func main() {
	serverPort := ":12345" // выбрали порт

	// Механизм graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Запускаем HTTP-сервер
	http.HandleFunc("/hello", helloHandler) // маршрутизация GET /hello
	http.HandleFunc("/data", dataHandler)   // маршрутизация POST /data

	fmt.Println("Сервер запущен на порту", serverPort)

	go func() {
		err := http.ListenAndServe(serverPort, nil) // запуск HTTP-сервера
		if err != nil {
			fmt.Println("Ошибка при запуске сервера:", err)
		}
	}()

	// Ожидание сигнала завершения
	<-sigs // Ждем сигнала остановки
	fmt.Println("Завершение работы сервера...")
}
