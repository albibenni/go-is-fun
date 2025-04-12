package main

import (
	"fmt"
)

func main() {

	var value float32 = 10.1
	r1, r2 := practice(&value)
	fmt.Printf("hello %v, %v \n", r1, r2)

	s1, s2, s3 := multiple()
	fmt.Printf("multiple return %v, %v, %v \n", s1, s2, s3)
}
