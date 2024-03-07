package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vpcmps/gopportunities/docs"
	"github.com/vpcmps/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	docs.SwaggerInfo.Title = "Gopportunities"
	docs.SwaggerInfo.Description = "Gopportunities API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.GET("/opening", handler.GetOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
