package main

import (
	"flag" // пакет, для работы с флагами
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
	age := 0

	fmt.Printf("Введите ваше число: ")
	fmt.Scan(&age)

	if age%2 == 0 {
		fmt.Printf("%d - четное", age)
	} else {
		fmt.Printf("%d - нечетное", age)
	}

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
