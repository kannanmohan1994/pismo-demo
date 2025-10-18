package account

import (
	"context"
	"pismo/internal/entity/models"
	"pismo/internal/entity/request"
	"pismo/utils"
)

func (u *usecase) CreateAccount(ctx context.Context, req *request.CreateAccountRequest) (result *models.Accounts, err error) {
	account := &models.Accounts{
		DocumentNumber: req.DocumentNumber,
	}
	isExist, err := u.account.CheckAccountExists(ctx, req.DocumentNumber)
	if err != nil {
		return
	} else if isExist {
		return nil, utils.ErrDocumentNumberAlreadyExist
	}
	return u.account.CreateAccount(ctx, account)
}

func (u *usecase) GetAccount(ctx context.Context, req *request.GetAccountRequest) (result *models.Accounts, err error) {
	return u.account.GetAccount(ctx, req.AccountID)
}
