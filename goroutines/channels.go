package main

import "fmt"

func Buff(channel chan string) {
	fmt.Println("Inside Buff")
	channel <- "test"
	channel <- "test2"
	close(channel)
}
