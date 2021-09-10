package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
	"net/http"
)

func CreatePaymentController(c echo.Context) error {
	var payment models.Payments
	c.Bind(&payment)

	newPayment, err := database.CreatePayment(&payment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Create Payment Method", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Create New Payment Method", newPayment))
}
