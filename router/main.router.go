package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ranggakrisnaa/task-5-pbi-btpns-rangga/controllers"
	"github.com/ranggakrisnaa/task-5-pbi-btpns-rangga/middlewares"
)

func Router() *gin.Engine {
	router := gin.Default()

	appAuth := router.Group("/api/v1/auth")
	{
		appAuth.POST("/register", controllers.Register)
		appAuth.POST("/login", controllers.Login)
	}

	appMain := router.Group("/api/v1")
	{
		appMain.Use(middlewares.AuthMiddleware())
		appMain.PUT("/user/:id", controllers.UpdateUser)
		appMain.DELETE("/user/:id", controllers.DeleteUser)
	}

	return router
}
