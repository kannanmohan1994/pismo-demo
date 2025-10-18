package validation

import (
	"pismo/internal/usecase"
	"pismo/internal/usecase/account"
	"pismo/internal/usecase/transaction"
	"pismo/internal/usecase/user"
)

type Validation struct {
	AccountValidation     account.UsecaseInterface
	TransactionValidation transaction.UsecaseInterface
	UserValidation        user.UsecaseInterface
}

func Init(uc *usecase.Usecase) *Validation {
	return &Validation{
		AccountValidation:     InitAccountValidation(uc.Account),
		TransactionValidation: InitTransactionValidation(uc.Transaction),
		UserValidation:        InitUserValidation(uc.User),
	}
}
