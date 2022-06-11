package main

import (
	"fmt"
	"time"
)

//a goroutine is a lightweigth thread of execution

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	//synchronous call
	//the blocking call, blocked because we are stuck at this line
	f("direct")

	//a new go routine will execute concurrently with the calling one
	go f("goroutine")

	//you can also start a go routine for an anon funciton
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//the two functions are now asynchronously finishing
	time.Sleep(time.Second)
	fmt.Println("done")
}

/*
we see the output of the blocking call
then the output of the goroutines may be interleaved
*/
