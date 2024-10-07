package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func add[T constraints.Ordered](x T, y T) T {
	return x + y
}

func main() {
	value := add(5.5, 3.5)
	fmt.Println(value)
	value2 := add(5, 3)
	fmt.Println(value2)
}
