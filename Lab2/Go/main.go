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
	num := 0

	fmt.Printf("%d - %s", num, numPos(num))
}

func task3() {
	count := 11

	for i := 1; i < count; i++ {
		fmt.Printf("%d, ", i)
	}
}

func task4() {
	str := ""

	fmt.Scan(&str)

	fmt.Printf("Строка '%s' - %d симв.", str, strLen(str))
}

func task5() {
	rect := Rectangle{10, 20}
	rect.Square()
}

func task6() {
	num1 := 30
	num2 := 15

	fmt.Printf("fisrt - %d, second - %d, avg - %f", num1, num2, sred(num1, num2))
}

func numPos(n int) string {
	if n > 0 {
		return "positive"
	}
	if n < 0 {
		return "negative"
	}
	return "zero"
}

func strLen(s string) int {
	sum := 0

	for range s {
		sum++
	}

	return sum
}

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Square() int {
	s := r.Width * r.Height

	fmt.Printf("Ширина - %d, высота - %d, площадь - %d", r.Width, r.Height, s)

	return s
}

func sred(first int, second int) float64 {
	return (float64(first) + float64(second)) / 2
}
