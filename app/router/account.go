package router

import (
	"github.com/gin-gonic/gin"
)

func (c *Capsule) AccountRoutes(r *gin.RouterGroup) {
	accountHandler := c.Handler.AccountHandler

	accountV1 := r.Group("/accounts")

	accountV1.GET("/:accountid", accountHandler.HandleGetAccount)
	accountV1.POST("", accountHandler.HandleCreateAccount)
}
