package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) { //methods with pointer receivers take either a value or a pointer as the receiver when they are called:
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) { //functions with a pointer argument must take a pointer

	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{3, 4}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
