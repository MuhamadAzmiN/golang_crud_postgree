// main.go
package main

import (
	"myapp/db"
	"myapp/routes"
	"github.com/labstack/echo/v4"
)

func main() {
    db.ConnectDB()
    e := echo.New()
    routes.InitSiswaRoutes(e)
    routes.InitAuthRoutes(e)
    e.Logger.Fatal(e.Start(":8080"))
}
