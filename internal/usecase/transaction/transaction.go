package transaction

import (
	"context"
	"pismo/internal/entity/models"
	"pismo/internal/entity/request"
	"pismo/internal/repo/account"
	"pismo/internal/repo/operationtype"
	"pismo/internal/repo/transaction"
	"pismo/logger"
)

type UsecaseInterface interface {
	CreateTransaction(ctx context.Context, req *request.CreateTransactionRequest) (result *models.Transactions, err error)
}

type usecase struct {
	logger        logger.Log
	transaction   transaction.TransactionRepository
	account       account.AccountRepository
	operationtype operationtype.OperationTypeRepository
}

func InitTransactionUsecase(transaction transaction.TransactionRepository,
	account account.AccountRepository,
	operationtype operationtype.OperationTypeRepository,
	logger logger.Log) UsecaseInterface {
	return &usecase{
		transaction:   transaction,
		account:       account,
		operationtype: operationtype,
		logger:        logger,
	}
}
