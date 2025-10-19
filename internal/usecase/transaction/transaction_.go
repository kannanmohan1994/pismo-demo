package transaction

import (
	"context"
	"fmt"
	"pismo/internal/entity/models"
	"pismo/internal/entity/request"
	"pismo/utils"
	"time"
)

func (u *usecase) CreateTransaction(ctx context.Context, req *request.CreateTransactionRequest) (result *models.Transactions, err error) {
	if _, err = u.account.GetAccount(ctx, req.AccountID); err != nil {
		return
	}

	operationType, err := u.operationtype.GetOperationType(ctx, req.OperationTypeID)
	if err != nil {
		return
	}

	isAmountCredit := (req.Amount > 0)
	isOperationCredit := operationType.IsCredit

	// handles operation credit mismatch with amount
	if isAmountCredit != isOperationCredit {
		if isOperationCredit {
			err = fmt.Errorf(utils.ErrOperationTypeCredit, operationType.Description)
		} else {
			err = fmt.Errorf(utils.ErrOperationTypeNonCredit, operationType.Description)
		}
		u.logger.Errorf(err.Error(), "db_operation_type", operationType.Description, "amount", req.Amount)
		return result, err
	}

	txns := &models.Transactions{
		AccountID:       req.AccountID,
		OperationTypeID: req.OperationTypeID,
		Amount:          req.Amount,
		EventDate:       time.Now().UTC(),
	}
	return u.transaction.CreateTransaction(ctx, txns)
}
