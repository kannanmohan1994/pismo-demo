package user

import (
	"xm/internal/entity/models"
)

func (d *repo) GetUser(name, password string) (result *models.User, err error) {
	d.logger.Infof("Begin Repo - GetUser")

	err = d.db.Table("users").Where("name = ? AND password = ?", name, password).First(&result).Error
	if err != nil {
		d.logger.Errorf("error fetching from company", err.Error())
		return result, err
	}

	d.logger.Infof("End Repo - GetUser - %+v", result)
	return result, err
}

func (d *repo) CreateUser(user *models.User) (result *models.User, err error) {
	d.logger.Infof("Begin Repo - CreateUser")

	err = d.db.Table("users").Create(&user).Error
	if err != nil {
		d.logger.Errorf("error creating company", err.Error())
		return result, err
	}

	d.logger.Infof("End Repo - CreateUser %+v", user.ID)
	return user, err
}
