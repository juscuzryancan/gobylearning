package main

import "fmt"

type rect struct {
	width, height int
}

// methods are denoted by
// keyword func (receiver type) functionName() return type
func (r *rect) area() int {
	return r.width * r.height
}

//methods can be defined for either pointer or value receiver types
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	/*
		go handles conversion between values and pointers for method calls
		use a pointer to avoid coping on method calls or to allow the method to mutate the receiving struct
	*/
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
