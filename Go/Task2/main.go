package main

import "fmt"

func main() {

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
