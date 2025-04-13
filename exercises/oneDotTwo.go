package main

import (
	"fmt"
	"os"
)

func printArgByLines() {
	args := os.Args
	var arMap = make(map[int]string)
	for key := range len(args) {
		fmt.Printf("index: %v, value: %v\n", key, args[key])
		arMap[key] = args[key]
	}
	for index, value := range arMap {
		fmt.Printf("map: index %v, value %v \n", index, value)
	}
}
