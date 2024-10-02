package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	people := make(map[string]int)
	addPerson(people, "Иван", 25)
	addPerson(people, "Анна", 30)
	addPerson(people, "Сергей", 35)

	fmt.Println("Список людей:")
	printPeople(people)
}

func task2() {
	people := make(map[string]int)
	addPerson(people, "Иван", 25)
	addPerson(people, "Анна", 30)
	addPerson(people, "Сергей", 35)

	fmt.Printf("Средний возраст: %.2f\n", averageAge(people))
}

func task3() {
	people := make(map[string]int)
	addPerson(people, "Иван", 25)
	addPerson(people, "Анна", 30)
	addPerson(people, "Сергей", 35)

	fmt.Println("Список до удаления Анны:")
	printPeople(people)

	removePerson(people, "Анна")
	fmt.Println("Список после удаления Анны:")
	printPeople(people)
}

func task4() {
	fmt.Println("Введите строку:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(strings.ToUpper(strings.TrimSpace(input)))
}

func task5() {
	fmt.Println("Введите числа через пробел:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	numStrings := strings.Split(input, " ")
	sum := 0

	for _, numStr := range numStrings {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			sum += num
		} else {
			fmt.Println("Ошибка преобразования числа:", numStr)
		}
	}
	fmt.Println("Сумма чисел:", sum)
}

func task6() {
	fmt.Println("Введите числа через пробел:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	numStrings := strings.Split(input, " ")
	var numbers []int

	for _, numStr := range numStrings {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			numbers = append(numbers, num)
		} else {
			fmt.Println("Ошибка преобразования числа:", numStr)
		}
	}

	fmt.Println("Массив в обратном порядке:")
	for i := len(numbers) - 1; i >= 0; i-- {
		fmt.Printf("%d ", numbers[i])
	}
	fmt.Println()
}

func addPerson(people map[string]int, name string, age int) {
	people[name] = age
}

func printPeople(people map[string]int) {
	for name, age := range people {
		fmt.Printf("%s: %d\n", name, age)
	}
}

func removePerson(people map[string]int, name string) {
	delete(people, name)
}

func averageAge(people map[string]int) float64 {
	if len(people) == 0 {
		return 0
	}
	totalAge := 0
	for _, age := range people {
		totalAge += age
	}
	return float64(totalAge) / float64(len(people))
}
