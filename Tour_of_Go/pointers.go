package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} //An array's length is part of its type, so arrays cannot be resized
	fmt.Println(primes)

	var s []int = primes[1:4] //Slicing an array, works like python list slicing
	fmt.Println(s)
}
