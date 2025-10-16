package repo

import (
	company "xm/internal/repo/company"
	"xm/internal/repo/user"
	"xm/logger"

	"gorm.io/gorm"
)

type Repo struct {
	Logger  logger.Log
	Company company.CompanyRepository
	User    user.UserRepository
}

func InitRepo(db *gorm.DB, logger logger.Log) *Repo {
	return &Repo{
		Logger:  logger,
		Company: company.InitCompanyRepo(db, logger),
		User:    user.InitUserRepo(db, logger),
	}
}
