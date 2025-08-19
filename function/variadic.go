package main

import "fmt"

func main() {
	fmt.Printf("------\n")

	fmt.Printf("sum - %d\n", sum(1, 2, 5, 10))
}

func sum(numbers ...int) (total int) {
	total = 0
	for _, num := range numbers {
		total += num
	}
	return
}
