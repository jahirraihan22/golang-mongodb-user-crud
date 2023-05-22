package routes

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"ums/interfaces"
	"ums/src/service"
)

func Init(echo *echo.Echo) {
	var userApiGroup = echo.Group("/api/user")
	userApiGroup.Use(echojwt.WithConfig(service.MiddlewareControl()))
	userApiRouting(userApiGroup)

	var externalCommandApiGroup = echo.Group("/api/command")
	externalCommandApiGroup.Use(echojwt.WithConfig(service.MiddlewareControl()))
	externalCommandApiRouting(externalCommandApiGroup)

	// without group
	echo.POST("/login", service.Login) // login
}

func userApiRouting(userApiGroup *echo.Group) {
	userApiGroup.POST("", interfaces.GetUserCRUD().Create)
	userApiGroup.GET("/:id", interfaces.GetUserCRUD().Get)
	userApiGroup.PUT("/:id", interfaces.GetUserCRUD().Update)
	userApiGroup.DELETE("/:id", interfaces.GetUserCRUD().Delete)
	userApiGroup.GET("/all", interfaces.GetUserCRUD().GetAll)

}

func externalCommandApiRouting(externalCommandApiGroup *echo.Group) {
	// test curl command request
	externalCommandApiGroup.GET("/get-image-via-url", interfaces.GetCommand().CurlRequest)
	externalCommandApiGroup.GET("/create-vm", interfaces.GetCommand().CreateVm)
}
