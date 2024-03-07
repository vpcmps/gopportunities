package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, status int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, gin.H{
		"message":     msg,
		"errorStatus": status,
	})
}

func sendCreated(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("Operation %s successful", op),
		"data":    data,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Operation %s successful", op),
		"data":    data,
	})
}
