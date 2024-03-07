package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vpcmps/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a existing job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Request body"
// @Param request body UpdateOpeningRequest true "Request body"
// @Success 200 {object} ShowOpeningResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("Error fetching opening with id: %s. Error: %s", id, err)
		sendError(ctx, http.StatusNotFound, "Error fetching opening")
		return
	}
	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary != 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("UpdateOpeningHandler: %s", err)
		sendError(ctx, http.StatusInternalServerError, "Error updating opening")
		return
	}

	sendSuccess(ctx, "update-opening", opening)
}
