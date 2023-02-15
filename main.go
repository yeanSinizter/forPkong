package main

import (
	"eiei/clean_architecture/db"
	"eiei/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	SqlDb, teardown := db.InitSql()
	defer teardown()

	routes.InitRoute(e, SqlDb)

	e.Logger.Fatal(e.Start(":3000"))

}
