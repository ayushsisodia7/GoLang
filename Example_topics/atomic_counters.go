//Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops atomic.Uint64 //We’ll use an atomic integer type to represent our (always-positive) counter.

	var wg sync.WaitGroup //A WaitGroup will help us wait for all goroutines to finish their work.

	for i := 0; i < 50; i++ { //We’ll start 50 goroutines that each increment the counter exactly 1000 times.
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {

				ops.Add(1) //To atomically increment the counter we use Add.
			}

			wg.Done()
			fmt.Println("ops:", ops.Load())
		}()
	}

	wg.Wait() //Wait until all the goroutines are done.

	fmt.Println("ops:", ops.Load()) //Here no goroutines are writing to ‘ops’, but using Load it’s safe to atomically read a value even while other goroutines are (atomically) updating it.
}
