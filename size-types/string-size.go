package main

import "fmt"

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))         // 13 (bytes)
	fmt.Println(len([]rune(s))) // 9 (actual characters)
}
