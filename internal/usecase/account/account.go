package account

import (
	"context"
	"pismo/internal/entity/models"
	"pismo/internal/entity/request"
	"pismo/internal/repo/account"
	"pismo/logger"
)

type UsecaseInterface interface {
	CreateAccount(ctx context.Context, req *request.CreateAccountRequest) (result *models.Accounts, err error)
	GetAccount(ctx context.Context, req *request.GetAccountRequest) (result *models.Accounts, err error)
}

type usecase struct {
	logger  logger.Log
	account account.AccountRepository
}

func InitAccountUsecase(account account.AccountRepository, logger logger.Log) UsecaseInterface {
	return &usecase{
		logger:  logger,
		account: account,
	}
}
