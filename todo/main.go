package main

import (
	"github.com/gin-gonic/gin"
	"todo/routes"
)

func main() {
	app := gin.Default()
	routes.SetupRouter(app)
	app.Run(":8003")
}