package main

import "fmt"

func main() {
	// make([]Type, length, capacity)
	s1 := make([]int, 5)       // Creates slice with length 5, capacity 5, zero values
	s2 := make([]int, 3, 10)   // Creates slice with length 3, capacity 10
	s3 := make([]string, 0, 5) // Empty slice with capacity 5
	fmt.Printf("Creates slice with length 5, capacity 5, zero values: %v\n", s1)
	fmt.Printf("Creates slice with length 3, capacity 10: %v\n", s2)
	s2 = append(s2, 10)
	fmt.Printf("Creates slice with length now of 4 with append, capacity 10: %v\n", s2)
	fmt.Printf("Creates slice with length now of %d with append, capacity -func cap() %d: %v\n", len(s2), cap(s2), s2)
	fmt.Printf("Empty slice with capacity 5: %v\n", s3)
	fmt.Printf("Empty slice with capacity 5: %v\n", s3)
}
