package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

// Генерация пары ключей RSA и сохранение в файлы
func generateKeys(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	// Генерация закрытого ключа
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, fmt.Errorf("ошибка при генерации ключа: %v", err)
	}

	// Получение открытого ключа из закрытого
	publicKey := &privateKey.PublicKey

	// Сохранение закрытого ключа в файл
	privateKeyFile, err := os.Create("private.pem")
	if err != nil {
		return nil, nil, fmt.Errorf("ошибка при создании файла private.pem: %v", err)
	}
	defer privateKeyFile.Close()

	// Сохранение открытого ключа в файл
	publicKeyFile, err := os.Create("public.pem")
	if err != nil {
		return nil, nil, fmt.Errorf("ошибка при создании файла public.pem: %v", err)
	}
	defer publicKeyFile.Close()

	// Сохраняем закрытый ключ
	privBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	err = pem.Encode(privateKeyFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privBytes,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("ошибка при сохранении закрытого ключа: %v", err)
	}

	// Сохраняем открытый ключ
	pubBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, fmt.Errorf("ошибка при сохранении открытого ключа: %v", err)
	}
	err = pem.Encode(publicKeyFile, &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubBytes,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("ошибка при сохранении открытого ключа: %v", err)
	}

	fmt.Println("Ключи успешно сохранены в файлы private.pem и public.pem")
	return privateKey, publicKey, nil
}

// Подписание сообщения с использованием закрытого ключа
func signMessage(privateKey *rsa.PrivateKey, message string) ([]byte, error) {
	// Хешируем сообщение с использованием SHA-256
	hash := sha256.New()
	hash.Write([]byte(message))
	hashed := hash.Sum(nil)

	// Подписываем хеш с использованием закрытого ключа
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return nil, fmt.Errorf("ошибка при подписании сообщения: %v", err)
	}

	return signature, nil
}

// Проверка подписи с использованием открытого ключа
func verifySignature(publicKey *rsa.PublicKey, message string, signature []byte) bool {
	// Хешируем сообщение с использованием SHA-256
	hash := sha256.New()
	hash.Write([]byte(message))
	hashed := hash.Sum(nil)

	// Проверяем подпись с использованием открытого ключа
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed, signature)
	if err != nil {
		fmt.Println("Ошибка проверки подписи:", err)
	}
	return err == nil
}

func main() {
	// Генерация ключей
	privateKey, publicKey, err := generateKeys(2048)
	if err != nil {
		log.Fatalf("Ошибка при генерации ключей: %v", err)
	}

	// Пример подписания сообщения
	message := "Hello, this is a secret message"
	signature, err := signMessage(privateKey, message)
	if err != nil {
		log.Fatalf("Ошибка при подписании сообщения: %v", err)
	}

	fmt.Println("Подписанное сообщение: ", message)
	fmt.Printf("Подпись: %x\n", signature)

	// Пример проверки подписи
	isValid := verifySignature(publicKey, message, signature)
	if isValid {
		fmt.Println("Подпись действительна.")
	} else {
		fmt.Println("Подпись недействительна.")
	}

	// Пример передачи подписанных сообщений
	fmt.Println("\n--- Симуляция передачи сообщения ---")
	// Здесь, как пример, мы передаем подписанное сообщение между двумя сторонами
	// (сторона 1 подписала, сторона 2 проверяет подпись)
	fmt.Println("Сторона 2 проверяет подпись:")
	if verifySignature(publicKey, message, signature) {
		fmt.Println("Сообщение успешно проверено.")
	} else {
		fmt.Println("Ошибка при проверке сообщения.")
	}
}
