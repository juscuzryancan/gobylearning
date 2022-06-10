package main

import "fmt"

/*
	go supports embedding of structs and interfaces to create a composition of types
*/
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

//container embeds base, the field does not need a name
type container struct {
	base
	str string
}

func main() {
	//when creating structs withs literals we initalize the embedding explicitly
	// here the embedded type serves as the field name
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	//we can access the base's fields directly on co
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	//we can also spell out the full path with embedded type name
	fmt.Println("also num:", co.base.num)

	//since container embeds base, the base also has the methods of a container
	//this invokes a method from base directly on co
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	//this works because go is a value oriented language
	//we instantiate d as a describer and use co's value to be the value of the object
	//this in turn allows d to call the describe method
	var d describer = co
	fmt.Println("describer:", d.describe())
}
