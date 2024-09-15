package routes

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(base *gin.RouterGroup) {
	auth := base.Group("/auth")
	auth.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "login",
		})
	})
	// auth.POST("/login", login)
	// auth.POST("/register", register)
}