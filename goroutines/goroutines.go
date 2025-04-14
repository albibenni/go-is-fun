package main

import (
	"fmt"
	"time"
)

func Sleep() {
	for range 10 {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("Print after sleep of 400 ms\n")
	}
}

func SleepWithGoRoutines(channel chan string) {
	go runGoRoutinesAndSleep(channel)
}

func runGoRoutinesAndSleep(channel chan string) {
	for range 10 {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("Print after sleep of 400 ms\n")
	}
	channel <- "Done"
}
