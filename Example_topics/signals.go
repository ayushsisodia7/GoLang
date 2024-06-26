package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) //signal.Notify registers the given channel to receive notifications of the specified signals.

	done := make(chan bool, 1)

	go func() {

		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

	//When we run this program it will block waiting for a signal.
	//By typing ctrl-C (which the terminal shows as ^C) we can send a SIGINT signal, causing the program to print interrupt and then exit.
}
