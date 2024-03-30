package main

import "os"

func main() {

	panic("a problem") //A common use of panic is to abort if a function returns an error value that we don’t know how to (or want to) handle.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

	//When first panic in main fires, the program exits without reaching the rest of the code.
	//If you’d like to see the program try to create a temp file, comment the first panic out.
}
