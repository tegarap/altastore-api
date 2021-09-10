package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
)

func CreateNewCategoriesController(c echo.Context) error {
	category := models.Categories{}
	c.Bind(&category)

	newCategory, err := database.CreateNewCategories(&category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", newCategory))
}

func GetAllCategoriesController(c echo.Context) error {
	categories := []models.Categories{}

	category, err := database.GetAllCategories(&categories)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", category))
}
