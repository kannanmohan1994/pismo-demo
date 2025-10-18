package transaction

import (
	"context"
	"pismo/internal/entity/models"
	"pismo/logger"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, txn *models.Transactions) (result *models.Transactions, err error)
}

type repo struct {
	logger logger.Log
	db     *gorm.DB
}

func InitTransactionRepo(db *gorm.DB, logger logger.Log) TransactionRepository {
	return &repo{
		logger: logger,
		db:     db,
	}
}
