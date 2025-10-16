package company

import (
	"fmt"
	"time"
	"xm/consts"
	"xm/internal/entity/models"
	"xm/internal/entity/request"

	"github.com/google/uuid"
)

func (u *usecase) CreateCompany(req *request.CreateCompanyRequest) (result *models.Company, err error) {
	u.logger.Infof("Begin Usecase - CreateCompany")

	companyId := uuid.New().String()
	company := &models.Company{
		ID:            companyId,
		Name:          req.Name,
		Description:   req.Description,
		EmployeeCount: req.EmployeeCount,
		IsRegistered:  req.IsRegistered,
		Type:          req.Type,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	}
	result, err = u.company.CreateCompany(company)
	if err != nil {
		return nil, err
	}

	if notifyErr := u.notifier.Notify(consts.NotifyEventCompleted,
		[]byte(fmt.Sprintf("completed CreateCompany with id: %s", companyId))); notifyErr != nil {
		u.logger.Warnw("error notifying", "action", "CreateCompany", "company-id", companyId)
	}
	return result, err
}

func (u *usecase) GetCompany(id string) (result *models.Company, err error) {
	u.logger.Infof("Begin Usecase - GetCompany")

	result, err = u.company.GetCompany(id)
	if err != nil {
		return result, err
	}

	if notifyErr := u.notifier.Notify(consts.NotifyEventCompleted,
		[]byte(fmt.Sprintf("completed GetCompany with id: %s", id))); notifyErr != nil {
		u.logger.Warnw("error notifying", "action", "GetCompany", "company-id", id)
	}
	return result, nil
}

func (u *usecase) PatchCompany(id string, req *request.PatchCompanyRequest) (result *models.Company, err error) {
	u.logger.Infof("Begin Usecase - PatchCompany")

	_, err = u.company.GetCompany(id)
	if err != nil {
		return result, err
	}

	company := &models.Company{
		Name:          req.Name,
		Description:   req.Description,
		EmployeeCount: req.EmployeeCount,
		IsRegistered:  req.IsRegistered,
		Type:          req.Type,
	}
	result, err = u.company.PatchCompany(id, company)
	if err != nil {
		return result, err
	}

	if notifyErr := u.notifier.Notify(consts.NotifyEventCompleted,
		[]byte(fmt.Sprintf("completed PatchCompany with id: %s", id))); notifyErr != nil {
		u.logger.Warnw("error notifying", "action", "PatchCompany", "company-id", id)
	}
	return result, nil
}

func (u *usecase) DeleteCompany(id string) (err error) {
	u.logger.Infof("Begin Usecase - PatchCompany")

	_, err = u.company.GetCompany(id)
	if err != nil {
		return err
	}

	err = u.company.DeleteCompany(id)
	if err != nil {
		return err
	}

	if notifyErr := u.notifier.Notify(consts.NotifyEventCompleted,
		[]byte(fmt.Sprintf("completed DeleteCompany with id: %s", id))); notifyErr != nil {
		u.logger.Warnw("error notifying", "action", "DeleteCompany", "company-id", id)
	}
	return nil
}
