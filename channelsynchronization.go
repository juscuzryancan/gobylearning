package main

import (
	"fmt"
	"time"
)

/*
this is where the goroutine is ran
the done channel will notifiy another goroutine that this
function's work is done
*/
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	//send a val to notify that we're done
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	//block until we receive a notification from the worker on the channel
	//if this line were to be removed, the program would exit before the worker even started
	<-done
}
