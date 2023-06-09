package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ums/routes"
	"ums/src/database"
)

func main() {
	database.Init()

	var server = echo.New()

	server.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Running..........")
	})

	routes.Init(server)
	server.Logger.Fatal(server.Start(":12375"))
}
