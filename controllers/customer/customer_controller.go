package customer

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
)

func LoginCustomerController(c echo.Context) error {
	var customer models.Customers
	c.Bind(&customer)

	customers, err := database.LoginCustomer(&customer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Login Failed", nil))
	}

	response := Response{
		ID: customers.ID,
		Name: customers.Name,
		Email: customers.Email,
		Phone: customers.Phone,
		Address: customers.Address,
		Gender: customers.Gender,
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Login Success", response))
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

	response := Response{
		ID: customer.ID,
		Name: customer.Name,
		Email: customer.Email,
		Phone: customer.Phone,
		Address: customer.Address,
		Gender: customer.Gender,
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Register Success", response))
}

func GetAllCustomersController(c echo.Context) error {
	customers := []models.Customers{}

	allCustomers, err := database.GetAllCustomers(&customers)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", allCustomers))
}
