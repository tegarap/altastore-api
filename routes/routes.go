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
	e.PUT("/categories/:id", category.UpdateCategoryController)
	e.DELETE("/categories/:id", category.DeleteCategoryController)
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
	jwtAuth.POST("/payments", payment.CreatePaymentMethodController)
	jwtAuth.GET("/payments", payment.GetAllPaymentMethodController)
	jwtAuth.GET("/payments/:id", payment.GetSinglePaymentMethodController)
	jwtAuth.DELETE("/payments/:id", payment.DeletePaymentMethodController)
	jwtAuth.PUT("/payments/:id", payment.UpdatePaymentMethodController)


	//---------------------------------------
	//	CARTS
	//---------------------------------------
	e.POST("/carts", cart.CreateNewCartController)
	e.GET("/carts", cart.GetAllCartsController)
	e.GET("/carts/:id", cart.GetSingleCartController)
	e.GET("/customers/:id/carts", cart.GetCustomersCartsController)
	e.GET("/customers/:id/carts/:cart_id", cart.GetSingleCustomersCartController)
	e.DELETE("/carts/:id", cart.DeleteCartController)
	e.PUT("/carts/:id", cart.UpdatedCartStatusController)
	//---------------------------------------
	//	CART DETAILS
	//---------------------------------------
	e.POST("/add/products", cart.AddProductOnCart)
	e.GET("/carts/detail", cart.GetAllCartDetailController)
	e.GET("/carts/detail/:id", cart.GetSingleCartDetailController)
	e.PUT("/edit/products/:id", cart.UpdatedProductOnCartController)
	e.DELETE("/delete/products/:id", cart.DeleteProductOnCartController)

	//---------------------------------------
	//	TRANSACTIONS
	//---------------------------------------
	jwtAuth.POST("/transactions", transaction.CreateTransactionController)
	jwtAuth.GET("/transactions", transaction.GetTransactionController)

	return e
}
