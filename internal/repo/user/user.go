package user

import (
	"xm/internal/entity/models"
	"xm/logger"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) (result *models.User, err error)
	GetUser(name, password string) (result *models.User, err error)
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
