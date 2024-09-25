package main

import (
	"Go/fact" // пакет для факториала
	"flag"    // пакет, для работы с флагами
	"fmt"
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
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	factorial := fact.Factorial(num)
	fmt.Printf("Факториал числа %d: %d\n", num, factorial)
}

func task2() {

}

func task3() {

}

func task4() {

}

func task5() {

}

func task6() {

}
