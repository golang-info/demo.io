package helloWorld

import (
	"demo.io/demo-echo/utils"
	"github.com/labstack/echo/v4"
)

func HelloWorld(c echo.Context) error {
	//return c.String(http.StatusOK, "Hello, World!")
	return utils.NewResponse("Hello, World!")

}
