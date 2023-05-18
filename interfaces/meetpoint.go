package interfaces

import (
	"github.com/labstack/echo/v4"
	"ums/src/service"
)

type Crud interface {
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Get(ctx echo.Context) error
	GetAll(ctx echo.Context) error
}

type Command interface {
	CurlRequest(ctx echo.Context) error
	CreateVm(ctx echo.Context) error
}

func GetUserCRUD() Crud {
	return &service.UserManagement{}
}

func GetCommand() Command {
	return &service.CommandManagement{}
}
