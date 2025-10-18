package router

import (
	"github.com/gin-gonic/gin"
)

func (c *Capsule) TransactionRoutes(r *gin.RouterGroup) {
	transactionHandler := c.Handler.TransactionHandler

	transactionV1 := r.Group("/transactions")
	transactionV1.POST("", transactionHandler.HandleCreateTransaction)
}
