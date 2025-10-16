package handler

import (
	"xm/internal/usecase"
	"xm/logger"
)

type Handler struct {
	CompanyHandler *companyHandler
	UserHandler    *userHandler
	HealthHandler  *healthHandler
}

func Init(uc *usecase.Usecase, logger logger.Log) *Handler {
	return &Handler{
		CompanyHandler: InitCompanyHandler(uc.Company, logger),
		UserHandler:    InitUserHandler(uc.User, logger),
		HealthHandler:  InitHealthHandler(),
	}
}
