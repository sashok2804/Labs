package main

import (
	"Go/fact" // пакет для факториала
	"Go/str"
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
	num := 0
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	factorial := fact.Factorial(num)
	fmt.Printf("Факториал числа %d: %d\n", num, factorial)
}

func task2() {
	s := ""
	fmt.Print("Введите строку: ")
	fmt.Scan(&s)

	revStr := str.Reverse(s)
	fmt.Printf("Перевернутая строка: %s", revStr)
}

func task3() {
	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Массив: ", arr)
}

func task4() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[0:3]

	fmt.Printf("Массив: %v, срез от 0 до 3 эл.: %v\n", arr, slice)

	slice = append(slice, 99)

	fmt.Printf("Добавлено число 99 в срез: %v\n", slice)

	slice = append(slice[:2], slice[3:]...)

	fmt.Printf("Удален элемент с индексом 2: %v\n", slice)

}

func task5() {
	strings := []string{"mama", "papas", "gojwoiawd", "nnnnnnnnnnnnnnnnnnnnnn"}
	longStr := ""

	fmt.Printf("Строки: %v\n", strings)

	for _, str := range strings {
		if len(str) > len(longStr) {
			longStr = str
		}
	}

	fmt.Printf("Длинная строка: %s", longStr)
}

func task6() {

}
