package main

import (
	"fmt"
	"os"
)

func printArg() {
	for index := range len(os.Args) {
		fmt.Printf("Args: %v \n", os.Args[index])
	}

}
