package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/controllers"
)

func New() *echo.Echo {
	e := echo.New()

	//---------------------------------------
	//	ADMIN
	//---------------------------------------
	e.POST("/admin/login", controllers.LoginAdminController)
	e.POST("/admin/register", controllers.RegisterAdminController)

	//---------------------------------------
	//	CUSTOMER
	//---------------------------------------
	e.POST("/customers/login", controllers.LoginCustomerController)
	e.POST("/customers/register", controllers.RegisterCustomerController)

	//---------------------------------------
	//	CATEGORIES
	//---------------------------------------
	e.POST("/categories", controllers.CreateNewCategoriesController)
	e.GET("/categories", controllers.GetAllCategoriesController)
	e.GET("/categories/:id", controllers.GetSingleCategoryController)
	//---------------------------------------
	//	PRODUCTS
	//---------------------------------------
	e.POST("/products", controllers.CreateNewProductController)
	e.GET("/products", controllers.GetAllProductController)
	e.GET("/products/:id", controllers.GetSingleProductController)
	e.DELETE("/products/:id", controllers.DeleteProductController)
	//---------------------------------------
	//	CARTS
	//---------------------------------------
	e.POST("/carts", controllers.CreateNewCartController)
	return e
}
