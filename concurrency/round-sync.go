package main

import "fmt"

func main() {
	done := make(chan bool)

	go func() {
		// do work
		fmt.Println("Doing work goroutine")
		done <- true
		fmt.Println("DONE TRUE")
	}()

	<-done // wait for goroutine to finish

	fmt.Println("AFTER DONE TRUE")
}
