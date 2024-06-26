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

func (v *Vertex) Scale(f float64) {

	//You can declare methods with pointer receivers.
	//This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)

	//For example, the Scale method here is defined on *Vertex.

	//Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

	//Try removing the * from the declaration of the Scale function and observe how the program's behavior changes.

	//With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.)
	//The Scale method must have a pointer receiver to change the Vertex value declared in the main function.

	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
