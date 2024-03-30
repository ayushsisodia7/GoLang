package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{ //Map literals are like struct literals, but the keys are required.
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": { //If the top-level type is just a type name, you can omit it from the elements of the literal. So Vertex is not required for keys
		37.42202, -122.08408,
	},
}

func main() {
	/*m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}*/
	fmt.Println(m)

	//add or update element in map
	//add

	m["Microsoft"] = Vertex{43.56, 23.67} //add if key not present in map
	fmt.Println(m)
	m["Microsoft"] = Vertex{765.34, 65.24} //update if key already present in map
	fmt.Println(m)

	delete(m, "Microsoft") //delete key along with its value
	fmt.Println(m)

	v, ok := m["Google"] //here it checks if key present, if yes assignes value to first variable and assignes true value to second variable; if not assigns 0 to first variable and false to second variable
	fmt.Println("The value of v is :", v, "Present?", ok)
}
