package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("----Traditional for loop")
	// Traditional for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}
	fmt.Println("----while loop with for")
	// Go doesn't have a while keyword, but you can use for
	count := 0
	for count < 3 {
		fmt.Printf("Count: %d\n", count)
		count++
	}

	fmt.Println("----infinite loop")
	counter := 0
	for {
		fmt.Printf("Running... %d\n", counter)
		counter++

		if counter >= 3 {
			break // Exit the loop
		}

		time.Sleep(1 * time.Second)
	}

	fmt.Println("----slice loop")
	fruits := []string{"apple", "banana", "cherry", "date"}

	// Loop with both index and value
	for index, fruit := range fruits {
		fmt.Printf("%d: %s\n", index, fruit)
	}

	// Loop with just the value (ignore index with _)
	for _, fruit := range fruits {
		fmt.Printf("Fruit: %s\n", fruit)
	}

	// Loop with just the index
	for index := range fruits {
		fmt.Printf("Index: %d\n", index)
	}

	fmt.Println("----map loop")
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Carol": 35,
	}

	// Loop through map
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// Loop through keys only
	for name := range ages {
		fmt.Printf("Name: %s\n", name)
	}

	fmt.Println("----string range loop")
	text := "Hello"

	// Loop through string characters
	for index, char := range text {
		fmt.Printf("Index %d: %c (Unicode: %d)\n", index, char, char)
	}

	fmt.Println("----nested loops")
	// Create a simple multiplication table
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println() // New line after each row
	}

	fmt.Println("----continue and breaks loops")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range numbers {
		if num%2 == 0 {
			continue // Skip even numbers
		}

		if num > 7 {
			break // Stop when we reach numbers greater than 7
		}

		fmt.Printf("Odd number: %d\n", num)
	}

	fmt.Println("----label loops")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	target := 5

OuterLoop:
	for i, row := range matrix {
		for j, val := range row {
			if val == target {
				fmt.Printf("Found %d at position [%d][%d]\n", target, i, j)
				break OuterLoop // Break out of both loops
			}
		}
	}
}
