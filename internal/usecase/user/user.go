package user

import (
	"xm/internal/entity/response"
	"xm/internal/middleware"
	notifiers "xm/internal/notifiers"
	user "xm/internal/repo/user"
	"xm/logger"
)

type UsecaseInterface interface {
	CreateUser(name, password string) (resp *response.UserResponse, err error)
	LoginUser(name, password string) (resp *response.UserResponse, err error)
}

type usecase struct {
	logger     logger.Log
	notifier   notifiers.Notifier
	middleware middleware.Middleware
	user       user.UserRepository
}

func InitUserUsecase(user user.UserRepository, middleware middleware.Middleware, logger logger.Log,
	notifier notifiers.Notifier) UsecaseInterface {
	return &usecase{
		logger:     logger,
		notifier:   notifier,
		middleware: middleware,
		user:       user,
	}
}
