package request

import "xm/consts"

type CreateCompanyRequest struct {
	Name          string             `json:"name" validate:"required"`
	Description   string             `json:"description"`
	EmployeeCount *int               `json:"employeeCount" validate:"required"`
	IsRegistered  *bool              `json:"isRegistered" validate:"required"`
	Type          consts.CompanyType `json:"companyType" validate:"required"`
}

type PatchCompanyRequest struct {
	Name          string             `json:"name,omitempty"`
	Description   string             `json:"description,omitempty"`
	EmployeeCount *int               `json:"employeeCount,omitempty"`
	IsRegistered  *bool              `json:"isRegistered,omitempty"`
	Type          consts.CompanyType `json:"companyType,omitempty"`
}
