package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wg sync.WaitGroup // WaitGroup для отслеживания горутин

// функция main начнём тут
func main() {
	serverPort := ":12345"                            // выбрали порт
	tcpListener, err := net.Listen("tcp", serverPort) // слушаем порт
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		os.Exit(1) // закрываем программу если ошибка
	}
	defer tcpListener.Close() // закрыть потом сервер

	fmt.Println("Сервер ждёт клиентов на порту", serverPort)

	// Механизм graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs // Ждем сигнала остановки
		fmt.Println("Завершение работы сервера...")
		tcpListener.Close() // Закрываем слушатель
	}()

	for { // бесконечный цикл чтобы принимать соединения
		clientConn, acceptErr := tcpListener.Accept()
		if acceptErr != nil {
			fmt.Println("Ошибка при приеме соединения:", acceptErr)
			continue // пропускаем, пробуем следующее соединение
		}

		wg.Add(1)                             // Увеличиваем счетчик горутин
		go handleClientConnection(clientConn) // запускаем в горутине
	}

	wg.Wait() // Ждем завершения всех горутин
	fmt.Println("Все соединения закрыты. Сервер остановлен.")
}

// handleClientConnection обрабатывает каждый клиент
func handleClientConnection(clientConn net.Conn) {
	defer wg.Done()          // Уменьшаем счетчик горутин
	defer clientConn.Close() // закрываем соединение после работы

	buffer := make([]byte, 1024) // выделяем место под сообщение клиента

	// читаем данные от клиента
	n, readErr := clientConn.Read(buffer)
	if readErr != nil {
		fmt.Println("Ошибка чтения данных:", readErr)
		return // если ошибка, выходим
	}

	// просто пишем сообщение на экран сервера
	message := string(buffer[:n])
	fmt.Println("Получено сообщение от клиента:", message)

	// шлем клиенту подтверждение
	confirmation := "Сообщение получено! Спасибо :3"
	_, writeErr := clientConn.Write([]byte(confirmation))
	if writeErr != nil {
		fmt.Println("Ошибка при отправке подтверждения:", writeErr)
	}
}
