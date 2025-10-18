package account

import (
	"context"
	"errors"
	"pismo/internal/entity/models"
	"pismo/internal/entity/request"
	"pismo/internal/mocks"
	"pismo/logger"
	"pismo/utils"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	account "pismo/internal/mocks"
)

func TestService_CreateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name                          string
		request                       *request.CreateAccountRequest
		repoErr                       map[string]error
		wantErr                       error
		checkAccountExistsRepoCall    int
		createAccountRepoCall         int
		checkAccountExistsReturnValue bool
	}{
		{
			name: "should_return_error_for_get_account_repo_call",
			request: &request.CreateAccountRequest{
				DocumentNumber: "1234567890",
			},
			repoErr:                       map[string]error{"CheckAccountExists": errors.New("repo call fail")},
			wantErr:                       errors.New("repo call fail"),
			checkAccountExistsRepoCall:    1,
			checkAccountExistsReturnValue: false,
		},
		{
			name: "should_return_error_for_account_exists_call",
			request: &request.CreateAccountRequest{
				DocumentNumber: "1234567890",
			},
			wantErr:                       utils.ErrDocumentNumberAlreadyExist,
			checkAccountExistsRepoCall:    1,
			checkAccountExistsReturnValue: true,
		},
		{
			name: "should_return_error_for_create_account_repo_call",
			request: &request.CreateAccountRequest{
				DocumentNumber: "1234567890",
			},
			repoErr:                       map[string]error{"CreateAccount": errors.New("repo call fail")},
			wantErr:                       errors.New("repo call fail"),
			checkAccountExistsRepoCall:    1,
			createAccountRepoCall:         1,
			checkAccountExistsReturnValue: false,
		},
		{
			name: "success",
			request: &request.CreateAccountRequest{
				DocumentNumber: "1234567890",
			},
			checkAccountExistsRepoCall:    1,
			createAccountRepoCall:         1,
			checkAccountExistsReturnValue: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			repo := account.NewMockAccountRepository(ctrl)
			logger := logger.NewNoop()
			repo.EXPECT().CheckAccountExists(context.Background(), gomock.Any()).Return(tt.checkAccountExistsReturnValue, tt.repoErr["CheckAccountExists"]).Times(tt.checkAccountExistsRepoCall)
			repo.EXPECT().CreateAccount(context.Background(), gomock.Any()).Return(&models.Accounts{}, tt.repoErr["CreateAccount"]).Times(tt.createAccountRepoCall)
			s := InitAccountUsecase(repo, logger)

			_, err := s.CreateAccount(context.Background(), tt.request)
			if tt.wantErr != nil {
				require.EqualError(t, err, tt.wantErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestService_GetAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name          string
		request       *request.GetAccountRequest
		wantErr       error
		repoErr       map[string]error
		repoCallCount int
	}{
		{
			name:          "should_return_error_for_repo_call",
			request:       &request.GetAccountRequest{AccountID: 1},
			wantErr:       errors.New("repo call fail"),
			repoErr:       map[string]error{"GetAccount": errors.New("repo call fail")},
			repoCallCount: 1,
		},
		{
			name:          "success",
			request:       &request.GetAccountRequest{AccountID: 1},
			repoCallCount: 1,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewMockAccountRepository(ctrl)
			logger := logger.NewNoop()
			repo.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Return(&models.Accounts{}, tt.repoErr["GetAccount"]).Times(tt.repoCallCount)
			s := InitAccountUsecase(repo, logger)

			_, err := s.GetAccount(context.Background(), tt.request)
			if tt.wantErr != nil {
				require.EqualError(t, err, tt.wantErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
