package main

import (
	"fmt"  // пакет ввода и вывода данных
	"time" // пакет, для работы с датой
)

func main() {

	DateNow := time.Now() // Получение даты

	Year, Month, Day := DateNow.Date()        // Получаем у даты методом Date() - год, месяц, день
	Hour, Minutes, Seconds := DateNow.Clock() // Получаем у даты методом Clock() - час, минуты, секунды

	fmt.Printf("Today %d %s, %d year.\n", Day, Month, Year)       // Выводим дату
	fmt.Printf("Current time: %d:%d:%d.", Hour, Minutes, Seconds) // Выводим время
}
