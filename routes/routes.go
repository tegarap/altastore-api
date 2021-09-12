package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tegarap/altastore-api/controllers/admin"
	"github.com/tegarap/altastore-api/controllers/cart"
	"github.com/tegarap/altastore-api/controllers/category"
	"github.com/tegarap/altastore-api/controllers/customer"
	"github.com/tegarap/altastore-api/controllers/payment"
	"github.com/tegarap/altastore-api/controllers/product"
	"github.com/tegarap/altastore-api/controllers/transaction"
	"os"
)

func New() *echo.Echo {
	e := echo.New()
	jwtAuth := e.Group("")
	jwtAuth.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT"))))
	//customerAuth := e.Group("")
	//customerAuth.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT"))))


	//---------------------------------------
	//	ADMIN
	//---------------------------------------
	e.POST("/admin/login", admin.LoginAdminController)
	e.POST("/admin/register", admin.RegisterAdminController)
	// Only admin can access
	jwtAuth.GET("/admin/profile", admin.GetAdminProfileController)

	//---------------------------------------
	//	CUSTOMER
	//---------------------------------------
	e.POST("/customers/login", customer.LoginCustomerController)
	e.POST("/customers/register", customer.RegisterCustomerController)
	jwtAuth.GET("/customers/profile", customer.GetCustomerProfileController)
	// Only admin can access
	jwtAuth.GET("/customers", customer.GetAllCustomersController)

	//---------------------------------------
	//	CATEGORIES
	//---------------------------------------
	e.POST("/categories", category.CreateNewCategoriesController)
	e.GET("/categories", category.GetAllCategoriesController)
	e.GET("/categories/:id", category.GetSingleCategoryController)

	//---------------------------------------
	//	PRODUCTS
	//---------------------------------------
	e.POST("/products", product.CreateNewProductController)
	e.GET("/products", product.GetAllProductController)
	e.GET("/products/:id", product.GetSingleProductController)
	e.DELETE("/products/:id", product.DeleteProductController)
	e.PUT("/products/:id", product.UpdateProductController)

	//---------------------------------------
	//	PAYMENTS
	//---------------------------------------
	e.POST("/payments", payment.CreatePaymentController)
	e.GET("/payments", payment.GetAllPaymentController)
	e.GET("/payments/:id", payment.GetSinglePaymentController)

	//---------------------------------------
	//	CARTS
	//---------------------------------------
	e.POST("/carts", cart.CreateNewCartController)
	e.GET("/carts", cart.GetAllCartsController)
	e.GET("/carts/:id", cart.GetSingleCartController)

	//---------------------------------------
	//	CART DETAILS
	//---------------------------------------
	e.POST("/add/products", cart.CreateNewCartDetailController)
	e.GET("/carts/detail", cart.GetAllCartDetailController)
	e.GET("/carts/detail/:id", cart.GetSingleCartDetailController)
	e.PUT("/edit/products/:id", cart.UpdatedProductOnCartController)
	e.DELETE("/delete/products/:id", cart.DeleteProductOnCartController)

	//---------------------------------------
	//	TRANSACTIONS
	//---------------------------------------
	e.POST("/transactions", transaction.CreateTransactionController)
	e.GET("/transactions", transaction.GetAllTransactionController)
	e.GET("/transactions/:id", transaction.GetSingleTransactionController)

	return e
}
