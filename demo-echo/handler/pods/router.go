package pods

import (
	"github.com/labstack/echo/v4"
)

func InitRouter(group *echo.Group) {
	group.GET("/GetPods", GetPods)
}