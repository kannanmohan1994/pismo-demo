package account

import (
	"context"
	"pismo/internal/entity/models"
	"pismo/logger"

	"gorm.io/gorm"
)

type AccountRepository interface {
	GetAccount(ctx context.Context, id int) (result *models.Accounts, err error)
	CheckAccountExists(ctx context.Context, documentNumber string) (isExist bool, err error)
	CreateAccount(ctx context.Context, obj *models.Accounts) (result *models.Accounts, err error)
}

type repo struct {
	logger logger.Log
	db     *gorm.DB
}

func InitAccountRepo(db *gorm.DB, logger logger.Log) AccountRepository {
	return &repo{
		logger: logger,
		db:     db,
	}
}
