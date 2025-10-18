package handler

import (
	"net/http"
	"pismo/internal/entity/response"
	"pismo/utils"

	"github.com/gin-gonic/gin"
)

type healthHandler struct{}

func InitHealthHandler() *healthHandler {
	return &healthHandler{}
}

// HandleGetHealth godoc
// @Summary Health check
// @Description Returns the current health status of the service.
// @Tags Health
// @Produce json
// @Success 200 {object} utils.Response
// @Router /health [get]
func (h *healthHandler) HandleGetHealth(c *gin.Context) {
	result := response.Health{
		Status:  "normal",
		Message: "system running normally",
	}
	c.JSON(http.StatusOK, utils.Send(result, nil, ""))
}
