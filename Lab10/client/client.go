package main

import (
	"crypto/tls"
	"fmt"
	"os"
)

func main() {
	serverAddress := "localhost:8443" // Адрес сервера

	// Загружаем сертификат сервера (проверка доверенности)
	caCert, err := os.ReadFile("server.crt")
	if err != nil {
		fmt.Println("Ошибка при чтении сертификата CA:", err)
		os.Exit(1)
	}

	// Создаем пул сертификатов для проверки сервера
	certPool := tls.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)

	// Настроим TLS-конфигурацию для клиента, указывая сертификаты для проверки
	tlsConfig := &tls.Config{
		RootCAs: certPool, // Указываем сертификат для проверки сервера
	}

	// Подключаемся к серверу
	conn, err := tls.Dial("tcp", serverAddress, tlsConfig)
	if err != nil {
		fmt.Println("Не удалось подключиться к серверу:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Подключение установлено! Введите сообщение для отправки:")

	// Чтение сообщения от пользователя
	var message string
	fmt.Scanln(&message)

	// Отправка сообщения на сервер
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Ошибка при отправке сообщения:", err)
		return
	}

	// Чтение ответа от сервера
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Ошибка при получении ответа:", err)
		return
	}

	// Печать ответа
	fmt.Printf("Ответ от сервера: %s\n", string(buf[:n]))
}
