package response

import (
	"xm/internal/entity/models"
)

type GetCompanyResponse struct {
	models.Company
}

type CreateCompanyResponse struct {
	models.Company
}
