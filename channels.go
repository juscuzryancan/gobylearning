package main

import "fmt"

// channels are the pipes that connect concurrent goroutines
// you can send values into channels from one go routines and receive those values into another goroutine

func main() {

	//create a new channel with make(chan val-type)
	//channels are type by the values they convey
	messages := make(chan string)

	//send a value into a channel using the channel <- syntax
	//this send "ping" to the messages channel we made
	go func() { messages <- "ping" }()

	//the <-channel syntax receives a value from the channell
	msg := <-messages
	fmt.Println(msg)
	/*
		sends and receives block until both the sender and receiver are ready
		this allowed us to wait for the ping message without synchronizing
	*/
}
