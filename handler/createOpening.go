package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := struct {
		role string
	}{}
	ctx.BindJSON(&request)

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST Opening",
	})
}
