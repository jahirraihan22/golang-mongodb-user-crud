package routes

import (
	"github.com/labstack/echo/v4"
	"ums/interfaces"
)

func Init(echo *echo.Echo) {
	var userApiGroup = echo.Group("/api/user")
	userApiRouting(userApiGroup)
}

func userApiRouting(userApiGroup *echo.Group) {
	userApiGroup.POST("", interfaces.GetUserCRUD().Create)
	userApiGroup.GET("/:id", interfaces.GetUserCRUD().Get)
	userApiGroup.PUT("/:id", interfaces.GetUserCRUD().Update)
	userApiGroup.DELETE("/:id", interfaces.GetUserCRUD().Delete)
	userApiGroup.GET("/all", interfaces.GetUserCRUD().GetAll)
}