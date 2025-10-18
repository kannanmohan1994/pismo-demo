package handler

import (
	"context"
	"net/http"
	"pismo/internal/entity/request"
	"pismo/internal/entity/response"
	"pismo/internal/usecase/account"
	"pismo/logger"

	"github.com/gin-gonic/gin"

	_ "pismo/utils"
)

type accountHandler struct {
	accountUC account.UsecaseInterface
	logger    logger.Log
}

func InitAccountHandler(uc account.UsecaseInterface, logger logger.Log) *accountHandler {
	return &accountHandler{
		accountUC: uc,
		logger:    logger,
	}
}

// HandleCreateAccount godoc
// @Summary Create account
// @Description Creates a new account with the provided document number.
// @Tags Accounts
// @Accept json
// @Produce json
// @Param request body request.CreateAccountRequest true "Create account payload"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /accounts [post]
func (h *accountHandler) HandleCreateAccount(c *gin.Context) {
	handleRequest(
		c,
		&request.CreateAccountRequest{},
		func(ctx context.Context, r RequestHandler) (code int, data interface{}, err error) {
			req := r.(*request.CreateAccountRequest)
			resp := &response.CreateAccountResponse{}
			result, err := h.accountUC.CreateAccount(ctx, req)
			if err != nil {
				return
			}
			resp.Encode(result)
			return http.StatusOK, resp, err
		},
	)
}

// HandleGetAccount godoc
// @Summary Get account
// @Description Retrieves account information by account identifier.
// @Tags Accounts
// @Produce json
// @Param accountid path string true "Account ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /accounts/{accountid} [get]
func (h *accountHandler) HandleGetAccount(c *gin.Context) {
	handleRequest(
		c,
		&request.GetAccountRequest{},
		func(ctx context.Context, r RequestHandler) (code int, data interface{}, err error) {
			req := r.(*request.GetAccountRequest)
			resp := &response.GetAccountResponse{}
			result, err := h.accountUC.GetAccount(ctx, req)
			if err != nil {
				return
			}
			resp.Encode(result)
			return http.StatusOK, resp, err
		},
	)
}
