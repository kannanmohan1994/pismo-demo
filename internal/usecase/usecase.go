package usecase

import (
	"xm/internal/middleware"
	notifier "xm/internal/notifiers"
	"xm/internal/repo"
	company "xm/internal/usecase/company"
	"xm/internal/usecase/user"
	"xm/logger"
)

type Usecase struct {
	Company company.UsecaseInterface
	User    user.UsecaseInterface
}

func Init(repo *repo.Repo, middleware middleware.Middleware, logger logger.Log, notifier notifier.Notifier) *Usecase {
	return &Usecase{
		Company: company.InitCompanyUsecase(repo.Company, logger, notifier),
		User:    user.InitUserUsecase(repo.User, middleware, logger, notifier),
	}
}
