package request

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

type GetAccountRequest struct {
	AccountID int `json:"account_id"`
}

func (req *GetAccountRequest) Decode(c *gin.Context) (err error) {
	req.AccountID, err = strconv.Atoi(c.Param("accountid"))
	return err
}

func (req *CreateAccountRequest) Decode(c *gin.Context) (err error) {
	return c.BindJSON(&req)
}
