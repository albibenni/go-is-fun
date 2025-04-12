package main

import (
	"fmt"
)

func main() {
	// Example using MultiplyByTwoAndThree
	value := float32(10.1)
	double, triple := MultiplyByTwoAndThree(&value)
	fmt.Printf("Original value: %.1f\n", value)
	fmt.Printf("Double: %.1f, Triple: %.1f\n", double, triple)

	// Example using GenerateMultiple
	first, second, third := GenerateMultiple()
	fmt.Printf("First three multiples: %d, %d, %d\n", first, second, third)
}
