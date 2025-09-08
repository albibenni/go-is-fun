package main

import (
	"log"
	"os"

	"github.com/albibenni/go-exercises/auth/routes"
	"github.com/albibenni/go-exercises/auth/types"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("local.env")

	Key:= GenerateRandomKey()
	SetJWTKey(key)

	r := gin.Default()
	routes.SetupRoutes(r)

	port := os.Getenv(types.ServerPath.Value())

	r.Run(":" + port)
	log.Println("Server is running on port: ", port)
}
