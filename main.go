package main

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/lib/middleware"
	"github.com/tegarap/altastore-api/routes"
)

func main() {
	config.InitDb()
	e := routes.New()
	middleware.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
