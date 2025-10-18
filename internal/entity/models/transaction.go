package models

import "time"

type Transactions struct {
	ID              int       `json:"transaction_id" gorm:"primaryKey"`
	AccountID       int       `json:"account_id"`
	OperationTypeID int       `json:"operation_type_id"`
	Amount          float64   `json:"amount"`
	EventDate       time.Time `json:"event_date"`
}
