package main

import "fmt"

//structs are typed collections of fields
type person struct {
	name string
	age  int
}

// *person => the * is part of the type as this is saying that the variable is holding a pointer to a person
// &p => holds the appointed memory address that will be returned
// *p will be the value of what's inside of the memory address
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	/*
		you can safely return a pointer to local variable
		as a local variable will survive the scope of the function
	*/
	return &p
}

func main() {

	//creates a new struct
	fmt.Println(person{"Bob", 20})

	//you can name the fields while initializing
	fmt.Println(person{name: "Alice", age: 30})

	//omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})

	//& prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	//idiomatic to encapsulate new struct creation in constructer functions
	fmt.Println(newPerson("Jon"))

	//access struct fields with a dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	//you can also use dots with struct pointers - pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	//structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
}
