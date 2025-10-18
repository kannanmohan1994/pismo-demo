package transaction

import (
	"context"
	"pismo/internal/entity/models"
	"pismo/internal/entity/request"
	"time"
)

func (u *usecase) CreateTransaction(ctx context.Context, req *request.CreateTransactionRequest) (result *models.Transactions, err error) {
	if _, err = u.account.GetAccount(ctx, req.AccountID); err != nil {
		return
	}

	if _, err = u.operationtype.GetOperationType(ctx, req.OperationTypeID); err != nil {
		return
	}
	txns := &models.Transactions{
		AccountID:       req.AccountID,
		OperationTypeID: req.OperationTypeID,
		Amount:          req.Amount,
		EventDate:       time.Now().UTC(),
	}
	return u.transaction.CreateTransaction(ctx, txns)
}
