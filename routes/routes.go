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
	e.PUT("/products/:id", controllers.UpdateProductController)

	//---------------------------------------
	//	DELIVERIES
	//---------------------------------------
	e.POST("/deliveries", controllers.CreateDeliveryController)
	e.GET("/deliveries", controllers.GetAllDeliveryController)
	e.GET("/deliveries/:id", controllers.GetSingleDeliveryController)

	//---------------------------------------
	//	PAYMENTS
	//---------------------------------------
	e.POST("/payments", controllers.CreatePaymentController)
	e.GET("/payments", controllers.GetAllPaymentController)
	e.GET("/payments/:id", controllers.GetSinglePaymentController)

	//---------------------------------------
	//	CARTS
	//---------------------------------------
	//e.POST("/carts", controllers.CreateNewCartController)


	//---------------------------------------
	//	TRANSACTIONS
	//---------------------------------------
	e.POST("/transactions", controllers.CreateTransactionController)
	e.GET("/transactions", controllers.GetAllTransactionController)
	e.GET("/transactions/:id", controllers.GetSingleTransactionController)


	return e
}
