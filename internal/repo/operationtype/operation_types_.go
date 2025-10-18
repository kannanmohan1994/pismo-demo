package operationtype

import (
	"context"
	"errors"
	"pismo/internal/entity/models"
	"pismo/utils"

	"gorm.io/gorm"
)

func (r *repo) GetOperationType(ctx context.Context, id int) (operationType *models.OperationType, err error) {
	operationType = &models.OperationType{ID: id}
	result := r.db.First(operationType)
	if err = result.Error; err != nil {
		r.logger.WithContext(ctx).Errorf("operation type fetch failed", "id", id, "error", err.Error())
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return operationType, utils.ErrOperationTypeNotFound
		} else {
			return operationType, utils.ErrOperationTypeFetch
		}
	}
	return operationType, nil
}
