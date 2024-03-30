//To wait for multiple goroutines to finish, we can use a wait group.

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) { //This is the function we’ll run in every goroutine
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second) //Sleep to simulate an expensive task.
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup //This WaitGroup is used to wait for all the goroutines launched here to finish.
	//Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer.

	for i := 1; i <= 5; i++ { //Launch several goroutines and increment the WaitGroup counter for each.
		wg.Add(1)

		i := i //Avoid re-use of the same i value in each goroutine closure

		go func() { //Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done.
			//This way the worker itself does not have to be aware of the concurrency primitives involved in its execution.

			//he first line of the function contains the statement defer wg.Done(). This means that the wg.Done() function will be called just before the worker function returns.
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait() //Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
	//Note that this approach has no straightforward way to propagate errors from workers.
}
