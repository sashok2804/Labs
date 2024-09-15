package main

import (
	"flag" // пакет, для работы с флагами
	"fmt"  // пакет ввода и вывода данных
	"time" // пакет, для работы с датой
)

func main() {

	task := flag.Int("t", 1, "task num") // создаем флаг t, 1 - значение по умолчанию, "task num" - описание флага

	flag.Parse() // парсим флаги

	switch *task {
	case 1:
		task1()
	case 2:
		task2()
	case 3:
		task3()
	case 4:
		task4()
	case 5:
		task5()
	case 6:
		task6()
	default:
		task1()
	}
}

func task1() {
	DateNow := time.Now() // Получение даты

	Year, Month, Day := DateNow.Date()        // Получаем у даты методом Date() - год, месяц, день
	Hour, Minutes, Seconds := DateNow.Clock() // Получаем у даты методом Clock() - час, минуты, секунды

	fmt.Printf("Today %d %s, %d year.\n", Day, Month, Year)       // Выводим дату
	fmt.Printf("Current time: %d:%d:%d.", Hour, Minutes, Seconds) // Выводим время
}

func task2() {
	var i int = 100          // целое число
	var f32 float32 = 3.14   // дробное число
	var f64 float64 = 3.2314 // дробное число с большей точностью
	var b bool = true        // булево значение
	var s string = "Hello"   // строка

	fmt.Printf("%T - %d\n", i, i)       // вывод целого
	fmt.Printf("%T - %f\n", f32, f32)   // вывод дробного
	fmt.Printf("%T - %.2f\n", f64, f64) // вывод дробного округленный до двух знаков
	fmt.Printf("%T - %t\n", b, b)       // вывод булево
	fmt.Printf("%T - %s\n", s, s)       // вывод строки
}

func task3() {
	var y int = 44 // Полная запись обьявления
	x := 56        // Краткая запись обьявления

	fmt.Printf("X: %d, Y: %d, sum: %d\n", x, y, x+y)
}

func task4() {
	y := 33
	x := 44

	fmt.Printf("X: %d, Y: %d, sum: %d\n", x, y, x+y)
}

func task5() {
	x := 124.17684
	y := 432.48724
	fmt.Printf("X: %f, Y:%f\nSum: %f, dif: %f", x, y, plus(x, y), minus(x, y))
}

func task6() {
	arr := [3]int{1, 2, 97} // обьявляем массив с тремя элементами
	sum := 0

	fmt.Print(arr) // выводим массив

	for i := 0; i < len(arr); i++ { //суммируем все элементы массива
		sum += arr[i]
	}

	fmt.Printf(", AVG: %f", float64(sum)/float64(len(arr))) // приводим сумму и длину массива к дробной части и выводим среднее значение
}

// функция сложения двух чисел
func plus(a float64, b float64) float64 {
	return a + b
}

// функция вычитания двух чисел
func minus(a float64, b float64) float64 {
	return a - b
}
