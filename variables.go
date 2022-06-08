package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	//variables are initialized with zero-values
	var e int
	fmt.Println(e)

	//short hand syntax :=
	f := "apple"
	fmt.Println(f)
}
