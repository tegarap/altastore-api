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
	return e
}
