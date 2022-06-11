package main

import "fmt"

/*
	this is an example of a generic function
	mapkeys takes a map of any type and return a slice of its keys
	this function has two type parameters K and V
	K has the comparable constraint meaning we can compare values with the == and != operator
	V has the any constraint, meaning that it's not restricted in any way ( any is an alias for interface{} )
*/
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// as an example of a generic type, List is a singly-linked list with values of any type
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

//define methods just like how we do on regular types
//we must keep the type parameters in place
//List[T] not List
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	//when invoking generic functions we rely on type inference
	fmt.Println("keys m:", MapKeys(m))

	//but you can also specify them explicitly
	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
