package main

import (
	"flag"
	"fmt"
	"math"
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
	default:
		task1()
	}
}
func task1() {
	p := Person{name: "Alice", age: 30}
	p.Info()

	p.Birthday()
	p.Info()
}

func task2() {
	c := Circle{radius: 5}
	fmt.Printf("Circle Area: %.2f\n", c.Area())
}

func task3() {
	c := Circle{radius: 5}
	r := Rectangle{width: 4, height: 5}

	shapes := []Shape{c, r}
	PrintAreas(shapes)
}
func task4() {
	b := Book{title: "Biblia", author: "IISYS"}
	fmt.Println(b.String())
}

type Person struct {
	name string
	age  int
}

func (p Person) Info() {
	fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)
}

func (p *Person) Birthday() {
	p.age++
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Shape interface {
	Area() float64
}

func PrintAreas(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}
}

type Book struct {
	title  string
	author string
}

type Stringer interface {
	String() string
}

func (b Book) String() string {
	return fmt.Sprintf("Title: %s, Author: %s", b.title, b.author)
}
