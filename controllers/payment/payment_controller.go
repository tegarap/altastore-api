package payment

import (
	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/lib/middleware"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
	"net/http"
	"strconv"
)

func CreatePaymentMethodController(c echo.Context) error {
	_, isAdmin := middleware.ExtractToken(c)

	if isAdmin != true {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Only Admin can Access", nil))
	}

	var payment models.Payments
	c.Bind(&payment)

	newPayment, err := database.CreatePaymentMethod(&payment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Create Payment Method", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Create New Payment Method", newPayment))
}

func GetAllPaymentMethodController(c echo.Context) error {
	var payment []models.Payments

	payments, err := database.GetAllPaymentMethod(&payment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Get All Payment Method", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Get All Payment Method", payments))
}

func GetSinglePaymentMethodController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Invalid Parameter", nil))
	}

	payment, errr := database.GetSinglePaymentMethod(id)
	if errr != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Get Payment Method", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Get Payment Method", payment))
}
