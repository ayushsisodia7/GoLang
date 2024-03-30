package main

import "fmt"

func add(x, y int) int { //x and y are of same data type
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
