package response

import (
	"pismo/internal/entity/models"
)

type CreateAccountResponse struct {
	*models.Accounts
}

type GetAccountResponse struct {
	*models.Accounts
}

func (r *CreateAccountResponse) Encode(result *models.Accounts) {
	r.Accounts = result
}

func (r *GetAccountResponse) Encode(result *models.Accounts) {
	r.Accounts = result
}
