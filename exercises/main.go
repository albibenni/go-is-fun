package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("------ Exercises -------\n")

	startHere := time.Now()
	printArg()
	getElapsedTime(startHere)

	fmt.Printf("------ Exercises -------\n")

	startHere2 := time.Now()
	WithJoin()
	getElapsedTime(startHere2)
	fmt.Printf("------ Exercises -------\n")

	printArgByLines()

	fmt.Printf("------ Exercises -------\n")
}

func getElapsedTime(t time.Time) {
	fmt.Printf("%d elapsed\n", time.Since(t).Nanoseconds())
}
