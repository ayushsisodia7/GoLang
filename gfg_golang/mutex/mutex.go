package main

import (
	"fmt"
	"sync"
)

func writeLoop(m map[int]int, mux *sync.Mutex) {
	for j := 0; j < 200; j++ {
		for i := 0; i < 100; i++ {
			mux.Lock()
			m[i] = i
			mux.Unlock()
		}
	}
}

func readLoop(m map[int]int, mux *sync.Mutex) {
	for j := 0; j < 200; j++ {
		mux.Lock()
		for k, v := range m {
			fmt.Println(k, " - ", v)
		}
		mux.Unlock()
	}
}

func main() {
	m := map[int]int{}
	mux := &sync.Mutex{}
	go writeLoop(m, mux)
	go readLoop(m, mux)
	block := make(chan struct{})
	<-block
}
