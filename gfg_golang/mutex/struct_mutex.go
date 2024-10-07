package main

import "fmt"

type Rectangle struct {
	height int
	width  int
}

func (r Rectangle) Area() int { //value receiver function
	return r.width * r.height
}

// would receive a value and then multiple the height and width with the scale -> change the value of height, width itself
func (r *Rectangle) Scale(factor int) { //pointer receiver function
	r.width *= factor
	r.height *= factor
}

func main() {
	rec := Rectangle{width: 10, height: 5}
	fmt.Println(rec.Area())
	rec.Scale(10)
	fmt.Println(rec.Area())

}
