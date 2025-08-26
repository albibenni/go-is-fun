package main

import "fmt"

func main() {
	fmt.Println("Starting main function")
	fmt.Println("-------")
	multipleDefers()
	fmt.Println("-------")

	multipleDefersFor()
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
func multipleDefersFor() {
	for i := 0; i < 10; i++ {
		defer fmt.Printf("defer %d\n", i)
	}
}
