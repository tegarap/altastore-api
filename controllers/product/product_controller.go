package product

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/lib/database"
	"github.com/tegarap/altastore-api/models"
	util "github.com/tegarap/jsonres"
)

func GetAllProductController(c echo.Context) error {
	product := []models.Products{}

	newProduct, rowAffected, err := database.GetAllProduct(&product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", newProduct))
}

func GetSingleProductController(c echo.Context) error {
	productId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}

	product, rowAffected, err := database.GetSingleProduct(productId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", product))
}

func CreateNewProductController(c echo.Context) error {
	product := models.Products{}
	c.Bind(&product)

	newProduct, err := database.CreateNewProduct(&product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", newProduct))
}

func DeleteProductController(c echo.Context) error {
	productId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	result, rowAffected, err := database.DeleteProduct(productId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", result))
}

func UpdateProductController(c echo.Context) error {
	productId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}

	newProduct := models.Products{}
	c.Bind(&newProduct)

	updatedProduct, rowAffected, err := database.UpdateProduct(productId, &newProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("failed", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("success", updatedProduct))
}
