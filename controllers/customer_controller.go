package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
	"net/http"
)

func LoginCustomerController(c echo.Context) error {
	var customer models.Customers
	c.Bind(&customer)

	_, err := database.LoginCustomer(&customer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Login Failed", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Login Success", nil))
}

func RegisterCustomerController(c echo.Context) error {
	regCustomer := models.Customers{}
	c.Bind(&regCustomer)

	if regCustomer.Name == "" || regCustomer.Email == "" || regCustomer.Password == "" {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Field are Required", nil))
	}

	customer, err := database.RegisterCustomer(&regCustomer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Register Failed", nil))
	}
	return c.JSON(http.StatusBadRequest, util.ResponseSuccess("Register Success", customer))
}
