package models

import (
	"time"
	"xm/consts"
)

type Company struct {
	ID            string             `json:"id" gorm:"primaryKey"`
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	EmployeeCount *int               `json:"employeeCount"`
	IsRegistered  *bool              `json:"isRegistered"`
	Type          consts.CompanyType `json:"companyType"`
	CreatedAt     time.Time          `json:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt"`
}
