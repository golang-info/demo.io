package helloWorld

import (
	"github.com/labstack/echo/v4"
)

func InitRouter(group *echo.Group) {
	group.GET("/helloWorld", HelloWorld)
}
