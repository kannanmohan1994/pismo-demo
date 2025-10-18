package usecase

import (
	"pismo/internal/middleware"
	"pismo/internal/repo"
	"pismo/internal/usecase/account"
	"pismo/internal/usecase/transaction"
	"pismo/internal/usecase/user"
	"pismo/logger"
)

type Usecase struct {
	Account     account.UsecaseInterface
	Transaction transaction.UsecaseInterface
	User        user.UsecaseInterface
}

func Init(repo *repo.Repo, tokenFunc middleware.TokenFunc, logger logger.Log) *Usecase {
	return &Usecase{
		Account:     account.InitAccountUsecase(repo.Account, logger),
		Transaction: transaction.InitTransactionUsecase(repo.Transaction, repo.Account, repo.OperationType, logger),
		User:        user.InitUserUsecase(repo.User, tokenFunc, logger),
	}
}
