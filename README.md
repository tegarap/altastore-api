# Alta Online Store API Server

REST API build using echo server.

These API were taken from projects mainly using [Echo Framework](https://echo.labstack.com/) and the code implementation
was inspired by Model View Controller [MVC](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller):

- **Models**<br/>Contains implementation model database and quary with ORM.
- **View**<br/>Not implemented in this code because not required User Interview (View)
- **Contoller (API)**<br/>API http handler or controller

Where full URLs are provided in responses they will be rendered as if service is running on 'localhost'.

## Admin

* [Register Admin](#) : `POST /admin/register`
* [Login Admin](#) : `POST /admin/login`
* [Admin Profile](#) : `POST /admin/profile`

## Customers

* [Register Customer](#) : `POST /customers/register`
* [Login Customer](#) : `POST /customers/login`
* [Admin Customer](#) : `POST /customers/profile`
* [Get All Customer](#) : `POST /customers`

## Categories

* [Create New Categories](#) : `POST /categories`
* [Get All Categories](#) : `GET /categories`
* [Get Single Category](#) : `GET /categories/:id`
* [Update Category](#) : `PUT /categories/:id`
* [Delete Category](#) : `DELETE /categories/:id`

## Products

* [Create New Product](#) : `POST /products`
* [Get All Product](#) : `GET /products`
* [Get Single Product](#) : `GET /products/:id`
* [Delete Product](#) : `DELETE /products/:id`
* [Update Product](#) : `PUT /products/:id`

## Products

* [Create New Payment Methods](#) : `POST /payments`
* [Get All Payment Method](#) : `GET /payments`
* [Get Single Payment Method](#) : `GET /payments/:id`
* [Delete Payment Method](#) : `DELETE /payments/:id`
* [Update Payment Method](#) : `PUT /payments/:id`

## Carts

* [Create New Cart](#) : `POST /carts`
* [Get All Cart](#) : `GET /carts`
* [Get Single Cart](#) : `GET /carts/:id`
* [Get All Customer Cart](#) : `GET /customers/carts`
* [Get Single Customer Cart](#) : `GET /customers/carts/:id`
* [Delete Customer Cart](#) : `DELETE /carts/:id`

## Cart Details

* [Add Product on Cart](#) : `POST /carts/products`
* [Update Products on Cart](#) : `PUT /carts/prouducts/:id`
* [Delete Products on Cart](#) : `DELETE /carts/products/:id`

## Transactions

* [Create Transaction](#) : `POST /transactions`
* [Get Transaction](#) : `GET /transactions`