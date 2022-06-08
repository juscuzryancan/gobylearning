package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	//range returns both the index, and value
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	//a go string is a read-only slice of bytes
	//a character is called a rune = integer that represents a Unicode code point
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
