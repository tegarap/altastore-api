package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
)

func LoginAdmin(admin *models.Admins) (interface{}, error){
	var err error
	if err = config.Db.Where("email = ? AND password = ?", admin.Email, admin.Password).First(admin).Error; err != nil {
		return nil, err
	}
	outAdmin := models.AdminResponse{}
	return outAdmin, nil
}

func RegisterAdmin(admin *models.Admins) (interface{}, error) {
	if err := config.Db.Create(&admin).Error; err != nil {
		return nil, err
	}
	outAdmin := models.AdminResponse{}
	return outAdmin, nil
}