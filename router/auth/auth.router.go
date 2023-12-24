package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/ranggakrisnaa/task-5-pbi-btpns-rangga/controllers"
)

func AuthRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Your authentication logic here
		controllers.Login(c)
		controllers.Register(c)
	}
}
