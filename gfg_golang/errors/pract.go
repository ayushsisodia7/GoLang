package main

import "fmt"

type Rectangle struct {
	length  int
	breadth int
}

type Square struct {
	side int
}

type shape interface {
	Area() int
	Perimeter() int
}

func (s Square) Area() int {
	return s.side * s.side
}

func (s Rectangle) Area() int {
	return s.length * s.breadth
}

func (s Square) Perimeter() int {
	return 4 * s.side
}

func (s Rectangle) Perimeter() int {
	return 2 * (s.length + s.breadth)
}

// func Solve(a shape) {
//  fmt.Println("Area = ", a.Area())
//  fmt.Println("Perimeter = ", a.Perimeter())
// }

func Area(s shape) {
	fmt.Println(s.Area())
}

func Perimeter(s shape) {
	fmt.Println(s.Perimeter())
}

func main() {
	var r Rectangle
	var s Square
	r.length, r.breadth = 10, 20
	s.side = 40
	Area(r)
	Area(s)
	Perimeter(r)
	Perimeter(s)
}
