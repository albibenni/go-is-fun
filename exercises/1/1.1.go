package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := range os.Args {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Printf("%s\n", s)
}
