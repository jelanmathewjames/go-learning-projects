package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.Engine) {
	base_api := app.Group("/api/v1")
	AuthRoutes(base_api)
	TodoRoutes(base_api)
}