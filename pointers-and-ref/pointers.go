package main

import "fmt"

func main() {
	fmt.Println("Starting main function")
	fmt.Println("-------")
	test()
	fmt.Println("Ending main function")

}

func test() {
	// pointer
	var x = 10
	var y *int
	y = &x

	// p = pointer = hexadecimal
	fmt.Printf("x: %d y: %p and y type: %T\n", x, y, y)
}
