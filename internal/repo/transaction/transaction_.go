package transaction

import (
	"context"
	"pismo/internal/entity/models"
	"pismo/utils"
)

func (r *repo) CreateTransaction(ctx context.Context, txn *models.Transactions) (result *models.Transactions, err error) {
	err = r.db.Create(txn).Error
	if err != nil {
		r.logger.WithContext(ctx).Errorf("error creating account",
			"account_id", txn.AccountID,
			"operation_type", txn.OperationTypeID,
			"error", err.Error())
		return txn, utils.ErrTransactionCreation
	}
	return txn, nil
}
