package operationtype

import (
	"context"
	"pismo/internal/entity/models"
	"pismo/logger"

	"gorm.io/gorm"
)

type OperationTypeRepository interface {
	GetOperationType(ctx context.Context, id int) (operationType *models.OperationType, err error)
}

type repo struct {
	logger logger.Log
	db     *gorm.DB
}

func InitOperationTypeRepo(db *gorm.DB, logger logger.Log) OperationTypeRepository {
	return &repo{
		logger: logger,
		db:     db,
	}
}
