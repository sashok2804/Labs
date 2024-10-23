package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	serverAddress := "localhost:12345" // адрес сервера

	fmt.Println("Попытка подключиться к серверу на", serverAddress)
	connection, connectionErr := net.Dial("tcp", serverAddress) // пробуем подключиться
	if connectionErr != nil {
		fmt.Println("Не удалось подключиться к серверу:", connectionErr)
		os.Exit(1) // завершаем программу при ошибке
	}
	defer connection.Close() // закрываем соединение в конце работы

	fmt.Println("Подключение успешное! Введите сообщение для отправки:")
	inputReader := bufio.NewReader(os.Stdin) // читаем ввод от пользователя

	message, readErr := inputReader.ReadString('\n') // читаем строку из консоли
	if readErr != nil {
		fmt.Println("Ошибка при чтении ввода:", readErr)
		return // выход при ошибке
	}

	_, writeErr := connection.Write([]byte(message)) // отправляем сообщение серверу
	if writeErr != nil {
		fmt.Println("Ошибка при отправке сообщения:", writeErr)
		return
	}

	responseBuffer := make([]byte, 1024)                  // создаем буфер для ответа от сервера
	n, readResponseErr := connection.Read(responseBuffer) // читаем ответ от сервера
	if readResponseErr != nil {
		fmt.Println("Ошибка при получении ответа:", readResponseErr)
		return
	}

	fmt.Println("Ответ от сервера:", string(responseBuffer[:n])) // выводим ответ на экран

	// завершаем соединение
	fmt.Println("Соединение завершено, спасибо за использование :)")
}
