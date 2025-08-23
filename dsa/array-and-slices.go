package main

import "fmt"

func main() {
	s := make([]int, 8)
	ar := [...]int{1, 2, 3, 4} // infer

	s = append(s, 10)

	fmt.Printf("type of slice s: %T, len: %d\n", s, len(s))
	fmt.Printf("type of array ar: %T\n", ar)
}
