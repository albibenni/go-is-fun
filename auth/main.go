package main

import (
	"context"
	"log"

	"github.com/albibenni/go-exercises/auth/config"
	"github.com/albibenni/go-exercises/auth/helpers"
	"github.com/albibenni/go-exercises/auth/routes"
	"github.com/albibenni/go-exercises/auth/types"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("local.env")

	// Initialize database connection
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		if err := db.Close(context.Background()); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	jwtSecret := types.JWTSecret.Value()
	helpers.SetJWTKey(jwtSecret)

	r := gin.Default()
	routes.SetupRoutes(r, db)

	port := types.ServerPort.Value()

	r.Run(":" + port)
	log.Println("Server is running on port: ", port)
}
