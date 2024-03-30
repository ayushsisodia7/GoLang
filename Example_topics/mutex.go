package main

import (
	"fmt"
	"sync"
)

type Container struct {
	//Container holds a map of counters;
	//since we want to update it concurrently from multiple goroutines, we add a Mutex to synchronize access.
	//Note that mutexes must not be copied, so if this struct is passed around, it should be done by pointer.
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {

	c.mu.Lock() //Lock the mutex before accessing counters; unlock it at the end of the function using a defer statement.
	c.counters[name]++
	defer c.mu.Unlock()
}

func main() {
	c := Container{ //Note that the zero value of a mutex is usable as-is, so no initialization is required here.

		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) { //This function increments a named counter in a loop.
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000) //Run several goroutines concurrently; note that they all access the same Container, and two of them access the same counter.
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait() //Wait for the goroutines to finish
	fmt.Println(c.counters)

	//Running the program shows that the counters updated as expected.
}
