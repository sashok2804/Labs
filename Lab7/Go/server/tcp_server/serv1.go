package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	serverPort := ":12345"                            // выбрали порт
	tcpListener, err := net.Listen("tcp", serverPort) // слушаем порт
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		os.Exit(1) // закрываем программу если ошибка
	}
	defer tcpListener.Close() // закрыть потом сервер

	fmt.Println("Сервер ждёт клиента на порту", serverPort)

	// Принимаем только одно соединение
	clientConn, acceptErr := tcpListener.Accept()
	if acceptErr != nil {
		fmt.Println("Ошибка при приеме соединения:", acceptErr)
		return // выходим, если произошла ошибка
	}
	defer clientConn.Close() // закрываем соединение после работы

	fmt.Println("Клиент подключился!")

	buffer := make([]byte, 1024) // выделяем место под сообщение клиента

	// читаем данные от клиента
	for {
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
}
