package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "POST Opening",
	})
}

func DeleteOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Opening",
	})
}

func GetOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}

func UpdateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "PUT Opening",
	})
}

func ListOpeningsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "GET Openings",
	})
}
