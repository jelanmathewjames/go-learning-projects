package routes

import (
	"github.com/gin-gonic/gin"
)

func TodoRoutes(base *gin.RouterGroup) {
	todo := base.Group("/todo")
	todo.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "list",
		})
	})
	// todo.POST("/", create)
	// todo.GET("/", list)
	// todo.GET("/:id", get)
	// todo.PUT("/:id", update)
	// todo.DELETE("/:id", delete)
}