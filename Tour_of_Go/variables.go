package main

import "fmt"

var c, python, java bool //default int->0 bool->false float32->0 string->empty
//Variables declared without an explicit initial value are given their zero value.

func main() {
	//var i int
	//var j int
	var i, j int = 1, 2 //initialize variables

	//fmt.Println(i, j, c, python, java)

	k := 3 //Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	//Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

	c, python, java := true, false, "no!" //same as var c, python, java = true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
