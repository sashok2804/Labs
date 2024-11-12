package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// Функция для шифрования строки с использованием AES
func encrypt(plainText, key string) (string, error) {
	// Преобразуем ключ в байты
	keyBytes := []byte(key)

	// Проверка длины ключа (должен быть 16, 24 или 32 байта)
	if len(keyBytes) != 16 && len(keyBytes) != 24 && len(keyBytes) != 32 {
		return "", fmt.Errorf("ключ должен быть длиной 16, 24 или 32 байта")
	}

	// Создаем блок-шифратор AES
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Создаем случайный вектор инициализации (IV) для режима CFB
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Шифруем текст с использованием режима CFB
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(plainText))

	// Возвращаем зашифрованный текст в виде строки base64
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// Функция для расшифровки строки с использованием AES
func decrypt(cipherText, key string) (string, error) {
	// Преобразуем ключ в байты
	keyBytes := []byte(key)

	// Проверка длины ключа (должен быть 16, 24 или 32 байта)
	if len(keyBytes) != 16 && len(keyBytes) != 24 && len(keyBytes) != 32 {
		return "", fmt.Errorf("ключ должен быть длиной 16, 24 или 32 байта")
	}

	// Декодируем строку из base64
	cipherTextBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	// Создаем блок-шифратор AES
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Вектор инициализации (IV) - это первые 16 байтов
	iv := cipherTextBytes[:aes.BlockSize]
	cipherTextBytes = cipherTextBytes[aes.BlockSize:]

	// Расшифровка
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherTextBytes, cipherTextBytes)

	// Возвращаем расшифрованную строку
	return string(cipherTextBytes), nil
}

func main() {
	// Ввод строки для шифрования
	fmt.Print("Введите строку для шифрования: ")
	var input string
	fmt.Scanln(&input)

	// Ввод секретного ключа
	fmt.Print("Введите секретный ключ (16, 24 или 32 байта): ")
	var key string
	fmt.Scanln(&key)

	// Шифрование строки
	encryptedText, err := encrypt(input, key)
	if err != nil {
		fmt.Println("Ошибка при шифровании:", err)
		return
	}

	// Печать зашифрованной строки в base64
	fmt.Println("Зашифрованный текст (base64):", encryptedText)

	// Ввод зашифрованного текста для расшифровки
	fmt.Print("\nВведите строку для расшифровки (base64): ")
	var encryptedInput string
	fmt.Scanln(&encryptedInput)

	// Расшифровка строки
	decryptedText, err := decrypt(encryptedInput, key)
	if err != nil {
		fmt.Println("Ошибка при расшифровке:", err)
		return
	}

	// Печать расшифрованной строки
	fmt.Println("Расшифрованный текст:", decryptedText)
}
