// Задание:
// Напишите программу, которая параллельно выполняет три функции:
// - Расчёт факториала числа
// - Генерация случайных чисел
// - Вычисление суммы числового ряда
// Каждая функция должна выполняться в своей горутине.
// Используйте time.Sleep() для имитации задержек и продемонстрируйте параллельное выполнение.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция для расчёта факториала числа
func calculateFactorial(n int, resultChan chan<- int) {
	fmt.Printf("Расчёт факториала для числа %d...\n", n)
	time.Sleep(2 * time.Second) // Имитация задержки выполнения
	factorial := 1
	for i := 2; i <= n; i++ {
		factorial *= i // Вычисление факториала
	}
	resultChan <- factorial
}

// Функция для генерации случайных чисел
func generateRandomNumbers(count int, resultChan chan<- []int) {
	fmt.Println("Генерация случайных чисел...")
	time.Sleep(1 * time.Second)
	randomNumbers := make([]int, count)
	for i := 0; i < count; i++ {
		randomNumbers[i] = rand.Intn(100) // Генерация случайного числа от 0 до 99
	}
	resultChan <- randomNumbers
}

// Функция для вычисления суммы числового ряда
func calculateSum(n int, resultChan chan<- int) {
	fmt.Printf("Вычисление суммы числового ряда от 1 до %d...\n", n)
	time.Sleep(3 * time.Second)
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	resultChan <- sum
}

func main() {
	// Создание каналов для получения результатов от горутин
	factorialChan := make(chan int)
	randomNumbersChan := make(chan []int)
	sumChan := make(chan int)

	// Запуск горутин
	go calculateFactorial(5, factorialChan)
	go generateRandomNumbers(5, randomNumbersChan)
	go calculateSum(10, sumChan)

	// Ожидание и вывод результатов по мере их готовности
	for i := 0; i < 3; i++ {
		select {
		case factorialResult := <-factorialChan:
			fmt.Printf("Факториал: %d\n", factorialResult)
		case randomNumbers := <-randomNumbersChan:
			fmt.Printf("Случайные числа: %v\n", randomNumbers)
		case sumResult := <-sumChan:
			fmt.Printf("Сумма числового ряда: %d\n", sumResult)
		}
	}
}
