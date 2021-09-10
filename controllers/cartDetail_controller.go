package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
)

func CreateNewCartDetailController(c echo.Context) error {
	cartDetail := models.CartsDetail{}
	c.Bind(&cartDetail)

	newCartDetail, err := database.CreateNewCartDetail(&cartDetail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", newCartDetail))
}
