package main

import (
	"fmt"
	"log"

	"github.com/albibenni/go-exercises/auth/helpers"
	"github.com/albibenni/go-exercises/auth/types"
	"github.com/joho/godotenv"
)

func test() {
	// Load environment variables
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set JWT key
	jwtSecret := types.JWTSecret.Value()
	fmt.Printf("JWT Secret from env: %s\n", jwtSecret)
	helpers.SetJWTKey(jwtSecret)

	// Generate a test token
	email := "test@example.com"
	userID := "test123"
	role := "ADMIN"

	token, _ := helpers.GenerateTokens(email, userID, role)
	fmt.Printf("Generated token: %s\n", token)

	// Try to validate the token
	claims, err := helpers.ValidateToken(token)
	if err != nil {
		fmt.Printf("Token validation failed: %v\n", err)
	} else {
		fmt.Printf("Token validation successful: %+v\n", claims)
	}
}
