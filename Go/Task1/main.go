package main

import "fmt"

func main() {

	fmt.Println("Введите первое число: ")
	var firstNum int
	fmt.Scanf("%d\n", &firstNum)

	fmt.Println("Введите второе число: ")
	var secondNum int
	fmt.Scanf("%d\n", &secondNum)

	fmt.Println("Введите знак операции:")
	var symbol string
	fmt.Scanf("%s\n", &symbol)

	fmt.Printf("Результат: %d\n", firstNum+secondNum)
}
