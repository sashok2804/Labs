package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wg sync.WaitGroup
var listener net.Listener
var clients = make(map[net.Conn]struct{})
var mu sync.Mutex

func handleConnection(conn net.Conn) {
	defer wg.Done()
	defer func() {
		mu.Lock()
		delete(clients, conn) // Удаляем соединение из списка активных
		mu.Unlock()
		conn.Close()
	}()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Клиент закрыл соединение.")
			} else {
				fmt.Println("Ошибка чтения:", err)
			}
			return
		}

		message := string(buf[:n])
		fmt.Println("Получено сообщение:", message)

		_, err = conn.Write([]byte("Сообщение получено"))
		if err != nil {
			fmt.Println("Ошибка отправки:", err)
			return
		}
	}
}

func shutdownServer() {
	fmt.Println("\nЗавершение работы сервера...")
	listener.Close()

	mu.Lock()
	for conn := range clients {
		conn.Close()
	}
	mu.Unlock()
}

func main() {
	var err error
	listener, err = net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Сервер запущен на порту 12345")

	// Обработка завершения работы сервера
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		shutdownServer()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			if opErr, ok := err.(*net.OpError); ok && opErr.Err.Error() == "use of closed network connection" {
				fmt.Println("Сервер закрыт, завершение работы.")
				break
			}
			fmt.Println("Ошибка при принятии соединения:", err)
			continue
		}

		mu.Lock()
		clients[conn] = struct{}{} // Добавляем соединение в список активных клиентов
		mu.Unlock()

		wg.Add(1)
		go handleConnection(conn)
	}

	wg.Wait()
}
