package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func max[T constraints.Ordered](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("Generics --- max 10 - 4: %d\n", max(10, 4))
}
