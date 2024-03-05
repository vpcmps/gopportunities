package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Code to initialize the router
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}