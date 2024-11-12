package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"
)

// Функция для вычисления хэша строки с использованием выбранного алгоритма
func calculateHash(input string, algorithm string) string {
	var hash []byte

	switch algorithm {
	case "MD5":
		hash = md5.New().Sum([]byte(input))
	case "SHA-256":
		hash = sha256.New().Sum([]byte(input))
	case "SHA-512":
		hash = sha512.New().Sum([]byte(input))
	default:
		fmt.Println("Неизвестный алгоритм хэширования.")
		return ""
	}

	// Возвращаем хэш в виде шестнадцатеричной строки
	return fmt.Sprintf("%x", hash)
}

// Функция для проверки целостности данных
func verifyIntegrity(input, providedHash, algorithm string) bool {
	// Убираем лишние пробелы и символы новой строки
	input = strings.TrimSpace(input)

	// Вычисляем хэш для входных данных
	computedHash := calculateHash(input, algorithm)
	return computedHash == providedHash
}

// Функция для печати ASCII-кодов каждого символа строки
func printAsciiCodes(input string) {
	for i, c := range input {
		fmt.Printf("Символ %d: '%c' (ASCII: %d)\n", i+1, c, c)
	}
}

func main() {
	// Чтение ввода от пользователя для выбора алгоритма
	fmt.Println("Выберите алгоритм хэширования:")
	fmt.Println("1. MD5")
	fmt.Println("2. SHA-256")
	fmt.Println("3. SHA-512")
	fmt.Print("Введите номер алгоритма (1-3): ")

	var choice int
	fmt.Scan(&choice)

	// Выбор алгоритма
	var algorithm string
	switch choice {
	case 1:
		algorithm = "MD5"
	case 2:
		algorithm = "SHA-256"
	case 3:
		algorithm = "SHA-512"
	default:
		fmt.Println("Неверный выбор. Программа завершена.")
		return
	}

	// Ввод строки для вычисления хэша
	fmt.Println("Введите строку для вычисления хэша:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Печатаем строку до вычисления хэша для отладки
	fmt.Printf("Строка для хэширования (до очистки): [%s]\n", input)
	printAsciiCodes(input) // Печатаем ASCII-коды каждого символа

	// Вычисление хэша
	hash := calculateHash(input, algorithm)
	fmt.Printf("Вычисленный хэш с использованием %s: %s\n", algorithm, hash)

	// Печать хэша, который можно вставить в проверку
	fmt.Println("\nВведите строку для проверки целостности данных:")

	// Ввод строки для проверки целостности данных
	fmt.Print("Введите строку для проверки: ")
	scanner.Scan()
	verifyInput := scanner.Text()

	// Печатаем строку для проверки до очистки
	fmt.Printf("Строка для проверки (до очистки): [%s]\n", verifyInput)
	printAsciiCodes(verifyInput) // Печатаем ASCII-коды каждого символа

	// Убираем лишние пробелы и символы новой строки
	verifyInput = strings.TrimSpace(verifyInput)

	// Печатаем строку для проверки после очистки
	fmt.Printf("Строка для проверки (после очистки): [%s]\n", verifyInput)
	printAsciiCodes(verifyInput) // Печатаем ASCII-коды после очистки

	// Вычисляем хэш для строки проверки
	computedHashForVerify := calculateHash(verifyInput, algorithm)
	fmt.Printf("Вычисленный хэш для строки проверки: %s\n", computedHashForVerify)

	// Запрос на ввод ожидаемого хэша для сравнения
	fmt.Print("Введите ожидаемый хэш для проверки: ")
	scanner.Scan()
	providedHash := scanner.Text()

	// Печатаем ожидаемый хэш для проверки
	fmt.Printf("Ожидаемый хэш для проверки: [%s]\n", providedHash)

	// Сравниваем хэши и выводим результат
	if verifyIntegrity(verifyInput, providedHash, algorithm) {
		fmt.Println("Целостность данных подтверждена: хэши совпадают.")
	} else {
		fmt.Println("Целостность данных нарушена: хэши не совпадают.")
	}
}
