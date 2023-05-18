package main

import (
	"net/http"
	"ums/routes"
	"ums/src/models"

	"github.com/labstack/echo/v4"
)

func main() {
	models.InitializeDatabase()

	var server = echo.New()

	server.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Running..........")
	})

	routes.Init(server)
	server.Logger.Fatal(server.Start(":12375"))
}
