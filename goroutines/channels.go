package main

import "fmt"

// Buff sends two test messages to the provided channel and then closes it.
// This function is typically used to demonstrate channel communication in Go.
//
// Parameters:
//   - channel: A string channel to which messages will be sent
//
// The function sends "test" and "test2" messages to the channel and then closes it
// to signal that no more messages will be sent.
func Buff(channel chan string) {
	fmt.Println("Inside Buff")
	channel <- "test"
	channel <- "test2"
	close(channel)
}
