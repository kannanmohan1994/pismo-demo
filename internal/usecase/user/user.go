package user

import (
	context "context"
	"pismo/internal/entity/response"
	"pismo/internal/middleware"
	user "pismo/internal/repo/user"
	"pismo/logger"
)

type UsecaseInterface interface {
	CreateUser(ctx context.Context, name, password string) (resp *response.UserResponse, err error)
	LoginUser(ctx context.Context, name, password string) (resp *response.UserResponse, err error)
}

type usecase struct {
	logger    logger.Log
	tokenFunc middleware.TokenFunc
	user      user.UserRepository
}

func InitUserUsecase(user user.UserRepository, tokenFunc middleware.TokenFunc, logger logger.Log) UsecaseInterface {
	return &usecase{
		logger:    logger,
		tokenFunc: tokenFunc,
		user:      user,
	}
}
