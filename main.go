package altastore_api

import (
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/routes"
)

func main() {
	config.InitDb()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
