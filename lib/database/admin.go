package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
)

func LoginAdmin(admin *models.Admins) (*models.Admins, error){
	var err error
	if err = config.Db.Where("email = ? AND password = ?", admin.Email, admin.Password).First(admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func RegisterAdmin(admin *models.Admins) (*models.Admins, error) {
	if err := config.Db.Create(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}