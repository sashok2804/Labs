package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	adr := "localhost:12345" // адрес сервера

	fmt.Println("Попытка подключиться к серверу на", adr)
	conn, oshibka := net.Dial("tcp", adr) // пробуем подключиться
	if oshibka != nil {
		fmt.Println("Не удалось подключиться к серверу:", oshibka)
		os.Exit(1) // завершаем программу при ошибке
	}
	defer conn.Close() // закрываем соединение в конце работы

	fmt.Println("Подключение успешное! Введите сообщение для отправки:")
	vvod := bufio.NewReader(os.Stdin) // читаем ввод от пользователя

	soobshenie, oshibka2 := vvod.ReadString('\n') // читаем строку из консоли
	if oshibka2 != nil {
		fmt.Println("Ошибка при чтении ввода:", oshibka2)
		return // выход при ошибке
	}

	_, oshibka3 := conn.Write([]byte(soobshenie)) // отправляем сообщение серверу
	if oshibka3 != nil {
		fmt.Println("Ошибка при отправке сообщения:", oshibka3)
		return
	}

	otvet := make([]byte, 1024)     // создаем буфер для ответа от сервера
	n, oshibka4 := conn.Read(otvet) // читаем ответ от сервера
	if oshibka4 != nil {
		fmt.Println("Ошибка при получении ответа:", oshibka4)
		return
	}

	fmt.Println("Ответ от сервера:", string(otvet[:n])) // выводим ответ на экран

	// завершаем соединение
	fmt.Println("Соединение завершено, спасибо за использование :)")
}
