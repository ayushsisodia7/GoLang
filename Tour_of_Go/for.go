package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ { //syntax for init - condition - post
		sum += i
	}
	fmt.Println(sum)

	sum1 := 1
	c := 0
	//C's while is spelled for in Go.
	for sum1 < 1000 { //The init and post statements are optional. here only condition is kept in for loop.
		sum1 += sum1
		c++
	}
	fmt.Println(sum1, c)

	/*
		for {
			//infinite loop
		}
	*/
}
