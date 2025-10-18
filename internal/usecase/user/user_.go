package user

import (
	"context"
	"pismo/internal/entity/models"
	"pismo/internal/entity/response"
)

func (u *usecase) CreateUser(ctx context.Context, name, password string) (resp *response.UserResponse, err error) {
	user := &models.User{
		Name:     name,
		Password: password,
	}

	accessToken, err := u.tokenFunc(user.ID)
	if err != nil {
		return nil, err
	}

	_, err = u.user.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &response.UserResponse{
		AccessToken: accessToken,
	}, nil
}
func (u *usecase) LoginUser(ctx context.Context, name, password string) (resp *response.UserResponse, err error) {
	user, err := u.user.GetUser(ctx, name, password)
	if err != nil {
		return nil, err
	}

	accessToken, err := u.tokenFunc(user.ID)
	if err != nil {
		return nil, err
	}

	return &response.UserResponse{
		UserID:      user.ID,
		AccessToken: accessToken,
	}, nil
}
