package repo

import (
	account "pismo/internal/repo/account"
	"pismo/internal/repo/operationtype"
	"pismo/internal/repo/transaction"
	"pismo/internal/repo/user"
	"pismo/logger"

	"gorm.io/gorm"
)

type Repo struct {
	Logger        logger.Log
	Account       account.AccountRepository
	OperationType operationtype.OperationTypeRepository
	Transaction   transaction.TransactionRepository
	User          user.UserRepository
}

func InitRepo(db *gorm.DB, logger logger.Log) *Repo {
	return &Repo{
		Logger:        logger,
		Account:       account.InitAccountRepo(db, logger),
		OperationType: operationtype.InitOperationTypeRepo(db, logger),
		Transaction:   transaction.InitTransactionRepo(db, logger),
		User:          user.InitUserRepo(db, logger),
	}
}
