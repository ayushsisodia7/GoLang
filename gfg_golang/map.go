package main

import "fmt"

func main() {
	my_map := map[string]int{
		"key1": 1,
		"key2": 2,
	}

	//add new key value pairs outside map defination
	my_map["key3"] = 10
	my_map["key4"] = 20

	//iterate through a map
	for key, value := range my_map {
		fmt.Println("key as --", key, " and value as -- ", value)
	}

	//find value of a key
	value := my_map["key1"]
	fmt.Println("value of the key is ", value)

	//check if a key is present or not
	value1, foundOrNot := my_map["key5"]
	if foundOrNot {
		fmt.Println("value with the given key is found as ", value1)
	} else {
		fmt.Println("key does not exist")
	}

	//print length of map
	fmt.Println("length of the map is ", len(my_map))
}
