package routes

import (
	"github.com/albibenni/go-exercises/auth/middleware"
	"github.com/albibenni/go-exercises/auth/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/signup", services.Signup())
	router.POST("/login", services.Login())

	protected := router.Group("/")
	protected.Use(middleware.Authenticate())

	{
		protected.GET("/users", services.Getusers())
		protected.GET("/user/:id", services.Getuser())
	}
}
