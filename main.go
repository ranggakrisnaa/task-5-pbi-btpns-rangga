package main

import (
	"github.com/gin-gonic/gin"
	db "github.com/ranggakrisnaa/task-5-pbi-btpns-rangga/database"
)

func main() {
	db.Init()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
