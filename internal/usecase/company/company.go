package company

import (
	"xm/internal/entity/models"
	"xm/internal/entity/request"
	notifiers "xm/internal/notifiers"
	company "xm/internal/repo/company"
	"xm/logger"
)

type UsecaseInterface interface {
	GetCompany(id string) (result *models.Company, err error)
	CreateCompany(req *request.CreateCompanyRequest) (result *models.Company, err error)
	PatchCompany(id string, req *request.PatchCompanyRequest) (result *models.Company, err error)
	DeleteCompany(id string) (err error)
}

type usecase struct {
	logger   logger.Log
	notifier notifiers.Notifier
	company  company.CompanyRepository
}

func InitCompanyUsecase(company company.CompanyRepository, logger logger.Log, notifier notifiers.Notifier) UsecaseInterface {
	return &usecase{
		logger:   logger,
		notifier: notifier,
		company:  company,
	}
}
