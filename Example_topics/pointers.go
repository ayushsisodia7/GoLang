package main

import "fmt"

func zeroval(ival int) { //zeroval will get a copy of ival distinct from the one in the calling function.
	ival = 0
}

func zeroptr(iptr *int) { //zeroptr in contrast has an *int parameter, meaning that it takes an int pointer.
	// The *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address.
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	//Assigning a value to a dereferenced pointer changes the value at the referenced address.
	fmt.Println("zeroptr:", i)

	//zeroval doesn’t change the i in main, but zeroptr does because it has a reference to the memory address for that variable.

	fmt.Println("pointer:", &i) //The &i syntax gives the memory address of i, i.e. a pointer to i.
}
