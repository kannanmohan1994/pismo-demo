package models

type Accounts struct {
	ID             int    `json:"account_id" gorm:"primaryKey"`
	DocumentNumber string `json:"document_number"`
}
