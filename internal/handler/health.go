package handler

import (
	"net/http"
	"xm/internal/entity/response"
	"xm/utils"

	"github.com/gin-gonic/gin"
)

type healthHandler struct{}

func InitHealthHandler() *healthHandler {
	return &healthHandler{}
}

func (h *healthHandler) HandleGetHealth(c *gin.Context) {
	result := response.Health{
		Status:  "normal",
		Message: "system running normally",
	}
	c.JSON(http.StatusOK, utils.Send(result))
}
