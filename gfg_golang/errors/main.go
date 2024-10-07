// package main

// import "fmt"

// type englishbot struct {
// }

// type spanishbot struct {
// }

// func (e englishbot) getGreeting() string {
// 	return "Hi from Ayush"
// }

// func (s spanishbot) getGreeting() string {
// 	return "Hola from Ayush"
// }

// func greet(f spanishbot) {
// 	fmt.Println(f.getGreeting())
// }

// // func greet(f englishbot) {
// // 	fmt.Println(f.getGreeting())
// // }

// func main() {
// 	//var eb englishbot
// 	var sb spanishbot
// 	greet(sb)
// }

package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
}

type spanishBot struct {
}

func (e englishBot) getGreeting() string {
	return "Hi from Ayush"
}

func (e spanishBot) getGreeting() string {
	return "Hola from Ayush"
}

func greet(f bot) {
	fmt.Println(f.getGreeting())
}

func main() {
	var eb englishBot
	var sb spanishBot
	greet(eb)
	greet(sb)
}
