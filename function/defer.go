package main

import "fmt"

func main() {
	fmt.Println("Starting main function")
	multipleDefers()
	fmt.Println("Ending main function")
}
func multipleDefers() {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")
}

// Output:
// Third defer
// Second defer
// First defer
