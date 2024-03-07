package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vpcmps/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Create an opening
// @Description Create a new job opening
// @Tags openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 201 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
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
