package database

import "github.com/tegarap/altastore-api/models"

func LoginAdmin(admin *models.Admins) (interface{}, error){
	var err error
	if err = config.DB.Where("email = ? AND password = ?", admin.Email, admin.Password).First(admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func RegisterAdmin(admin *models.Admins) (interface{}, error) {
	if err := config.DB.Create(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}