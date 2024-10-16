// Задание:
// Применение select для управления каналами:
// Создайте две горутины, одна из которых будет генерировать случайные числа,
// а другая — отправлять сообщения об их чётности/нечётности.
// Используйте конструкцию select для приёма данных из обоих каналов и вывода результатов в консоль.
// Продемонстрируйте, как select управляет многоканальными операциями.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Генерация случайных чисел
func generateNumbers(ch chan<- int) {
	for {
		num := rand.Intn(100)
		ch <- num
		time.Sleep(time.Second)
	}
}

// Проверка на чётность/нечётность
func checkEvenOdd(ch <-chan int, resultCh chan<- string) {
	for {
		num := <-ch
		if num%2 == 0 {
			resultCh <- fmt.Sprintf("Число %d чётное", num)
		} else {
			resultCh <- fmt.Sprintf("Число %d нечётное", num)
		}
	}
}

func main() {
	numCh := make(chan int)
	resultCh := make(chan string)

	go generateNumbers(numCh)
	go checkEvenOdd(numCh, resultCh)

	for {
		select {
		case result := <-resultCh:
			fmt.Println(result)
		}
	}
}
