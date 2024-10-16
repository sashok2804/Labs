package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// Функция для реверсирования строки
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Воркер, который будет обрабатывать задачи
func worker(id int, tasks <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task: %s\n", id, task)
		reversed := reverseString(task)
		results <- reversed
	}
}

func main() {
	// Открываем файл с данными
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Читаем строки из файла
	scanner := bufio.NewScanner(file)
	var tasks []string
	for scanner.Scan() {
		tasks = append(tasks, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var workerCount int
	fmt.Print("Enter number of workers: ")
	fmt.Scan(&workerCount)

	taskChan := make(chan string, len(tasks))   // Канал для задач
	resultChan := make(chan string, len(tasks)) // Канал для результатов

	var wg sync.WaitGroup

	// Запускаем воркеров
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, taskChan, resultChan, &wg)
	}

	for _, task := range tasks {
		taskChan <- task
	}

	close(taskChan)

	wg.Wait()

	close(resultChan)

	// Выводим результаты работы воркеров
	fmt.Println("Reversed lines:")
	for result := range resultChan {
		fmt.Println(result)
	}
}
