package main

import (
	"cmp"
	"fmt"
)

func max[T cmp.Ordered](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

func Print[T any](s T) {
	fmt.Println(s)
}

func main() {
	fmt.Printf("Generics --- max 10 - 4: %d\n", max(10, 4))
	fmt.Printf("Generics --- max A - a: %s\n", max("A", "a"))
	Print("test generic print")
	Print(10)
	Print[int](10)
}
