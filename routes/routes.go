package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/tegarap/altastore-api/controllers"
	"github.com/tegarap/altastore-api/controllers/admin"
	"github.com/tegarap/altastore-api/controllers/customer"
	"github.com/tegarap/altastore-api/controllers/payment"
	"github.com/tegarap/altastore-api/controllers/transaction"
)

func New() *echo.Echo {
	e := echo.New()

	//---------------------------------------
	//	ADMIN
	//---------------------------------------
	e.POST("/admin/login", admin.LoginAdminController)
	e.POST("/admin/register", admin.RegisterAdminController)

	//---------------------------------------
	//	CUSTOMER
	//---------------------------------------
	e.POST("/customers/login", customer.LoginCustomerController)
	e.POST("/customers/register", customer.RegisterCustomerController)

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

	//e.POST("/deliveries", controllers.CreateDeliveryController)
	//e.GET("/deliveries", controllers.GetAllDeliveryController)
	//e.GET("/deliveries/:id", controllers.GetSingleDeliveryController)

	//---------------------------------------
	//	PAYMENTS
	//---------------------------------------
	e.POST("/payments", payment.CreatePaymentController)
	e.GET("/payments", payment.GetAllPaymentController)
	e.GET("/payments/:id", payment.GetSinglePaymentController)

	//---------------------------------------
	//	CARTS
	//---------------------------------------
	//e.POST("/carts", controllers.CreateNewCartController)


	//---------------------------------------
	//	TRANSACTIONS
	//---------------------------------------
	e.POST("/transactions", transaction.CreateTransactionController)
	e.GET("/transactions", transaction.GetAllTransactionController)
	e.GET("/transactions/:id", transaction.GetSingleTransactionController)


	return e
}
