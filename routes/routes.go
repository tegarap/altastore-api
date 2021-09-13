package routes

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tegarap/altastore-api/controllers/admin"
	"github.com/tegarap/altastore-api/controllers/cart"
	"github.com/tegarap/altastore-api/controllers/category"
	"github.com/tegarap/altastore-api/controllers/customer"
	"github.com/tegarap/altastore-api/controllers/payment"
	"github.com/tegarap/altastore-api/controllers/product"
	"github.com/tegarap/altastore-api/controllers/transaction"
)

func New() *echo.Echo {
	e := echo.New()
	jwtAuth := e.Group("")
	jwtAuth.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT"))))

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
	jwtAuth.POST("/categories", category.CreateNewCategoriesController)
	e.GET("/categories", category.GetAllCategoriesController)
	e.GET("/categories/:id", category.GetSingleCategoryController)
	jwtAuth.PUT("/categories/:id", category.UpdateCategoryController)
	jwtAuth.DELETE("/categories/:id", category.DeleteCategoryController)
	//---------------------------------------
	//	PRODUCTS
	//---------------------------------------
	jwtAuth.POST("/products", product.CreateNewProductController)
	e.GET("/products", product.GetAllProductController)
	e.GET("/products/:id", product.GetSingleProductController)
	jwtAuth.DELETE("/products/:id", product.DeleteProductController)
	jwtAuth.PUT("/products/:id", product.UpdateProductController)

	//---------------------------------------
	//	PAYMENTS
	//---------------------------------------
	jwtAuth.POST("/payments", payment.CreatePaymentMethodController)
	jwtAuth.GET("/payments", payment.GetAllPaymentMethodController)
	jwtAuth.GET("/payments/:id", payment.GetSinglePaymentMethodController)
	jwtAuth.DELETE("/payments/:id", payment.DeletePaymentMethodController)
	jwtAuth.PUT("/payments/:id", payment.UpdatePaymentMethodController)

	//---------------------------------------
	//	CARTS
	//---------------------------------------
	jwtAuth.POST("/carts", cart.CreateNewCartController)
	jwtAuth.GET("/carts", cart.GetAllCartsController)
	jwtAuth.GET("/cart", cart.GetSingleCartController)
	jwtAuth.GET("/customers/carts", cart.GetCustomersCartsController)
	jwtAuth.GET("/customers/carts/:cart_id", cart.GetSingleCustomersCartController)
	jwtAuth.DELETE("/carts/:id", cart.DeleteCartController)
	e.PUT("/carts/:id", cart.UpdatedCartStatusController)
	//---------------------------------------
	//	CART DETAILS
	//---------------------------------------
	jwtAuth.POST("/add/products", cart.AddProductOnCart)
	jwtAuth.PUT("/edit/products/:id", cart.UpdatedProductOnCartController)
	jwtAuth.DELETE("/delete/products/:id", cart.DeleteProductOnCartController)

	//---------------------------------------
	//	TRANSACTIONS
	//---------------------------------------
	jwtAuth.POST("/transactions", transaction.CreateTransactionController)
	jwtAuth.GET("/transactions", transaction.GetTransactionController)

	return e
}
