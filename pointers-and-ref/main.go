package main

import (
	"fmt"
)

func main() {

	var value float32 = 10.1
	r1, r2 := practice(&value)
	fmt.Printf("hello %v, %v \n", r1, r2)
}
