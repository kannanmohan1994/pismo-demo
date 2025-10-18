package validation

import (
	"context"
	"pismo/internal/entity/response"
	"pismo/internal/usecase/user"
)

type UserValidation struct {
	userUC user.UsecaseInterface
}

func InitUserValidation(uc user.UsecaseInterface) user.UsecaseInterface {
	return &UserValidation{
		userUC: uc,
	}
}

func (u *UserValidation) CreateUser(ctx context.Context, name, password string) (resp *response.UserResponse, err error) {
	return u.userUC.CreateUser(ctx, name, password)
}
func (u *UserValidation) LoginUser(ctx context.Context, name, password string) (resp *response.UserResponse, err error) {
	return u.userUC.LoginUser(ctx, name, password)
}
