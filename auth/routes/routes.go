package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/signup", Signup())
	router.POST("/login", controllers.Login())

	protected := router.Group("/")
	protected.Use(middleware.Authenticate())
	{
		protected.GET("/users", controllers.Getusers())
		protected.GET("/user/:id", controllers.Getuser())
	}
}
