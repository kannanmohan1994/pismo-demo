package request

import "github.com/gin-gonic/gin"

type CreateTransactionRequest struct {
	AccountID       int     `json:"account_id"`
	OperationTypeID int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

func (req *CreateTransactionRequest) Decode(c *gin.Context) (err error) {
	return c.BindJSON(&req)
}
