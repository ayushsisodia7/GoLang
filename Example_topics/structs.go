package main

import "fmt"

type person struct { //This person struct type has name and age fields.
	name string
	age  int
}

func newPerson(name string) *person { //newPerson constructs a new person struct with the given name.

	p := person{name: name}
	p.age = 42
	return &p //You can safely return a pointer to local variable as a local variable will survive the scope of the function.
}

func main() {

	fmt.Println(person{"Bob", 20}) //This syntax creates a new struct.

	fmt.Println(person{name: "Alice", age: 30}) //You can name the fields when initializing a struct.

	fmt.Println(person{name: "Fred"}) //Omitted fields will be zero-valued.

	fmt.Println(&person{name: "Ann", age: 40}) //An & prefix yields a pointer to the struct.

	fmt.Println(newPerson("Jon")) //It’s idiomatic to encapsulate new struct creation in constructor functions

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) //Access struct fields with a dot.

	sp := &s //You can also use dots with struct pointers - the pointers are automatically dereferenced.
	fmt.Println(sp.age)

	sp.age = 51 //Structs are mutable.
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

}
