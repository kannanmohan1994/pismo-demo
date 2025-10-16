package company

import (
	"time"
	"xm/internal/entity/models"

	"gorm.io/gorm/clause"
)

func (d *repo) CreateCompany(company *models.Company) (result *models.Company, err error) {
	d.logger.Infof("Begin Repo - CreateCompany")

	err = d.db.Table("company").Create(&company).Error
	if err != nil {
		d.logger.Errorf("error creating company", err.Error())
		return result, err
	}

	d.logger.Infof("End Repo - CreateCompany - %v", company.ID)
	return company, err
}

func (d *repo) GetCompany(id string) (result *models.Company, err error) {
	d.logger.Infof("Begin Repo - GetCompany")

	err = d.db.Table("company").Where("id = ?", id).First(&result).Error
	if err != nil {
		d.logger.Errorf("error fetching from company", err.Error())
		return result, err
	}

	d.logger.Infof("End Repo - GetCompany - %+v", result)
	return result, err
}

func (d *repo) PatchCompany(id string, obj *models.Company) (result *models.Company, err error) {
	d.logger.Infof("Begin Repo - PatchCompany")

	updateMap := make(map[string]interface{}, 5)
	if len(obj.Name) > 0 {
		updateMap["name"] = obj.Name
	}
	if len(obj.Description) > 0 {
		updateMap["description"] = obj.Description
	}
	if obj.EmployeeCount != nil {
		updateMap["employee_count"] = obj.EmployeeCount
	}
	if obj.IsRegistered != nil {
		updateMap["is_registered"] = obj.IsRegistered
	}
	if len(obj.Type) > 0 {
		updateMap["type"] = obj.Type
	}

	if len(updateMap) > 0 {
		updateMap["updated_at"] = time.Now().UTC()

		result = &models.Company{ID: id}
		if err = d.db.Table("company").Model(&result).Clauses(clause.Returning{}).Updates(updateMap).Error; err != nil {
			d.logger.Errorf("error fetching from company", err.Error())
			return result, err
		}
	}

	d.logger.Infof("End Repo - PatchCompany - %+v", result)
	return result, err
}

func (d *repo) DeleteCompany(id string) (err error) {
	d.logger.Infof("Begin Repo - DeleteCompany")

	err = d.db.Table("company").Delete(&models.Company{}, "id = ?", id).Error
	if err != nil {
		d.logger.Errorf("error fetching from company", err.Error())
		return err
	}

	d.logger.Infof("End Repo - DeleteCompany")
	return err
}
