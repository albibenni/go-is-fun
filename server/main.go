package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	fmt.Printf("Listening on port %s\n", port)

}
