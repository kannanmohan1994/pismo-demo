package handler

import (
	"context"
	"net/http"
	"pismo/internal/entity/request"
	"pismo/internal/entity/response"
	"pismo/internal/usecase/transaction"
	"pismo/logger"

	"github.com/gin-gonic/gin"

	_ "pismo/utils"
)

type transactionHandler struct {
	transactionUC transaction.UsecaseInterface
	logger        logger.Log
}

func InitTransactionHandler(uc transaction.UsecaseInterface, logger logger.Log) *transactionHandler {
	return &transactionHandler{
		transactionUC: uc,
		logger:        logger,
	}
}

// HandleCreateTransaction godoc
// @Summary Create transaction
// @Description Creates a new transaction for the specified account and operation.
// @Tags Transactions
// @Accept json
// @Produce json
// @Param request body request.CreateTransactionRequest true "Create transaction payload"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /transactions [post]
func (h *transactionHandler) HandleCreateTransaction(c *gin.Context) {
	handleRequest(
		c,
		&request.CreateTransactionRequest{},
		func(ctx context.Context, r RequestHandler) (code int, data interface{}, err error) {
			req := r.(*request.CreateTransactionRequest)
			resp := &response.CreateTransactionResponse{}
			result, err := h.transactionUC.CreateTransaction(ctx, req)
			if err != nil {
				return
			}
			resp.Encode(result)
			return http.StatusOK, resp, err
		},
	)
}
