package main

import "fmt"

//has an int parameter
//args will be passed to it by value
//zeroval will get a copy of ival (distinct)
func zeroval(ival int) {
	ival = 0
}

//has an *int parameter
//the *iptr code dereferences the pointer from its memory adress to the current value at that address
//then assigning a value to a dereferenced pointer changes the value at the referenced address
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
