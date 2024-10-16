// Задание
// Разработка многопоточного калькулятора:
// Напишите многопоточный калькулятор, который одновременно может обрабатывать запросы на выполнение простых операций (+, -, *, /).
// Используйте каналы для отправки запросов и возврата результатов.
// Организуйте взаимодействие между клиентскими запросами и серверной частью калькулятора с помощью горутин.

package main

import (
	"fmt"
)

// Описание структуры запроса для калькулятора
type CalcRequest struct {
	Operand1  float64
	Operand2  float64
	Operation string
	Result    chan float64 // Канал для возврата результата
	Error     chan error   // Канал для возврата ошибки
}

// Функция-калькулятор, обрабатывающая запросы в отдельной горутине
func calculator(requests chan CalcRequest) {
	for req := range requests {
		var result float64
		var err error

		// Обработка операций
		switch req.Operation {
		case "+":
			result = req.Operand1 + req.Operand2
		case "-":
			result = req.Operand1 - req.Operand2
		case "*":
			result = req.Operand1 * req.Operand2
		case "/":
			if req.Operand2 != 0 {
				result = req.Operand1 / req.Operand2
			} else {
				err = fmt.Errorf("деление на ноль")
			}
		default:
			err = fmt.Errorf("неизвестная операция")
		}

		if err != nil {
			req.Error <- err
		} else {
			req.Result <- result
		}
	}
}

// Функция для отправки запроса и получения результата
func sendRequest(operand1, operand2 float64, operation string, requests chan CalcRequest) {

	result := make(chan float64)
	err := make(chan error)

	requests <- CalcRequest{
		Operand1:  operand1,
		Operand2:  operand2,
		Operation: operation,
		Result:    result,
		Error:     err,
	}

	select {
	case res := <-result:
		fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", operand1, operation, operand2, res)
	case e := <-err:
		fmt.Printf("Ошибка: %s\n", e.Error())
	}
}

func main() {
	requests := make(chan CalcRequest)

	go calculator(requests)

	sendRequest(5, 3, "+", requests)
	sendRequest(7, 2, "-", requests)
	sendRequest(6, 3, "*", requests)
	sendRequest(10, 0, "/", requests)
	sendRequest(9, 3, "/", requests)

	var input string
	fmt.Scanln(&input)
}
