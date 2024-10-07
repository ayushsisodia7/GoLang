package main

import "fmt"

func sub[T int | float64](x T, y T) T {
	return x - y
}

func main() {
	value := sub(5.5, 3.5)
	fmt.Println(value)
	value2 := sub(5, 3)
	fmt.Println(value2)
}
