package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup // WaitGroup для отслеживания горутин

// Структура для JSON-данных
type Data struct {
	Message string `json:"message"`
}

// Middleware для логирования
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // Запоминаем время начала обработки

		// Передаем управление следующему обработчику
		next.ServeHTTP(w, r)

		// Логируем метод, URL и время выполнения
		fmt.Printf("Метод: %s, URL: %s, Время: %s\n", r.Method, r.URL.Path, time.Since(start))
	})
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

	// Создаем новый маршрутизатор
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler) // маршрутизация GET /hello
	mux.HandleFunc("/data", dataHandler)   // маршрутизация POST /data

	// Оборачиваем маршрутизатор в middleware
	loggedMux := loggingMiddleware(mux)

	fmt.Println("Сервер запущен на порту", serverPort)

	go func() {
		err := http.ListenAndServe(serverPort, loggedMux) // запуск HTTP-сервера
		if err != nil {
			fmt.Println("Ошибка при запуске сервера:", err)
		}
	}()

	// Ожидание сигнала завершения
	<-sigs // Ждем сигнала остановки
	fmt.Println("Завершение работы сервера...")
}
