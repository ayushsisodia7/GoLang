package main

import "fmt"

func square(x int) int {
	return x * x
}

func main() {
	var list1 []int
	list1 = append(list1, 1, 2, 3, 4, 5)
	numbers := make(chan int, 60000000)
	for i := 0; i < 5; i++ {
		numbers <- list1[i]
	}
	for elem := range numbers {
		result := square(elem)
		fmt.Println(result)
	}
	close(numbers)
}

// func main() {
//     ss := make(chan string, 2)
//     ss <- "GFG"
//     ss <- "Golang channels"
//     ss <- "Sid from GFG"
//     fmt.Println(<-ss)
//     fmt.Println(<-ss)
//     fmt.Println(<-ss)
// }
