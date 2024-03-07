package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vpcmps/gopportunities/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)

	logger.Infof("CreateOpeningHandler: %+v", request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("CreateOpeningHandler: %s", err)
		sendError(ctx, http.StatusInternalServerError, "Error creating opening")
		return
	}

	sendCreated(ctx, "create-opening", opening)
}
