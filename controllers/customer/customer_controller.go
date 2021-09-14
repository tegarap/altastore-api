package customer

import (
	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/lib/middleware"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
	"net/http"
	"strconv"
)

func LoginCustomerController(c echo.Context) error {
	var customer models.Customers
	c.Bind(&customer)

	customers, err := database.LoginCustomer(&customer)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Login Failed", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Login Success", customers))
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

	return c.JSON(http.StatusOK, util.ResponseSuccess("Register Success", customer))
}

func GetCustomerProfileController(c echo.Context) error {
	id, isAdmin := middleware.ExtractToken(c)

	if isAdmin == true {
		customerId := c.QueryParam("id")
		getId, _ := strconv.Atoi(customerId)
		customer, err1 := database.GetCustomerProfile(getId)

		if err1 != nil {
			return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Get Customer Profile", nil))
		}

		return c.JSON(http.StatusOK, util.ResponseSuccess("Success Get Customer Profile", customer))
	} else {
		customer, err := database.GetCustomerProfile(id)

		if err != nil {
			return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Get Profile", nil))
		}

		return c.JSON(http.StatusOK, util.ResponseSuccess("Success Get Profile", customer))
	}
}

func GetAllCustomersController(c echo.Context) error {
	var customers []models.Customers

	_, isAdmin := middleware.ExtractToken(c)

	if isAdmin != true {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Only Admin can Access", nil))
	}

	allCustomers, err := database.GetAllCustomerSimple(&customers)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("success", allCustomers))
}
