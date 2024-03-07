package handler

import (
	"fmt"
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
// @Success 200 {object} ShowOpeningResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func GetOpeningHandler(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("Error fetching opening with id: %s. Error: %s", id, err)
		sendError(c, http.StatusNotFound, fmt.Sprintf("Error fetching opening with id: %s", id))
		return
	}

	sendSuccess(c, "show-opening", opening)
}
