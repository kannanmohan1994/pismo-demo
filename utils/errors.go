package utils

import (
	"errors"
)

var (
	ErrInvalidToken               = errors.New("invalid token")
	ErrEmptyToken                 = errors.New("empty token")
	ErrInvalidScheme              = errors.New("invalid scheme")
	ErrAccountCreation            = errors.New("account creation failed")
	ErrAccountFetch               = errors.New("account details fetch failed")
	ErrAccountNotFound            = errors.New("account not found")
	ErrDocumentNumberAlreadyExist = errors.New("document number already exists")
	ErrOperationTypeFetch         = errors.New("operation details fetch failed")
	ErrOperationTypeNotFound      = errors.New("operation not found")
	ErrTransactionCreation        = errors.New("transaction creation failed")
	ErrInvalidDocumentNumber      = errors.New("invalid document number: must be exactly 11 digits")
	ErrInvalidAmount              = errors.New("invalid amount")
	ErrInvalidAccountID           = errors.New("invalid account")
	ErrInvalidOperationTypeID     = errors.New("invalid operation type")

	// error template
	ErrOperationTypeNonCredit = "Invalid transaction: operation type '%s' is non-credit, but received a positive amount. Use a negative amount for this operation."
	ErrOperationTypeCredit    = "Invalid transaction: operation type '%s' is credit, but received a negative amount. Use a positive amount for this operation."
)

var (
	ErrCodeUnauthorized       = "UNAUTHORIZED"
	ErrCodeResourceCreation   = "RESOURCE_NOT_CREATED"
	ErrCodeResourceNotFound   = "RESOURCE_NOT_FOUND"
	ErrCodeResourceNotFetched = "RESOURCE_NOT_FETCHED"
	ErrCodeValidation         = "VALIDATION_ERROR"
)

var errCodeMap map[error]string = map[error]string{
	ErrInvalidToken:  ErrCodeUnauthorized,
	ErrEmptyToken:    ErrCodeUnauthorized,
	ErrInvalidScheme: ErrCodeUnauthorized,

	ErrAccountCreation:            ErrCodeResourceCreation,
	ErrTransactionCreation:        ErrCodeResourceCreation,
	ErrDocumentNumberAlreadyExist: ErrCodeResourceCreation,

	ErrAccountFetch:       ErrCodeResourceNotFetched,
	ErrOperationTypeFetch: ErrCodeResourceNotFetched,

	ErrAccountNotFound:       ErrCodeResourceNotFound,
	ErrOperationTypeNotFound: ErrCodeResourceNotFound,

	ErrInvalidDocumentNumber:  ErrCodeValidation,
	ErrInvalidAmount:          ErrCodeValidation,
	ErrInvalidAccountID:       ErrCodeValidation,
	ErrInvalidOperationTypeID: ErrCodeValidation,
}
