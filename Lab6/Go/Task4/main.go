// Задание:
// Синхронизация с помощью мьютексов:
// Реализуйте программу, в которой несколько горутин увеличивают общую переменную-счётчик.
// Используйте мьютексы (sync.Mutex) для предотвращения гонки данных.
// Включите и выключите мьютексы, чтобы увидеть разницу в работе программы.

package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex // Мьютекс для синхронизации доступа к переменной-счётчику
var counter int

func main() {
	var wg sync.WaitGroup // Определяем WaitGroup для ожидания завершения всех горутин

	numGoroutines := 5
	numIterations := 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1) // Увеличиваем счётчик WaitGroup на 1 для каждой запущенной горутины
		go func() {
			defer wg.Done() // Уменьшаем счётчик WaitGroup на 1 после завершения горутины

			for j := 0; j < numIterations; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	wg.Wait() // Ожидаем завершения всех горутин

	fmt.Printf("Итоговое значение счётчика: %d\n", counter)
}
