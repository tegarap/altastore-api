package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
	"net/http"
	"strconv"
)

func CreateDeliveryController(c echo.Context) error {
	var delivery models.Deliveries
	c.Bind(&delivery)

	newDelivery, err := database.CreateDelivery(&delivery)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Create New Delivery", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Create New Delivery", newDelivery))
}

func GetAllDeliveryController(c echo.Context) error {
	var delivery []models.Deliveries
	deliveries, err := database.GetAllDelivery(&delivery)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseSuccess("Fail to Get All Delivery", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Get All Delivery", deliveries))
}

func GetSingleDeliveryController(c echo.Context) error {
	deliveryId, er := strconv.Atoi(c.Param("id"))
	if er != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Invalid Parameter", nil))
	}

	delivery, err := database.GetSingleDelivery(deliveryId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to Get Single Delivery", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Get Single Delivery", delivery))
}
