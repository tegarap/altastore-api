package database

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/lib/middleware"
	"github.com/tegarap/altastore-api/models"
)

func LoginAdmin(admin *models.Admins) (interface{}, error){
	var err error
	if err = config.Db.Where("email = ? AND password = ?", admin.Email, admin.Password).First(admin).Error; err != nil {
		return nil, err
	}
	admin.Token, err = middleware.CreateToken(int(admin.ID), true)
	if err != nil {
		return nil, err
	}
	if err = config.Db.Save(admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func RegisterAdmin(admin *models.Admins) (interface{}, error) {
	if err := config.Db.Create(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func GetAdminProfile(id int) (interface{}, error) {
	admin := models.Admins{}
	if err := config.Db.Find(&admin, id).Error; err != nil {
		return nil, err
	}
	return admin, nil
}