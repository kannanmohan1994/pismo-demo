package handler

import (
	user "pismo/internal/usecase/user"
	"pismo/logger"
)

type userHandler struct {
	logger logger.Log
	userUC user.UsecaseInterface
}

func InitUserHandler(uc user.UsecaseInterface, logger logger.Log) *userHandler {
	return &userHandler{
		logger: logger,
		userUC: uc,
	}
}
