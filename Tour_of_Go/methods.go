package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 { // Same as func Abs(v Vertex) float64
	//A method is a function with a special receiver argument
	//The receiver appears in its own argument list between the func keyword and the method name.
	//In this example, the Abs method has a receiver of type Vertex named v.
	//a method is just a function with a receiver argument.

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//You can only declare a method with a receiver whose type is defined in the same package as the method.
//You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

type MyFloat float64 //You can declare a method on non-struct types, too.

func (f MyFloat) Abs() float64 { //In this example we see a numeric type MyFloat with an Abs method.
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
