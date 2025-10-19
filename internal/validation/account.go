package validation

import (
	"context"
	"pismo/internal/entity/models"
	"pismo/internal/entity/request"
	"pismo/internal/usecase/account"
	"pismo/utils"
	"regexp"
)

type AccountValidation struct {
	accountUC account.UsecaseInterface
}

func InitAccountValidation(uc account.UsecaseInterface) account.UsecaseInterface {
	return &AccountValidation{
		accountUC: uc,
	}
}

func (u *AccountValidation) CreateAccount(ctx context.Context, req *request.CreateAccountRequest) (result *models.Accounts, err error) {
	re := regexp.MustCompile(`^\d{11}$`)
	if !re.MatchString(req.DocumentNumber) {
		return nil, utils.ErrInvalidDocumentNumber
	}
	return u.accountUC.CreateAccount(ctx, req)
}

func (u *AccountValidation) GetAccount(ctx context.Context, req *request.GetAccountRequest) (result *models.Accounts, err error) {
	if req.AccountID <= 0 {
		return nil, utils.ErrInvalidAccountID
	}
	return u.accountUC.GetAccount(ctx, req)
}
