package transaction

import (
	"context"
	"errors"
	"pismo/internal/entity/models"
	"pismo/internal/entity/request"
	"pismo/internal/mocks"
	"pismo/logger"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService_CreateTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name                       string
		request                    *request.CreateTransactionRequest
		repoErr                    map[string]error
		wantErr                    error
		accountRepoCallCount       int
		operationTypeRepoCallCount int
		txnRepoCallCount           int
	}{
		{
			name: "should_return_error_for_account_repo_call",
			request: &request.CreateTransactionRequest{
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          100,
			},
			repoErr:              map[string]error{"GetAccount": errors.New("account repo call fail")},
			wantErr:              errors.New("account repo call fail"),
			accountRepoCallCount: 1,
		},
		{
			name: "should_return_error_for_operation_type_repo_call",
			request: &request.CreateTransactionRequest{
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          100,
			},
			repoErr:                    map[string]error{"GetOperationType": errors.New("op type repo call fail")},
			wantErr:                    errors.New("op type repo call fail"),
			accountRepoCallCount:       1,
			operationTypeRepoCallCount: 1,
		},
		{
			name: "should_return_error_for_txn_repo_call",
			request: &request.CreateTransactionRequest{
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          100,
			},
			repoErr:                    map[string]error{"CreateTransaction": errors.New("txn repo call fail")},
			wantErr:                    errors.New("txn repo call fail"),
			accountRepoCallCount:       1,
			operationTypeRepoCallCount: 1,
			txnRepoCallCount:           1,
		},
		{
			name: "success",
			request: &request.CreateTransactionRequest{
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          100,
			},
			accountRepoCallCount:       1,
			operationTypeRepoCallCount: 1,
			txnRepoCallCount:           1,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			txnRepo := mocks.NewMockTransactionRepository(ctrl)
			accountRepo := mocks.NewMockAccountRepository(ctrl)
			operationTypeRepo := mocks.NewMockOperationTypeRepository(ctrl)
			logger := logger.NewNoop()
			txnRepo.EXPECT().CreateTransaction(context.Background(), gomock.Any()).Return(&models.Transactions{}, tt.repoErr["CreateTransaction"]).Times(tt.txnRepoCallCount)
			accountRepo.EXPECT().GetAccount(context.Background(), gomock.Any()).Return(&models.Accounts{}, tt.repoErr["GetAccount"]).Times(tt.accountRepoCallCount)
			operationTypeRepo.EXPECT().GetOperationType(context.Background(), gomock.Any()).Return(&models.OperationType{}, tt.repoErr["GetOperationType"]).Times(tt.operationTypeRepoCallCount)
			s := InitTransactionUsecase(txnRepo, accountRepo, operationTypeRepo, logger)

			_, err := s.CreateTransaction(context.Background(), tt.request)
			if tt.wantErr != nil {
				require.EqualError(t, err, tt.wantErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
