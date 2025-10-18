package models

type OperationType struct {
	ID          int    `json:"operation_type_id" gorm:"primaryKey"`
	Description string `json:"description"`
}
