package routes

import (
	"github.com/albibenni/go-exercises/auth/config"
	"github.com/albibenni/go-exercises/auth/controller"
	"github.com/albibenni/go-exercises/auth/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *config.DB) {
	router.POST("/signup", controller.Signup(db))
	router.POST("/login", controller.Login(db))

	protected := router.Group("/")
	protected.Use(middleware.Authenticate())

	{
		protected.GET("/users", controller.GetUsers(db))
		protected.GET("/user/:id", controller.GetUser(db))
	}
}
