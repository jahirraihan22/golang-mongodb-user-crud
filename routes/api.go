package routes

import (
	"ums/interfaces"

	"github.com/labstack/echo/v4"
)

func Init(echo *echo.Echo) {
	var userApiGroup = echo.Group("/api/user")
	userApiRouting(userApiGroup)

	var externalCommandApiGroup = echo.Group("/api/command")

	externalCommandApiRouting(externalCommandApiGroup)
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
