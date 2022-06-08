package main

import "fmt"

func main() {

	//single condition forloop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//classical
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//without conditionals it will repeatedly loop
	//the for loop is the while loop

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
