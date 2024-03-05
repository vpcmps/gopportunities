package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Code to initialize the router
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	router.Run()
}
