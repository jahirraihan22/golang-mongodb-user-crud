package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ums/router"
	"ums/src/models"
)

func main() {
	models.InitializeDatabase()

	var server = echo.New()

	server.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Running..........")
	})

	router.Init(server)
	server.Logger.Fatal(server.Start(":12375"))
}
