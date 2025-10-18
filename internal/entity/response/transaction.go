package response

import "pismo/internal/entity/models"

type CreateTransactionResponse struct {
	*models.Transactions
}

func (r *CreateTransactionResponse) Encode(result *models.Transactions) {
	r.Transactions = result
}
