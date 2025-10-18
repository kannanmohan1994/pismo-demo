package user

import (
	"context"
	"pismo/internal/entity/models"
)

func (r *repo) GetUser(ctx context.Context, name, password string) (result *models.User, err error) {
	err = r.db.Table("users").Where("name = ? AND password = ?", name, password).First(&result).Error
	if err != nil {
		return result, err
	}
	return result, err
}

func (r *repo) CreateUser(ctx context.Context, user *models.User) (result *models.User, err error) {
	err = r.db.Table("users").Create(&user).Error
	if err != nil {
		return result, err
	}
	return user, err
}
