package main

import "fmt"

/*
	by default channels are unbuffered, meaning they will only
	accept sends (chan <-) if there is a corresponding receive(<- chan)
	ready to receive the sent value. Buffered channels accept a limited
	number of values without a corresponding receiver for those values
*/

func main() {
	//channel of strings that buffer up to 2 values
	messages := make(chan string, 2)

	//since this channel is buffered
	//these values can be sent into the channel
	//without a corresponding concurrent receive
	messages <- "buffered"
	messages <- "channel"

	//this is how we receive them later
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
