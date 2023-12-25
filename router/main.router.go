package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ranggakrisnaa/task-5-pbi-btpns-rangga/controllers"
)

func Router() *gin.Engine {
	router := gin.Default()

	app := router.Group("/api/v1/auth")
	{
		app.POST("/register", controllers.Register)
		app.POST("/login", controllers.Login)
	}

	return router
}
