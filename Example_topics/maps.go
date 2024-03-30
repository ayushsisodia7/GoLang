package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int) //To create an empty map, use the builtin make: make(map[key-type]val-type).

	m["k1"] = 7 //Set key/value pairs using typical name[key] = val syntax.
	m["k2"] = 13

	fmt.Println("map:", m) //Printing a map with e.g. fmt.Println will show all of its key/value pairs.

	v1 := m["k1"] //Get a value for a key with name[key].
	fmt.Println("value1:", v1)

	v3 := m["k3"] //If the key doesn’t exist, the zero value of the value type is returned.
	fmt.Println("value3:", v3)

	fmt.Println("length of map:", len(m)) //The builtin len returns the number of key/value pairs when called on a map.

	delete(m, "k2") //The builtin delete removes key/value pairs from a map.
	fmt.Println("map after deletion:", m)

	clear(m) //To remove all key/value pairs from a map, use the clear builtin.
	fmt.Println("map after clear:", m)

	_, prs := m["k2"] //The optional second return value when getting a value from a map indicates if the key was present in the map.
	//This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	//Here we didn’t need the value of key itself, so we ignored it with the blank identifier _.
	fmt.Println("prs:", prs)

	m1 := make(map[string]int)
	m1["key1"] = 10
	m1["key2"] = 20
	p1, p2 := m1["key1"] //value of key, present status
	fmt.Println(p1, p2)  //this will print as (value of key,present(t/f))

	n := map[string]int{"foo": 1, "bar": 2} //You can also declare and initialize a new map in the same line with this syntax.
	fmt.Println("map:", n)                  //Note that maps appear in the form map[k:v k:v] when printed with fmt.Println.

	n2 := map[string]int{"foo": 1, "bar": 2} //The maps package contains a number of useful utility functions for maps.
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
