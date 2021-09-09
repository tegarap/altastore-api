package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
	"net/http"
)

func LoginAdminController(c echo.Context) error {
	var admin models.Admins
	c.Bind(&admin)

	admins, err := database.LoginAdmin(&admin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Login Failed", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Login Success", admins))
}

func RegisterAdminController(c echo.Context) error {
	regAdmin := models.Admins{}
	c.Bind(&regAdmin)

	if regAdmin.Name == "" || regAdmin.Email == "" || regAdmin.Password == "" {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Field are Required", nil))
	}

	admin, err := database.RegisterAdmin(&regAdmin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Register Failed", nil))
	}
	return c.JSON(http.StatusBadRequest, util.ResponseSuccess("Register Success", admin))
}
