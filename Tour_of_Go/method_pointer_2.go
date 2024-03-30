package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())    //methods with value receivers take either a value or a pointer as the receiver when they are called:
	fmt.Println(AbsFunc(v)) //Functions that take a value argument must take a value of that specific type

	p := &Vertex{4, 3}
	fmt.Println(p.Abs()) //In this case, the method call p.Abs() is interpreted as (*p).Abs().
	fmt.Println(AbsFunc(*p))
}
