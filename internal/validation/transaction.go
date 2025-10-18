package validation

import (
	"context"
	"pismo/internal/entity/models"
	"pismo/internal/entity/request"
	"pismo/internal/usecase/transaction"
	"pismo/utils"
)

type TransactionValidation struct {
	transactionUC transaction.UsecaseInterface
}

func InitTransactionValidation(uc transaction.UsecaseInterface) transaction.UsecaseInterface {
	return &TransactionValidation{
		transactionUC: uc,
	}
}

func (u *TransactionValidation) CreateTransaction(ctx context.Context, req *request.CreateTransactionRequest) (result *models.Transactions, err error) {
	if req.AccountID <= 0 {
		return nil, utils.ErrInvalidAccountID
	}
	if req.OperationTypeID <= 0 {
		return nil, utils.ErrInvalidOperationTypeID
	}
	if req.Amount == 0 {
		return nil, utils.ErrInvalidAmount
	}
	return u.transactionUC.CreateTransaction(ctx, req)
}
