package user

import (
	"context"
	"pismo/internal/entity/models"
	"pismo/logger"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (result *models.User, err error)
	GetUser(ctx context.Context, name, password string) (result *models.User, err error)
}

type repo struct {
	logger logger.Log
	db     *gorm.DB
}

func InitUserRepo(db *gorm.DB, logger logger.Log) UserRepository {
	return &repo{
		logger: logger,
		db:     db,
	}
}
