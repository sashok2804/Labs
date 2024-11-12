package main

import (
	"crypto/tls"
	"fmt"
	"os"
)

func main() {
	// Путь к сертификатам
	certFile := "server.crt"
	keyFile := "server.key"

	// Загружаем сертификат и ключ
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		fmt.Println("Ошибка при загрузке сертификата и ключа:", err)
		os.Exit(1)
	}

	// Создаем TLS-конфигурацию
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert}, // Указываем сертификат для сервера
	}

	// Запускаем сервер на порту 8443
	listener, err := tls.Listen("tcp", ":8443", tlsConfig)
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Сервер работает на порту 8443. Ожидаем подключений...")

	for {
		// Принимаем соединения от клиента
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при принятии соединения:", err)
			continue
		}

		// Приводим net.Conn к tls.Conn
		tlsConn, ok := conn.(*tls.Conn)
		if !ok {
			fmt.Println("Ошибка: соединение не является TLS-соединением")
			conn.Close()
			continue
		}

		// Обработка клиента
		go handleRequest(tlsConn)
	}
}

func handleRequest(conn *tls.Conn) {
	defer conn.Close()

	// Чтение данных
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Ошибка при чтении данных:", err)
		return
	}

	fmt.Printf("Получено сообщение от клиента: %s\n", string(buf[:n]))

	// Отправляем ответ
	_, err = conn.Write([]byte("Сообщение получено!"))
	if err != nil {
		fmt.Println("Ошибка при отправке ответа:", err)
	}
}
