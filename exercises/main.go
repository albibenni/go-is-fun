package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("------ Exercises -------\n")

	startHere := time.Now()
	// printArg()
	// GetElapsedTime(startHere)
	//
	// fmt.Printf("------ Exercises -------\n")
	//
	// startHere2 := time.Now()
	// WithJoin()
	// fmt.Printf("------ Exercises -------\n")
	//
	// printArgByLines()
	Names()

	GetElapsedTime(startHere)
	fmt.Printf("------ Exercises -------\n")
}

func GetElapsedTime(t time.Time) {
	fmt.Printf("%d elapsed\n", time.Since(t).Nanoseconds())
}
