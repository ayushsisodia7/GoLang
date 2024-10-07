package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		mu     sync.RWMutex
		my_map = make(map[string]int)
	)

	go func() {
		mu.Lock()
		my_map["key1"] = 100
		my_map["key2"] = 200
		mu.Unlock()
	}()

	go func() {
		mu.Lock()
		fmt.Println(my_map["key1"])
		mu.Unlock()
	}()
}
