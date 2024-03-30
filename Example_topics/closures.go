package main

import "fmt"

func intSeq() func() int { //This function intSeq returns another function, which we define anonymously in the body of intSeq.
	//The returned function closes over the variable i to form a closure.
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq() //We call intSeq, assigning the result (a function) to nextInt.
	//This function value captures its own i value, which will be updated each time we call nextInt.

	fmt.Println(nextInt()) //See the effect of the closure by calling nextInt a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq() //To confirm that the state is unique to that particular function, create and test a new one.
	fmt.Println(newInts())
}
