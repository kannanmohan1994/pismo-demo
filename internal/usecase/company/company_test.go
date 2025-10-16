package company

import (
	"errors"
	"testing"
	"xm/consts"
	"xm/internal/entity/models"
	"xm/internal/entity/request"
	notifier "xm/internal/notifiers"
	"xm/logger"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var employeeCount int = 500
var isRegistered bool = false

func TestService_CreateCompany(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name          string
		request       *request.CreateCompanyRequest
		err           map[string]error
		repoCallCount int
	}{
		{
			name: "should_return_error_for_repo_call",
			request: &request.CreateCompanyRequest{
				Name:          "xm",
				Description:   "Forex & CFD Trading on Stocks, Indices, Oil, Gold by XM",
				EmployeeCount: &employeeCount,
				IsRegistered:  &isRegistered,
				Type:          consts.CompanyTypeCorporations,
			},
			err:           map[string]error{"CreateCompany": errors.New("repo call fail")},
			repoCallCount: 1,
		},
		{
			name: "success",
			request: &request.CreateCompanyRequest{
				Name:          "xm",
				Description:   "Forex & CFD Trading on Stocks, Indices, Oil, Gold by XM",
				EmployeeCount: &employeeCount,
				IsRegistered:  &isRegistered,
				Type:          consts.CompanyTypeCorporations,
			},
			repoCallCount: 1,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewMockCompanyRepository(ctrl)
			logger := logger.NewNoop()
			notifier := notifier.NewNoop()
			repo.EXPECT().CreateCompany(gomock.Any()).Return(&models.Company{}, tt.err["CreateCompany"]).Times(tt.repoCallCount)
			s := InitCompanyUsecase(repo, logger, notifier)

			_, err := s.CreateCompany(tt.request)
			if tt.err != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestService_GetCompany(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name          string
		companyId     string
		err           map[string]error
		repoCallCount int
	}{
		{
			name:          "should_return_error_for_repo_call",
			companyId:     "company-id-1",
			err:           map[string]error{"GetCompany": errors.New("repo call fail")},
			repoCallCount: 1,
		},
		{
			name:          "success",
			companyId:     "company-id-1",
			repoCallCount: 1,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewMockCompanyRepository(ctrl)
			logger := logger.NewNoop()
			notifier := notifier.NewNoop()
			repo.EXPECT().GetCompany(gomock.Any()).Return(&models.Company{}, tt.err["GetCompany"]).Times(tt.repoCallCount)
			s := InitCompanyUsecase(repo, logger, notifier)

			_, err := s.GetCompany(tt.companyId)
			if tt.err != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestService_PatchCompany(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name               string
		companyId          string
		request            *request.PatchCompanyRequest
		err                map[string]error
		getRepoCallCount   int
		patchRepoCallCount int
	}{
		{
			name:      "should_return_error_for_get_company_repo_call",
			companyId: "company-id-1",
			request: &request.PatchCompanyRequest{
				Name:          "xm",
				Description:   "forex trading",
				EmployeeCount: &employeeCount,
				IsRegistered:  &isRegistered,
				Type:          consts.CompanyTypeCooperative,
			},
			err:                map[string]error{"GetCompany": errors.New("repo call fail")},
			getRepoCallCount:   1,
			patchRepoCallCount: 0,
		},
		{
			name:      "should_return_error_for_patch_company_repo_call",
			companyId: "company-id-1",
			request: &request.PatchCompanyRequest{
				Name:          "xm",
				Description:   "forex trading",
				EmployeeCount: &employeeCount,
				IsRegistered:  &isRegistered,
				Type:          consts.CompanyTypeCooperative,
			},
			err:                map[string]error{"PatchCompany": errors.New("repo call fail")},
			getRepoCallCount:   1,
			patchRepoCallCount: 1,
		},
		{
			name:      "success",
			companyId: "company-id-1",
			request: &request.PatchCompanyRequest{
				Name:          "xm",
				Description:   "forex trading",
				EmployeeCount: &employeeCount,
				IsRegistered:  &isRegistered,
				Type:          consts.CompanyTypeCooperative,
			},
			getRepoCallCount:   1,
			patchRepoCallCount: 1,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewMockCompanyRepository(ctrl)
			logger := logger.NewNoop()
			notifier := notifier.NewNoop()
			repo.EXPECT().GetCompany(gomock.Any()).Return(&models.Company{}, tt.err["GetCompany"]).Times(tt.getRepoCallCount)
			repo.EXPECT().PatchCompany(gomock.Any(), gomock.Any()).Return(&models.Company{}, tt.err["PatchCompany"]).Times(tt.patchRepoCallCount)
			s := InitCompanyUsecase(repo, logger, notifier)

			_, err := s.PatchCompany(tt.companyId, tt.request)
			if tt.err != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestService_DeleteCompany(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name                string
		companyId           string
		err                 map[string]error
		getRepoCallCount    int
		deleteRepoCallCount int
	}{
		{
			name:                "should_return_error_for_get_company_repo_call",
			companyId:           "company-id-1",
			err:                 map[string]error{"GetCompany": errors.New("repo call fail")},
			getRepoCallCount:    1,
			deleteRepoCallCount: 0,
		},
		{
			name:                "should_return_error_for_delete_company_repo_call",
			companyId:           "company-id-1",
			err:                 map[string]error{"DeleteCompany": errors.New("repo call fail")},
			getRepoCallCount:    1,
			deleteRepoCallCount: 1,
		},
		{
			name:                "success",
			companyId:           "company-id-1",
			getRepoCallCount:    1,
			deleteRepoCallCount: 1,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewMockCompanyRepository(ctrl)
			logger := logger.NewNoop()
			notifier := notifier.NewNoop()
			repo.EXPECT().GetCompany(gomock.Any()).Return(&models.Company{}, tt.err["GetCompany"]).Times(tt.getRepoCallCount)
			repo.EXPECT().DeleteCompany(gomock.Any()).Return(tt.err["DeleteCompany"]).Times(tt.deleteRepoCallCount)
			s := InitCompanyUsecase(repo, logger, notifier)

			err := s.DeleteCompany(tt.companyId)
			if tt.err != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
