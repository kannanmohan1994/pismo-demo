package account

import (
	"context"
	"errors"
	"pismo/internal/entity/models"
	"pismo/utils"

	"gorm.io/gorm"
)

func (r *repo) CreateAccount(ctx context.Context, account *models.Accounts) (result *models.Accounts, err error) {
	err = r.db.Create(&account).Error
	if err != nil {
		r.logger.WithContext(ctx).Errorf("error creating account",
			"document_number", account.DocumentNumber,
			"error", err.Error())
		return result, utils.ErrAccountCreation
	}
	return account, nil
}

func (r *repo) GetAccount(ctx context.Context, id int) (account *models.Accounts, err error) {
	account = &models.Accounts{ID: id}
	result := r.db.First(account)
	if err = result.Error; err != nil {
		r.logger.WithContext(ctx).Errorf("account fetch failed", "id", id, "error", err.Error())
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return account, utils.ErrAccountNotFound
		} else {
			return account, utils.ErrAccountFetch
		}
	}
	return account, nil
}

func (r *repo) CheckAccountExists(ctx context.Context, documentNumber string) (isExist bool, err error) {
	var count int64
	result := r.db.Model(&models.Accounts{}).Where("document_number = ?", documentNumber).Count(&count)
	if err = result.Error; err != nil {
		return isExist, utils.ErrAccountFetch
	}
	return count > 0, nil
}
