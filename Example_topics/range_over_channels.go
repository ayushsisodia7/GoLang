//We can also use this syntax to iterate over values received from a channel.

package main

import "fmt"

func main() {

	queue := make(chan string, 2) //We’ll iterate over 2 values in the queue channel.
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue { //This range iterates over each element as it’s received from queue.
		//Because we closed the channel above, the iteration terminates after receiving the 2 elements.
		fmt.Println(elem)
	}

	//This example also showed that it’s possible to close a non-empty channel but still have the remaining values be received.
}
