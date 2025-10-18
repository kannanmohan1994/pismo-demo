package handler

import (
	"pismo/internal/validation"
	"pismo/logger"
)

type Handler struct {
	AccountHandler     *accountHandler
	TransactionHandler *transactionHandler
	UserHandler        *userHandler
	HealthHandler      *healthHandler
}

func Init(uc *validation.Validation, logger logger.Log) *Handler {
	return &Handler{
		AccountHandler:     InitAccountHandler(uc.AccountValidation, logger),
		TransactionHandler: InitTransactionHandler(uc.TransactionValidation, logger),
		UserHandler:        InitUserHandler(uc.UserValidation, logger),
		HealthHandler:      InitHealthHandler(),
	}
}
