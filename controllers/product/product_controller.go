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
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get all product data", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get all product data", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully get all product data", newProduct))
}

func GetSingleProductController(c echo.Context) error {
	productId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Invalid product id", nil))
	}

	product, rowAffected, err := database.GetSingleProduct(productId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get product data", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to get product data", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully get product data", product))
}

func CreateNewProductController(c echo.Context) error {
	product := models.Products{}
	c.Bind(&product)

	newProduct, err := database.CreateNewProduct(&product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to create new product", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully create new product", newProduct))
}

func DeleteProductController(c echo.Context) error {
	productId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Invalid product id", nil))
	}
	result, rowAffected, err := database.DeleteProduct(productId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to delete product data", nil))
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to delete product data", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully delete product data", result))
}

func UpdateProductController(c echo.Context) error {
	productId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Invalid product id", nil))
	}

	newProduct := models.Products{}
	c.Bind(&newProduct)

	updatedProduct, rowAffected, err := database.UpdateProduct(productId, &newProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to update product data ", nil))
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Failed to updated product data", nil))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Successfully update product data", updatedProduct))
}
