package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vpcmps/gopportunities/schemas"
)

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
