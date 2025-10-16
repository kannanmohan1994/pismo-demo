package company

import (
	"xm/internal/entity/models"
	"xm/logger"

	"gorm.io/gorm"
)

type CompanyRepository interface {
	GetCompany(id string) (result *models.Company, err error)
	CreateCompany(obj *models.Company) (result *models.Company, err error)
	PatchCompany(id string, obj *models.Company) (result *models.Company, err error)
	DeleteCompany(id string) (err error)
}

type repo struct {
	logger logger.Log
	db     *gorm.DB
}

func InitCompanyRepo(db *gorm.DB, logger logger.Log) CompanyRepository {
	return &repo{
		logger: logger,
		db:     db,
	}
}
