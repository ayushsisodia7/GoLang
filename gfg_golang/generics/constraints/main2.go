package main

import "fmt"

type inttype int

func add[T ~int](x T, y T) T { // ~  operator tell go that any value under the hood of int is accepted.
	return x + y
}

func main() {
	ans := add(5, 3)
	fmt.Println(ans)
	x := inttype(5)
	y := inttype(11)
	ans2 := add(x, y)
	fmt.Println(ans2)
}
